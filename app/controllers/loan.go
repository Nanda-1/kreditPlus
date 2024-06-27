package controllers

import (
	"kreditPlus/app/database"
	"kreditPlus/app/models"
	"kreditPlus/app/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoanRepo struct {
	Db *gorm.DB
}

func NewLoanRepo() *LoanRepo {
	db := database.DB

	return &LoanRepo{Db: db}
}

func (l *LoanRepo) CreateLoan(c *gin.Context) {
	res := models.Reqjson{Success: true}

	type LoanReq struct {
		models.Loan
		TenorID int `json:"tenor_id"`
	}
	req := LoanReq{}
	err := c.ShouldBind(&req)
	if err != nil {
		res.Success = false
		res.Error = err.Error()
		c.JSON(400, res)
		return
	}

	store := repositories.CreateLoan(&req.Loan, req.TenorID)
	if store != nil {
		res.Success = false
		res.Error = store.Error()
		c.JSON(400, res)
		return
	}

	res.Data = store
	c.JSON(200, res)
}
