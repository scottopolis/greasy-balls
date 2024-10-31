package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type product struct {
  ID string `json:"id"`
  Name string `json:"name"`
  Price float64 `json:"price"`
}

var products = []product{
  {ID: "1", Name: "Transition Repeater", Price: 9999.99},
  {ID: "2", Name: "Titleist Pro V1", Price: 29.99},
  {ID: "3", Name: "Maxxis Assegai", Price: 102.95},
}

func getProducts(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, products)
}

func postProducts( c *gin.Context) {
  var newProduct product

  if err := c.BindJSON(&newProduct); err != nil {
    return
  }

  products = append(products, newProduct)
  c.IndentedJSON(http.StatusCreated, newProduct)
}

func getProductByID(c *gin.Context) {
	id := c.Param("id")

	for _, a:= range products {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
  router := gin.Default()
  router.Use(cors.Default())
  router.GET("/products", getProducts)
  router.GET("/products/:id", getProductByID)
  router.POST("/products", postProducts)

  router.Run(":80")
}
