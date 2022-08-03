package controller

import (
	db "2022_07_HT/database"
	"2022_07_HT/models"
	"log"

	"github.com/gin-gonic/gin"
)

func ViewGraduate(c *gin.Context) {

	user := []models.User{}

	err := db.DB.Find(&user, "isStudent = ?", "false").Error // user에
	if err != nil {
		c.JSON(400, gin.H{
			"message": "view graduate find error",
		})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "view graduate success",
		"id":      user,
	})

}
