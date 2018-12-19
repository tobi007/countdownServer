package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tobi007/cdServer/controllers"
	"github.com/tobi007/cdServer/db"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)
	//router.Use(middlewares.AuthMiddleware())

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := controllers.NewUserController(db.GetDB())
			userGroup.GET("/:id", user.Retrieve)
		}

		eventGroup := v1.Group("/event")
		{
			event := controllers.NewEventController(db.GetDB())
			eventGroup.GET("/:id", event.GetEventByUserId)
			eventGroup.GET("", event.GetAllEvent)
			eventGroup.POST("", event.NewEvent)
			eventGroup.DELETE("/:id", event.DeleteEvent)
		}
	}
	return router

}
