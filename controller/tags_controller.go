package controller

import (
	"CRUD/data/request"
	"CRUD/data/response"
	"CRUD/helper"
	"CRUD/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagsController struct {
	TagsService services.TagsService
}

func NewTagsController(services services.TagsService) *TagsController {
	return &TagsController{
		TagsService: services,
	}
}

func (controller *TagsController) Create(ctx *gin.Context) {
	CreateTagsRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&CreateTagsRequest)
	helper.ErrorPanic(err)

	controller.TagsService.Create(CreateTagsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   CreateTagsRequest,
	}
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *TagsController) Update(ctx *gin.Context) {
	updateTagRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagRequest.Id = id

	controller.TagsService.Update(updateTagRequest)
}

func (controller *TagsController) Delete(ctx *gin.Context) {

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.TagsService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagsController) FindByID(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagResponse := controller.TagsService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagsController) FindAll(ctx *gin.Context) {
	tagResponse := controller.TagsService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tagResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
