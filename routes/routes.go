package routes

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()

	// Ininitialize routes
	initializeRoutes(r)

	// listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
