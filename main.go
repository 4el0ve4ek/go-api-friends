package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-api-friends/controller"
	"log"
	"net/http"
	"os"
)

// fun
func homeLink(c *gin.Context) {
	c.JSON(200, "Welcome home!")
}

func init() {
	// Load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	router := gin.New()
	server := controller.NewPhoneServer()
	router.GET("/", homeLink)
	router.GET("/user", server.GetUsers)
	router.GET("/user/:id", server.GetUserById)
	router.POST("/user", server.AddUser)
	router.GET("/city/:city", server.GetUserFromCity)

	router.POST("/login", server.LoginHandler)

	authorized := router.Group("/")
	authorized.GET("/refresh_token", server.RefreshHandler)
	authorized.Use(server.AuthMiddleWare())
	{
		authorized.POST("/city", server.ChangeCity)
		authorized.POST("/status", server.ChangeStatus)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}
