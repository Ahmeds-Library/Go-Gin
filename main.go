package main

import (
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "File not found"})
			return
		}
		savePath := filepath.Join("uploads", file.Filename)
		err = c.SaveUploadedFile(file, savePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
			return
		}
	})

	r.Run(":8080")
}

func Analyze()  {
	
}