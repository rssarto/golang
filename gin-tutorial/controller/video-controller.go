package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rssarto/gin-tutorial/entity"
	"github.com/rssarto/gin-tutorial/service"
	"github.com/rssarto/gin-tutorial/validators"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Video
}

type videoController struct {
	service service.VideoService
}

func (videoController *videoController) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(video)
	if err != nil {
		return err
	}
	return nil
}

func (videoController *videoController) FindAll() []entity.Video {
	return videoController.service.FindAll()
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoController{
		service: service,
	}
}
