// models/SubjectRepository.go
package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SubjectRepository handles database operations related to the Subject model.
type SubjectRepository struct {
	Db *gorm.DB
}

// NewSubjectRepository creates a new instance of SubjectRepository.
func NewSubjectRepository(db *gorm.DB) *SubjectRepository {
	return &SubjectRepository{Db: db}
}

// GetSubjects retrieves all subjects from the database and returns them as JSON.
func (r *SubjectRepository) GetSubjects(c *gin.Context) {
	var subjects []Subject
	r.Db.Find(&subjects)
	c.JSON(200, subjects)
}

// GetSubject retrieves a subject from the database based on the given ID and returns it as JSON.
func (r *SubjectRepository) GetSubject(c *gin.Context) {
	id := c.Param("id")
	var subject Subject
	r.Db.First(&subject, id)
	c.JSON(200, subject)
}

// CreateSubject adds a new subject to the database and returns it as JSON.
func (r *SubjectRepository) CreateSubject(c *gin.Context) {
	var newSubject Subject
	c.BindJSON(&newSubject)
	r.Db.Create(&newSubject)
	c.JSON(200, newSubject)
}

// UpdateSubject updates an existing subject in the database and returns it as JSON.
func (r *SubjectRepository) UpdateSubject(c *gin.Context) {
	id := c.Param("id")
	var subject Subject
	r.Db.First(&subject, id)
	c.BindJSON(&subject)
	r.Db.Save(&subject)
	c.JSON(200, subject)
}

// DeleteSubject deletes a subject from the database based on the given ID and returns a JSON response.
func (r *SubjectRepository) DeleteSubject(c *gin.Context) {
	id := c.Param("id")
	var subject Subject
	r.Db.Delete(&subject, id)
	c.JSON(200, gin.H{"id" + id: "is deleted"})
}