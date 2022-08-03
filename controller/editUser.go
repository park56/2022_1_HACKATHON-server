package controller

import (
	db "2022_07_HT/database"
	"2022_07_HT/models"
	"log"

	"github.com/gin-gonic/gin"
)

func EditUser(c *gin.Context) {

	user := &models.User{}

	err := c.Bind(user)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "editUser binding error",
		})
		log.Println(err)
		return
	}

	log.Println(user.Id, user.Pw)

	err = db.DB.Omit("Pw", "Name", "IsStudent", "Cardinal").Save(user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "editUser edit error",
		})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "editUser edit error",
	})
}
