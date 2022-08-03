package controller

import (
	db "2022_07_HT/database"
	"2022_07_HT/models"
	"log"

	"github.com/gin-gonic/gin"
)

func UserData(c *gin.Context) {

	user := &models.User{}
	id := c.Query("id") // 쿼리 값 중 id 읽어오기

	err := db.DB.Find(user, "id = ?", id).Error // id가 일치하는 row 가져오기
	if err != nil {
		c.JSON(400, gin.H{
			"message": "user update finding error",
		})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "user data success",
		"data":    user,
	})
}
