package service

import (
	"project/app/march_voiced/models"
	"project/app/march_voiced/models/dto"
)

// 添加文章
func InsertArticle(articleDto *dto.InsertArticleDto, userId int) (err error) {
	// 实例化要添加的文章
	article := new(models.Article)
	article.CreatBy = uint(userId)
	article.UpdateBy = uint(userId)
	article.Title = articleDto.Title
	article.Content = articleDto.Content
	article.Image = articleDto.Image
	article.Kind = articleDto.Kind
	article.Tag = articleDto.Tag
	article.Status = articleDto.Status

	// 调用dao方法
	err = article.InsertArticle()
	return
}

// 更新文章
func UpdateArticle(articleDto *dto.UpdateArticleDto, userId int) (err error) {
	// 实例化要修改的文章
	article := new(models.Article)
	article.UpdateBy = uint(userId)
	article.Title = articleDto.Title
	article.Content = articleDto.Content
	article.Image = articleDto.Image
	article.Kind = articleDto.Kind
	article.Tag = articleDto.Tag
	article.Status = articleDto.Status
	article.Type = articleDto.Type

	// 调用dao方法
	err = article.UpdateArticle(articleDto.ID)
	return
}
