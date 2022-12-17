package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rssarto/gin-tutorial/entity"
	"github.com/rssarto/gin-tutorial/service"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Video
	ShowAll(ctx *gin.Context)
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
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

	// err = validate.Struct(video)
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

func (videoController *videoController) Update(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	video.ID = id
	// err = validate.Struct(video)
	if err != nil {
		return err
	}
	videoController.service.Update(video)
	return nil
}

func (videoController *videoController) Delete(ctx *gin.Context) error {
	var video entity.Video
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	video.ID = id
	videoController.service.Delete(video)
	return nil
}

// var validate *validator.Validate

func New(service service.VideoService) VideoController {
	// validate = validator.New()
	// validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &videoController{
		service: service,
	}
}
