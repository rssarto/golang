package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rssarto/gin-tutorial/entity"
	"github.com/rssarto/gin-tutorial/service"
)

type VideoController interface {
	Save(ctx *gin.Context) entity.Video
	FindAll() []entity.Video
}

type videoController struct {
	service service.VideoService
}

func (videoController *videoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	return videoController.service.Save(video)
}

func (videoController *videoController) FindAll() []entity.Video {
	return videoController.service.FindAll()
}

func New(service service.VideoService) VideoController {
	return &videoController{
		service: service,
	}
}
