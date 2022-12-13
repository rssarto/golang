package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rssarto/gin-tutorial/entity"
	"github.com/rssarto/gin-tutorial/service"
	"github.com/rssarto/gin-tutorial/validators"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Video
	ShowAll(ctx *gin.Context)
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
	videoController.service.Save(video)
	return nil
}

func (videoController *videoController) FindAll() []entity.Video {
	return videoController.service.FindAll()
}

func (videoController *videoController) ShowAll(ctx *gin.Context) {
	//Fetch all available videos and assign to a variable
	videos := videoController.service.FindAll()

	//Create a gin data strucure to send data to the template
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}

	//Send the data to the template
	ctx.HTML(http.StatusOK, "index.html", data)
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoController{
		service: service,
	}
}
