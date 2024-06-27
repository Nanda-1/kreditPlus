package routes

import (
	"kreditPlus/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	customerRoute := r.Group("api/customers")
	{
		customerRoute.POST("/create", controllers.NewCostumerRepo().CreatedCostumer)
		customerRoute.GET("/get-all", controllers.NewCostumerRepo().Index)
	}

	loanRoute := r.Group("api/loans")
	{
		loanRoute.POST("/create", controllers.NewLoanRepo().CreateLoan)
	}

	return r
}
