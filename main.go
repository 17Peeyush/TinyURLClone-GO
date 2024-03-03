package main

import (
	"net/http"

	"github.com/17Peeyush/TinyURLClone-GO/db"
	"github.com/17Peeyush/TinyURLClone-GO/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	// m :=models.Meet{
	// 	UniqueID: "a1OXylP",
	// 	Url: "https://teams.microsoft.com/l/meetup-join/19%3ameeting_NTI4MjFkZDctZDUwMS00MmJhLWFlZTctNzM1N2VlZGI3MGIy%40thread.v2/0?context=%7b%22Tid%22%3a%22e6c83ca8-8322-4c7e-b01c-83474d4c1e19%22%2c%22Oid%22%3a%22e7b39be9-fa48-4c51-944b-32aee32a1e48%22%7d",
	// }
	// db.InsertData(m)
	server := gin.Default()

	// Load HTML templates from the templates directory
    server.LoadHTMLGlob("templates/*")
	
	server.GET("/",func(c *gin.Context){
		//Render the index.html template with gynamic data
		c.HTML(http.StatusOK, "index.html", gin.H{"message":"first page"})
	})
	// CORS middleware configuration
	server.Use(func(c *gin.Context) {
		// Allow requests from any origin
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow certain HTTP methods
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// Allow certain HTTP headers
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		// Allow credentials
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})
	server.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"Message":"Running server...","meeting link":"https://teams.microsoft.com/l/meetup-join/19%3ameeting_NTI4MjFkZDctZDUwMS00MmJhLWFlZTctNzM1N2VlZGI3MGIy%40thread.v2/0?context=%7b%22Tid%22%3a%22e6c83ca8-8322-4c7e-b01c-83474d4c1e19%22%2c%22Oid%22%3a%22e7b39be9-fa48-4c51-944b-32aee32a1e48%22%7d"})
	})
	routes.RegisterRoutes(server)
	server.Run()
}
