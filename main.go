package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rigolo-api/services/clients"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	db := createDatabaseConnection()
	r := gin.Default()
	r.Use(CORS)
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "UP",
		})
	})
	r.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "UP",
		})
	})

	// Context bootstrapping
	clients.Bootstrap(clients.ClientsServiceConfiguration{DB: db, Server: r})

	log.Fatal(r.Run())
}
func CORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
func createDatabaseConnection() *gorm.DB {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, username, dbName, password, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		panic("failed to connect database")
	}

	return db
}
