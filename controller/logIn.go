package controller

import (
	db "2022_07_HT/database"
	"2022_07_HT/helper"
	"2022_07_HT/models"
	"log"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	user := &models.User{} // 회원 가입용 객체 생성

	err := c.Bind(user) // 프론트의 데이터 바인딩
	if err != nil {
		c.JSON(400, gin.H{
			"message": "login binding error",
		})
		log.Println(err)
		return
	}

	inputPw := user.Pw                            // 유저 pw를 따로 저장
	result := db.DB.Find(user, "id = ?", user.Id) // User객체에 id가 동일한 db의 row를 저장

	if result.RowsAffected == 0 { // id가 존재하지 않을 때
		c.JSON(400, gin.H{
			"message": "id isn't exist",
		})
		log.Println(err)
		return
	}

	res := helper.CheckPasswordHash(user.Pw, inputPw) // pw가 맞는지 체크
	if !res {
		c.JSON(400, gin.H{
			"message": "pw is not",
		})
		log.Println(err)
		return
	}

	accessToken, err := helper.CreateJWT(user.Id) // token 제작
	if err != nil {
		c.JSON(400, gin.H{
			"message": "token issue error",
		})
		log.Println(err)
		return
	}

	c.SetCookie("access-token", accessToken, 60*60*24, "/", "localhost", false, false) // local에 토큰 저장

	c.JSON(200, gin.H{ // 클라이언트로 id, token 보내줌
		"message": "login success",
		"id":      user.Id,
		"name":    user.Name,
		"token":   accessToken,
	})
}
