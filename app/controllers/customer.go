package controllers

import (
	"kreditPlus/app/database"
	"kreditPlus/app/models"
	"kreditPlus/app/repositories"
	"log"
	"mime/multipart"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CostumerRepo struct {
	Db *gorm.DB
}

func NewCostumerRepo() *CostumerRepo {
	db := database.DB

	return &CostumerRepo{Db: db}
}

func (co *CostumerRepo) CreatedCostumer(c *gin.Context) {
	res := models.Reqjson{Success: true}
	Nik := c.PostForm("nik")
	FullName := c.PostForm("full_name")
	LegalName := c.PostForm("legal_name")
	PlaceBirthDate := c.PostForm("birth_place")
	DateBirthDate := c.PostForm("birth_date")
	Salary := c.PostForm("salary")
	salaryfloat, _ := strconv.ParseFloat(Salary, 64)
	ImgKtp, err := c.FormFile("photo_ktp")
	if err != nil {
		res.Success = false
		res.Error = "KTP image is required"
		c.JSON(400, res)
		return
	}
	ImgSelfie, err := c.FormFile("photo_selfie")
	if err != nil {
		res.Success = false
		res.Error = "Selfie image is required"
		c.JSON(400, res)
		return
	}

	if Nik == "" {
		res.Success = false
		res.Error = "Nik is required"
		c.JSON(400, res)
		return
	}

	log.Println(Nik)
	log.Println(ImgKtp.Filename)

	ktpPath, err := saveFile(ImgKtp, "uploads/ktp")
	if err != nil {
		res.Success = false
		res.Error = "Failed to save KTP image"
		c.JSON(500, res)
		return
	}

	selfiePath, err := saveFile(ImgSelfie, "uploads/selfie")
	if err != nil {
		res.Success = false
		res.Error = "Failed to save selfie image"
		c.JSON(500, res)
		return
	}

	store := &models.Customer{
		NIK:        Nik,
		FullName:   FullName,
		LegalName:  LegalName,
		BirthPlace: PlaceBirthDate,
		BirthDate:  DateBirthDate,
		Salary:     salaryfloat,
		IDPhoto:    ktpPath,
		Selfie:     selfiePath,
	}

	insert, err := repositories.StoreCostumer(store)
	if err != nil {
		res.Success = false
		res.Error = err.Error()
		c.JSON(400, res)
		return
	}

	res.Data = insert
	c.JSON(200, res)
}
func saveFile(file *multipart.FileHeader, folder string) (string, error) {
	var c *gin.Context
	filename := filepath.Base(file.Filename)
	path := filepath.Join(folder, filename)

	if err := c.SaveUploadedFile(file, path); err != nil {
		return "", err
	}

	return path, nil
}

func (co *CostumerRepo) Index(c *gin.Context) {
	res := models.Reqjson{Success: true}

	data, err := repositories.GetAllUsers()
	if err != nil {
		res.Success = false
		res.Error = err.Error()
		c.JSON(400, res)
		return
	}

	res.Data = data
	c.JSON(200, res)
}
