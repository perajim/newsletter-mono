package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/perajim/controllers"
)

//CORSMiddleware ...
//CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
func main() {
	r := gin.Default()

	//Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}

	r.Use(CORSMiddleware())

	/*** initalize newsletter  ***/
	newsletter := new(controllers.NewsletterController)

	r.POST("/newsletter", newsletter.Create)
	r.POST("/shipping/newsletter/:id/:idFile", newsletter.SendNewsletter)
	r.POST("/newsletter/recipient", newsletter.AddRecipient)
	r.POST("/newsletter/file", newsletter.SaveFile)
	r.GET("/suscription/:newsletter/:email", newsletter.RemoveEmail)
	r.GET("/newsletter/:newsletter/recipients", newsletter.GetRecipients)
	r.GET("/newsletters", newsletter.GetNewsletters)

	r.Run()
}
