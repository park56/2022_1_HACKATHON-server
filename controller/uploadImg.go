package controller

import (
	"2022_07_HT/modules"
	"log"

	"github.com/gin-gonic/gin"
)

func UploadImg(c *gin.Context) {

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Upload IMG formFile Error",
		})
		log.Println(err)
		return
	}

	fileName := header.Filename

	err = modules.MakeFile(file, fileName)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Upload IMG makeFile Error",
		})
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"message": "Upload IMG success",
	})

}
