package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/", func(c *gin.Context) {
		// log 폴더 존재 하지않을경우 생성
		_, err := os.Stat("./log")
		fmt.Println(err)
		if err != nil {
			os.Mkdir("./log", os.ModePerm)
		}
		// single file
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "./log/"+file.Filename)

		c.String(http.StatusCreated, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	router.Run(":9090")
}
