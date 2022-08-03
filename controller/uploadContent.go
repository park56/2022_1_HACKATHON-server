package controller

import (
	db "2022_07_HT/database"
	"2022_07_HT/models"
	"log"

	"github.com/gin-gonic/gin"
)

func UploadContent(c *gin.Context) {

	board := &models.Board{}

	err := c.Bind(board)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "upload content binding error",
		})
		log.Println(err)
		return
	}

	err = db.DB.Create(board).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "upload content upload error",
		})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "upload content success",
	})

}
