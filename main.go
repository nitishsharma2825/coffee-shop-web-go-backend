package main

import (
"net/http"
"github.com/gin-gonic/gin"
"coffeeshop/coffee"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Coffeeshop!",
		})
	})
	r.GET("/home", coffeeList)
	r.Run(":8080") 
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
