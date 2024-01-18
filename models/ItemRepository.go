// models/ItemRepository.go
package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemRepository struct {
	Db *gorm.DB
}

// ทำหน้าที่สร้าง Instance ของ ItemRepository และส่งคืนกลับไป ให้เรียกใช้งานได้
func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{Db: db}
}

// ทำหน้าที่ดึงข้อมูล Item ทั้งหมดจากฐานข้อมูล และส่งกลับไปให้ผู้ใช้งาน
// (r *ItemRepository) คือการระบุว่าเป็น method ของ ItemRepository
// r คือตัวแปรที่เป็น pointer ของ ItemRepository
// ฟังก์ชัน GetItems รับพารามิเตอร์เป็น c *gin.Context และมีการส่งกลับค่าเป็น JSON กลับไปให้ผู้ใช้งานผ่าน c.JSON(200, items) ซึ่ง c คือตัวแปรที่เป็น pointer ของ gin.Context
// ในกรณีนี้เราไม่ต้องใช้ return เพราะเราใช้ c.JSON แทน
// ฟังก์ชันนี้ parameter ที่รับมาจะเป็น pointer ของ gin.Context เพราะเราจะใช้ c.JSON ส่งค่ากลับไปให้ผู้ใช้งาน
func (r *ItemRepository) GetItems(c *gin.Context) {
	var items []Item
	r.Db.Find(&items)
	c.JSON(200, items)
}

// ทำหน้าที่เพิ่มข้อมูล Item ลงในฐานข้อมูล และส่งกลับไปให้ผู้ใช้งานผ่าน c.JSON(200, newItem)
func (r *ItemRepository) PostItem(c *gin.Context) {
	var newItem Item
	c.BindJSON(&newItem)
	r.Db.Create(&newItem)
	c.JSON(200, newItem)
}

// ฟังก์ชันค้นหา Item จากฐานข้อมูล โดยใช้ id ที่รับเข้ามาเป็นเงื่อนไขในการค้นหา
// ฟังก์ชันนี้ parameter ที่รับมาจะเป็น pointer ของ gin.Context เพราะเราจะใช้ c.JSON ส่งค่ากลับไปให้ผู้ใช้งาน
func (r *ItemRepository) GetItem(c *gin.Context) {
	id := c.Param("id")
	var item Item
	r.Db.First(&item, id)
	c.JSON(200, item)
}

// ฟังก์ชันอัพเดทข้อมูล Item ลงในฐานข้อมูล และส่งกลับไปให้ผู้ใช้งานผ่าน c.JSON(200, item)
func (r *ItemRepository) UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var item Item
	r.Db.First(&item, id)
	c.BindJSON(&item)
	r.Db.Save(&item)
	c.JSON(200, item)
}

// ฟังก์ชันลบข้อมูล Item ออกจากฐานข้อมูล และส่งกลับไปให้ผู้ใช้งานผ่าน c.JSON(200, gin.H{"id" + id: "is deleted"})
// gin.H ทำหน้าที่สร้าง map ของ key และ value และส่งกลับไปให้ผู้ใช้งาน
// H คือตัวย่อของ Header ใน HTTP
func (r *ItemRepository) DeleteItem(c *gin.Context) {
	id := c.Param("id")
	var item Item
	r.Db.Delete(&item, id)
	c.JSON(200, gin.H{"id" + id: "is deleted"})
}