package models

import (
	"fmt"
	"go.uber.org/zap"
	"project/app/march_voiced/models/bo"
	"project/app/march_voiced/models/dto"
	"project/common/global"
	"strconv"
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

func (seach *Seach) SearchUser(searchInfo dto.SearchActileDto,userId interface{}) ([]bo.SearchUserBo, error) {
	var user []SysUser
	follow := new(Follow)
	var data []bo.SearchUserBo
	fmt.Println(searchInfo.SearchWord)
	//word := "%"+searchInfo.SearchWord+ "%"
	err := global.Eloquent.Where(fmt.Sprintf("username like '%%%s%%' ",searchInfo.SearchWord)).Limit(int(searchInfo.Size)).Offset(int((searchInfo.Current - 1) * searchInfo.Size)).Find(&user).Error
	//db.Model(&User{}).Where("price >= ?",0).Count(&total)
	if err != nil {
		return nil,err
	}
	//遍历数据表格
	//然后判断是否关注
	for _, val := range user{
		atoi, err := strconv.Atoi(fmt.Sprintf("%v", userId))
		if err != nil {
			return nil,err
		}
		is, err := follow.IsFollow(atoi, int(val.ID))
		if err != nil {
			zap.L().Error("IsFollow failed", zap.Error(err))
			err = nil
		}
		data = append(data, bo.SearchUserBo{
			AvatarPath: val.AvatarPath,
			NickName:   val.Username,
			UserId:    val.ID,
			Signature: val.Signature,
			IsFollow:   is,
		})
	}
	return data,nil

}

