package main

import (
	"log"

	"github.com/backpacker69/storasd/db"
	"github.com/backpacker69/storasd/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	r := gin.Default()

	// Offer routes
	r.GET("/offers", handlers.GetOffers)
	r.GET("/offers/:id", handlers.GetOffer)
	r.POST("/offers", handlers.CreateOffer)
	r.PUT("/offers/:id", handlers.UpdateOffer)
	r.DELETE("/offers/:id", handlers.DeleteOffer)

	// Oracle routes
	r.GET("/oracles", handlers.GetOracles)
	r.GET("/oracles/:id", handlers.GetOracle)
	r.POST("/oracles", handlers.CreateOracle)
	r.PUT("/oracles/:id", handlers.UpdateOracle)
	r.DELETE("/oracles/:id", handlers.DeleteOracle)

	// User routes
	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUser)
	r.POST("/users", handlers.CreateUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	// Message routes
	r.GET("/messages", handlers.GetMessages)
	r.GET("/messages/:id", handlers.GetMessage)
	r.POST("/messages", handlers.CreateMessage)
	r.PUT("/messages/:id", handlers.UpdateMessage)
	r.DELETE("/messages/:id", handlers.DeleteMessage)

	r.Run(":8080")
}
