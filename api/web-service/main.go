package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var db *sql.DB

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	connectDb()
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductByID)
	router.POST("/products", postProducts)

	router.Run(":80")
}

func connectDb() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := 5432
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open database connection
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to db ")
}

func getProductsFromDB() ([]product, error) {
	rows, err := db.Query("SELECT id, name, description FROM products")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var products []product
	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description); err != nil {
			log.Println(err)
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func getProducts(c *gin.Context) {
	if db == nil {
		log.Println("Database connection not established")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	products, err := getProductsFromDB()
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product data"})
		return
	}

	c.IndentedJSON(http.StatusOK, products)
}

func postProducts(c *gin.Context) {
	var newProduct product

	if err := c.BindJSON(&newProduct); err != nil {
		log.Println("bind json error", err)
		return
	}

	result, err := addProduct(&newProduct)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusCreated, result)
}

func addProduct(p *product) (int, error) {
	_, err := db.Exec("INSERT INTO products (id, name, description) VALUES ($1, $2, $3)", p.ID, p.Name, p.Description)
	if err != nil {
		return 0, fmt.Errorf("addProduct: %v", err)
	}

	return p.ID, nil
}

func getProductByID(c *gin.Context) {
	var p product
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	row, err := db.Query("SELECT id, name, description FROM products WHERE id = $1", idInt)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product data"})
		return
	}

	defer row.Close()

	for row.Next() {
		if err := row.Scan(&p.ID, &p.Name, &p.Description); err != nil {
			log.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product data"})
			return
		}
	}
	fmt.Println(p)

	c.IndentedJSON(http.StatusOK, p)
}
