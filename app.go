package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log_server_port := os.Getenv("LOG_SERVER_PORT")
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/", func(c *gin.Context) {
		guid := c.PostForm("guid")
		// log 폴더 존재 하지않을경우 생성
		_, err := os.Stat("./log")
		if err != nil {
			os.Mkdir("./log", os.ModePerm)
		}
		// single file
		file, _ := c.FormFile("file")
		// Upload the file to specific dst.
		c.SaveUploadedFile(file, fmt.Sprintf("./log/%v_%v", guid, file.Filename))
		// 응답
		c.String(http.StatusCreated, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(log_server_port)
}
