package main

import (
	"web-microservice-starter/src/controller"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @title Articles Microservice
// @version 0.0.2
// @description This microservice serves as an example microservice

// @BasePath /webmicroservicestarter/api/v1
func setupRoutes(r *gin.Engine) {
	// Instantiate controllers
	articleController := controller.ArticleController{}

	// Set application context in URL - Do not edit this
	application := r.Group(viper.GetString("server.basepath"))
	{
		// Edit from here
		// All routes for API version V1
		v1 := application.Group("/api/v1")
		{
			docs := v1.Group("/docs")
			{
				docs.StaticFile("/swagger.json", "./docs/swagger.json")
				docs.StaticFile("/swagger.yaml", "./docs/swagger.yaml")
			}

			articles := v1.Group("/articles")
			{
				articles.GET(controller.GetArticlesPath, articleController.GetArticles)
				articles.GET(controller.GetArticlePath, articleController.GetArticle)
			}
		}
	}
}
