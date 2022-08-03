package main

import (
	"2022_07_HT/database"
	"2022_07_HT/handler"
	"2022_07_HT/middleware"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	r := gin.Default() // gin 프레임 워크

	database.Connect() // DB를 전역변수로 풀어주기

	r.Use(middleware.CORSMiddleware()) // CORS 허용

	handler.Route(r) // 라우팅

	r.Run() // 서버시작 loclahost:8080
}
