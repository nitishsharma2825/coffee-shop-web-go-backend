package main

import (
	"coffeeshop/coffee"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	portNumber := os.Getenv("APP_PORT")

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Coffeeshop!",
		})
	})
	r.GET("/home", coffeeList)
	
	fmt.Printf("Starting the app on port %s", portNumber)
	r.Run(fmt.Sprintf(":%s", portNumber)) 
}

func coffeeList(c *gin.Context) {
	coffee, _ := coffee.GetCoffees()

  // Call the HTML method of the Context to render a template
  c.HTML(
    http.StatusOK,
    "index.html",
    gin.H{
      "list": coffee.List,
    },
  )
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Coffeeshop!",
	})
}
