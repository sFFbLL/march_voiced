package models

import (
	"fmt"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/common/global"
)
type Seach struct {
}


func (seach *Seach)SearchActile(data []bo.SearchActileBo) ([]bo.SearchActileBo, error) {
	var user SysUser
	for index,val :=range data{
		err := global.Eloquent.First(&user, "id=?", val.CreateID).Error
		if err != nil {
			fmt.Println(err,"出错了")
			return nil,err
		}
		data[index].Nickname = user.Username
		var count int64
		global.Eloquent.Model(&ArticleComment{}).Where("article_id=?",val.ContentId).Count(&count)
		data[index].CommentTotal = fmt.Sprintf("%v",count)
		global.Eloquent.Model(&ArticleFavour{}).Where("article_id=?",val.ContentId).Count(&count)
		data[index].FavourTotal = fmt.Sprintf("%v",count)
	}
	return data,nil
}

func (seach *Seach) SearchUser(searchInfo dto.SearchActileDto) ([]bo.SearchUserBo, error) {
	var user []SysUser
	fmt.Println(searchInfo.SearchWord)
	//word := "%"+searchInfo.SearchWord+ "%"

	err := global.Eloquent.Where(fmt.Sprintf("username like '%%%s%%' ",searchInfo.SearchWord)).Limit(int(searchInfo.Size)).Offset(int((searchInfo.Current - 1) * searchInfo.Size)).Find(&user).Error
	//db.Model(&User{}).Where("price >= ?",0).Count(&total)
	if err != nil {
		return nil,err
	}
	//遍历数据表格
	//然后判断是否关注



	return nil,nil

}

