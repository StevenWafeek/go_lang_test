package routes

import (
	"OnlineServer/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRoutes(server *gin.Engine) {
	server.GET("/events", handelEvent)
	server.GET("/events/:id", getEvent)

	authindcated := server.Group("/")
	authindcated.Use(middlewares.Authraization)
	authindcated.POST("/events", PostEvent)
	authindcated.PUT("/events/:id", UpdateEvent)
	authindcated.DELETE("/events/:id", DeleteEvent)

	server.POST("/users", signUp)
	server.POST("/login", logIn)
}
