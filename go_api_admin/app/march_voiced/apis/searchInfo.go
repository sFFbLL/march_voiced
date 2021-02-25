package apis

import (
	"project/app/march_voiced/models/dto"
	"project/app/march_voiced/service"
	"project/utils/app"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SearchActile(c *gin.Context)  {
	p := new(dto.SearchActileDto)
	if err := c.ShouldBindQuery(p); err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	search := new(service.Search)
	searchActile, err := search.SearchActile(*p)
	if err != nil {
		return
	}
	app.ResponseSuccess(c,searchActile)

}

func SearchUser(c *gin.Context)  {
	userId, _ := c.Get("UserId")
	p := new(dto.SearchActileDto)
	if err := c.ShouldBindQuery(p); err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			app.ResponseError(c, app.CodeParamIsInvalid)
			return
		}
		app.ResponseError(c, app.CodeParamNotComplete)
		return
	}
	search := new(service.Search)
	searchUser, err := search.SearchUser(*p,userId)
	if err != nil {
		return
	}
	app.ResponseSuccess(c,searchUser)

}

