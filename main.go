package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"Message":"Running server...","meeting link":"https://teams.microsoft.com/l/meetup-join/19%3ameeting_NTI4MjFkZDctZDUwMS00MmJhLWFlZTctNzM1N2VlZGI3MGIy%40thread.v2/0?context=%7b%22Tid%22%3a%22e6c83ca8-8322-4c7e-b01c-83474d4c1e19%22%2c%22Oid%22%3a%22e7b39be9-fa48-4c51-944b-32aee32a1e48%22%7d"})
	})
	server.Run()
}
