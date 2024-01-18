// models/StudentRepository.go
package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// StudentRepository handles database operations related to the Student model.
type StudentRepository struct {
	Db *gorm.DB
}

// NewStudentRepository creates a new instance of StudentRepository.
func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{Db: db}
}

// GetStudents retrieves all students from the database and returns them as JSON.
func (r *StudentRepository) GetStudents(c *gin.Context) {
	var students []Student
	r.Db.Find(&students)
	c.JSON(200, students)
}

// GetStudent retrieves a student from the database based on the given ID and returns it as JSON.
func (r *StudentRepository) GetStudent(c *gin.Context) {
	id := c.Param("id")
	var student Student
	r.Db.First(&student, id)
	c.JSON(200, student)
}

// CreateStudent adds a new student to the database and returns it as JSON.
func (r *StudentRepository) CreateStudent(c *gin.Context) {
	var newStudent Student
	c.BindJSON(&newStudent)
	r.Db.Create(&newStudent)
	c.JSON(200, newStudent)
}

// UpdateStudent updates an existing student in the database and returns it as JSON.
func (r *StudentRepository) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student Student
	r.Db.First(&student, id)
	c.BindJSON(&student)
	r.Db.Save(&student)
	c.JSON(200, student)
}

// DeleteStudent deletes a student from the database based on the given ID and returns a JSON response.
func (r *StudentRepository) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	var student Student
	r.Db.Delete(&student, id)
	c.JSON(200, gin.H{"id" + id: "is deleted"})
}