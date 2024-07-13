package routes

import (
	"github.com/gin-gonic/gin"
	docs "github.com/odeassis/goodoor/docs"
	"github.com/odeassis/goodoor/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	handler.Init()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group("/api/v1")

	v1.GET("/opening", handler.ShowOpening)

	v1.POST("/opening", handler.CreateOpening)

	v1.DELETE("/opening", handler.DeleteOpening)

	v1.PUT("/opening", handler.UpdateOpening)

	v1.GET("/openings", handler.ListOpenings)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
