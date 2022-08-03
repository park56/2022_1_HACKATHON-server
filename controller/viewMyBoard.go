package controller

import (
	db "2022_07_HT/database"
	"2022_07_HT/models"
	"log"

	"github.com/gin-gonic/gin"
)

func ViewMyBoard(c *gin.Context) {

	user := &models.User{}
	c.Bind(user)

	boards := []models.Board{}

	err := db.DB.Where("userid = ?", user.Id).Find(&boards).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "view my board",
		})
		log.Println(err)
	}

	c.JSON(200, gin.H{

		"data": boards,
	})
}
