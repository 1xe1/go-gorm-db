// models/Item.go
package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name  string
	Price float64
}
