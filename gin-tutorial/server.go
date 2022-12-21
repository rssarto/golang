package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rssarto/gin-tutorial/api"
	"github.com/rssarto/gin-tutorial/docs"
	_ "github.com/rssarto/gin-tutorial/docs"
	"github.com/rssarto/gin-tutorial/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	gindum "github.com/tpkeeper/gin-dump"
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {

	// // Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Pragmatic Reviews - Video API"
	docs.SwaggerInfo.Description = "Pragmatic Reviews - Yoututbe Video API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}

	setupLogOutput()

	server := gin.New()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	//server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindum.Dump())
	server.Use(gin.Recovery(), middlewares.Logger(), gindum.Dump())

	server.POST("/api/login", api.Authenticate)

	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", api.GetVideos)

		apiRoutes.POST("/videos", api.CreateVideo)

		apiRoutes.PUT("/videos/:id", api.UpdateVideo)

		apiRoutes.DELETE("/videos/:id", api.DeleteVideo)
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", api.ViewVideos)
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Run(":8080")
}
