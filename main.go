// main.go
package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/1xe1/go-gorm-db/models"
	"github.com/1xe1/go-gorm-db/db"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Read database configuration from environment variables
	dbType := os.Getenv("DB_TYPE")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Connect to the database
	database, err := db.ConnectDatabase(dbType, dbUser, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// AutoMigrate the database
	err = database.AutoMigrate(&models.Item{}, &models.Student{}, &models.Subject{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Create repositories for each model
	itemRepo := models.NewItemRepository(database)
	studentRepo := models.NewStudentRepository(database)
	subjectRepo := models.NewSubjectRepository(database)

	// Initialize Gin router
	r := gin.Default()

	// Item routes
	r.GET("/items", itemRepo.GetItems)
	r.POST("/items", itemRepo.PostItem)
	r.GET("/items/:id", itemRepo.GetItem)
	r.PUT("/items/:id", itemRepo.UpdateItem)
	r.DELETE("/items/:id", itemRepo.DeleteItem)

	// Student routes
	r.GET("/students", studentRepo.GetStudents)
	r.POST("/students", studentRepo.CreateStudent)
	r.GET("/students/:id", studentRepo.GetStudent)
	r.PUT("/students/:id", studentRepo.UpdateStudent)
	r.DELETE("/students/:id", studentRepo.DeleteStudent)

	// Subject routes
	r.GET("/subjects", subjectRepo.GetSubjects)
	r.POST("/subjects", subjectRepo.CreateSubject)
	r.GET("/subjects/:id", subjectRepo.GetSubject)
	r.PUT("/subjects/:id", subjectRepo.UpdateSubject)
	r.DELETE("/subjects/:id", subjectRepo.DeleteSubject)

	// 404 route
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	// Run the server
	if err := r.Run(":5000"); err != nil {
		log.Fatalf("Server is not running: %v", err)
	}
}