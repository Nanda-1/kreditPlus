package routes

import (
	"kreditPlus/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	custumerRoute := r.Group("api/customer")
	custumerRoute.POST("/create", controllers.NewCostumerRepo().CreatedCostumer)
	custumerRoute.GET("/get-all", controllers.NewCostumerRepo().Index)

	return r
}
