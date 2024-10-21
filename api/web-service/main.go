package main

import (
  "net/http"

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

func main() {
  router := gin.Default()
  router.GET("/products", getProducts)

  router.Run("localhost:8080")
}
