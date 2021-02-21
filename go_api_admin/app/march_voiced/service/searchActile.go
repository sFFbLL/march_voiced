package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/common/global"
)

type Search struct {
}


func (search *Search) SearchActile(searchInfo dto.SearchActileDto)( []bo.SearchActileBo,error) {
	if searchInfo.Size<0 || searchInfo.Current<1 {
		fmt.Println("错误")
		return nil,errors.New("分页参数错误")
	}
	hightmap, err := QueryHeight(searchInfo, global.ElasticSearch)
	if err != nil {
		return nil,err
	}
	searchActile := new(models.Seach)
	actile, err := searchActile.SearchActile(hightmap)
	if err != nil {
		return nil,err
	}
	return actile,nil
}

func (search *Search) SearchUser(searchInfo dto.SearchActileDto)([]bo.SearchUserBo,error){
	if searchInfo.Size<0 || searchInfo.Current<1 {
		fmt.Println("错误")
		return nil,errors.New("分页参数错误")
	}
	searchUserDao := new(models.Seach)
	userInfo, err := searchUserDao.SearchUser(searchInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}



func QueryHeight(search dto.SearchActileDto,client *elastic.Client) (hightmap []bo.SearchActileBo,err error){
	var data []bo.SearchActileBo
	textFild:=elastic.NewHighlighterField("content")
	//idFild:=elastic.NewHighlighterField("id")
	//fild.PreTags("<em>")
	//fild.PostTags("</em>")
	titleFild:=elastic.NewHighlighterField("title")
	titleFild.FragmentSize(350)
	higthFild :=elastic.NewHighlight()
	higthFild.Fields(titleFild)
	higthFild.Fields(textFild)
	//higthFild.Fields(idFild)
	//两个里面必须有
	//query :=elastic.NewBoolQuery().Must(elastic.NewMatchPhraseQuery("title","css"),elastic.NewMatchPhraseQuery("text","css"))
	query :=elastic.NewBoolQuery().Should(elastic.NewMatchPhraseQuery("title",search.SearchWord),elastic.NewMatchPhraseQuery("content",search.SearchWord))
	res, err := client.Search("article").Query(query).Highlight(higthFild).Size(search.Size).From((search.Current-1)*search.Size).Do(context.Background())
	if err != nil {
		return nil,err
	}
	//fmt.Println(res.Hits.Hits)
	for _, hit := range res.Hits.Hits {
		aa :=hit.Highlight
		var search bo.SearchActileBo
		search.ContentId = hit.Id
		var qq dto.Article
		json.Unmarshal(hit.Source,&qq)
		search.CreatTime = fmt.Sprintf("%v",qq.CreateTime)
		search.CreateID = qq.CreateBy
		if aa["title"] == nil{
			search.Title=qq.Title
		}else{
			search.Title=aa["title"][0]
		}
		if aa["content"] == nil{
			search.Content=qq.Content
		}else{
			search.Content = aa["content"][0]
		}
		data = append(data, search)
	}
	return data,nil
}