package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rssarto/gin-tutorial/controller"
	"github.com/rssarto/gin-tutorial/dto"
	"github.com/rssarto/gin-tutorial/repository"
	"github.com/rssarto/gin-tutorial/service"
)

var (
	loginService    service.LoginService       = service.NewLoginService()
	jwtService      service.JWTService         = service.NewJwtService()
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)

	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	videoController controller.VideoController = controller.New(videoService)
)

// Authenticate godoc
// @Summary      Provides a JSON Web Token
// @Description  Authenticates a user and provides a JWT to Authorize API calls
// @Accept       application/x-www-form-urlencoded
// @Produce      json
// @Tags		 Authentication
// @Param        username   formData      string  true  "User Credentials"
// @Param        password   formData      string  true  "User Credentials"
// @Success      200  {object}  dto.JWT
// @Failure      401  {object}  dto.Response
// @Router       /login [post]
func Authenticate(ctx *gin.Context) {
	token := loginController.Login(ctx)
	if token != "" {
		ctx.JSON(http.StatusOK, &dto.JWT{
			Token: token,
		})
	} else {
		ctx.JSON(http.StatusUnauthorized, &dto.Response{
			Message: "Not Authorized",
		})
	}
}

// GetVideos godoc
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Summary 	List all videos
// @Description Get all existing videos
// @Tags 		videos
// @Produce 	json
// @Success		200	{array} 	entity.Video
// @Failure		401 {object}	dto.Response
// @Router		/videos	[get]
func GetVideos(ctx *gin.Context) {
	ctx.JSON(200, videoController.FindAll())
}

// CreateVideo godoc
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @param video body entity.Video true "Update Video"
// @Summary 	Create video
// @Description Create a new video
// @Tags 		videos
// @Produce 	json
// @Success		200	{array} 	entity.Video
// @Failure		401 {object}	dto.Response
// @Router		/videos	[post]
func CreateVideo(ctx *gin.Context) {
	err := videoController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Video input is Valid!!",
		})
	}
}

// UpdateVideo godoc
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @param id path int true "Video ID"
// @param video body entity.Video true "Update Video"
// @Summary 	Update video
// @Description Update video by id
// @Tags 		videos
// @Produce 	json
// @Success		200	{array} 	entity.Video
// @Failure		401 {object}	dto.Response
// @Router		/videos/{id}	[put]
func UpdateVideo(ctx *gin.Context) {
	err := videoController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Video input is Valid!!",
		})
	}
}

// DeleteVideo godoc
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @param id path int true "Video ID"
// @Summary 	Delete Video
// @Description Delete video by id
// @Tags 		videos
// @Produce 	json
// @Success		200	{array} 	entity.Video
// @Failure		401 {object}	dto.Response
// @Router		/videos/{id}	[delete]
func DeleteVideo(ctx *gin.Context) {
	err := videoController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Video input is Valid!!",
		})
	}
}

func ViewVideos(ctx *gin.Context) {
	videoController.ShowAll(ctx)
}
