package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

func productsHandler(c *gin.Context) {
	products := []Product{
		{100, "BassTune Headset 2.0", 200, "A headphone with a inbuilt high-quality microphone"},
		{101, "Fastlane Toy Car", 100, "A toy car that comes with a free HD camera"},
		{102, "ATV Gear Mouse", 75, "A high-quality mouse for office work and gaming"},
	}
	c.JSON(200, gin.H{
		"products": products,
	})
}
func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/products", productsHandler)
	r.Run(":50000")
}
