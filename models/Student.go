// models/Student.go
package models

import "gorm.io/gorm"

// Student represents the structure of your "students" table in the database.
type Student struct {
    gorm.Model // Gorm's default model that includes fields like ID, CreatedAt, UpdatedAt, and DeletedAt

    // Your custom fields
    FirstName string
    LastName  string
    Age       int
    Grade     string
}

// You can add more fields and methods based on your specific requirements.