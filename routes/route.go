package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/odeassis/goodoor/handler"
)

func initializeRoutes(routes *gin.Engine) {

	handler.Init()

	v1 := routes.Group("/api/v1")

	v1.GET("/opening", handler.ShowOpening)

	v1.POST("/opening", handler.CreateOpening)

	v1.DELETE("/opening", handler.DeleteOpening)

	v1.PUT("/opening", handler.UpdateOpening)

	v1.GET("/openings", handler.ListOpenings)

}
