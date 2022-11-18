package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rssarto/gin-tutorial/controller"
	"github.com/rssarto/gin-tutorial/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	firstNumber := 5
	var secondNumber *int = &firstNumber
	thirdNumber := firstNumber
	fmt.Printf("firstNumber = %d, secondNumber = %d, thirdNumber = %d\n", firstNumber, *secondNumber, thirdNumber)

	firstNumber = 10
	fmt.Printf("firstNumber = %d, secondNumber = %d, thirdNumber = %d\n", firstNumber, *secondNumber, thirdNumber)

	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
