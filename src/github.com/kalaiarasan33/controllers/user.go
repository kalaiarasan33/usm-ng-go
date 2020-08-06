package controllers

import (
	"net/http"
	"github.com/kalaiarasan33/usm-ng-go/models"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Name   string `json:"name" binding:"required"`
	Depart string `json:"depart" binding:"required"`
}
type UpdateUserInput struct {
	Name   string `json:"name"`
	Depart string `json:"depart"`
}

// GET METHOD /users

func Findusers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// POST METHOD /createuser
func Createuser(c *gin.Context) {
	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Depart: input.Depart}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})

}

// GET METHOD /user/:id

func Finduser(c *gin.Context) {

	var user []models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}

// PATCH /updateuser/:id

func Updateuser(c *gin.Context) {

	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(&input)

	c.JSON(http.StatusOK, gin.H{"data": user})

}

// DELETE /deleteuser/:id
func DeleteUser(c *gin.Context) {

	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&user).Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})

}
