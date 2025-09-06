package main

import (
	"coding/backend-api/config"
	"coding/backend-api/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Memuat variabel lingkungan dari file .env
	config.LoadEnv()

	// Inisialisasi koneksi database
	database.InitDB()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "halo Ini Dari Golang!",
		})
	})

	router.Run(": " + config.GetEnv("APP_PORT", "3000"))
}