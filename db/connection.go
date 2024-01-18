package db

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func ConnectDatabase(dbType, dbUser, dbPassword, dbHost, dbPort, dbName string) (*gorm.DB, error) {
	
	var dialector gorm.Dialector
	switch dbType {
	case "mysql":
// dsn := "admin:adminpassword@tcp(localhost:3306)/go_gorm?charset=utf8mb4&parseTime=True&loc=Local"
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
		fmt.Println("dsn:", dsn)
		dialector = mysql.Open(dsn)
	case "postgres":
	case "sqlite":
	default:
		return nil, fmt.Errorf("unknown db type %s", dbType)
	}
	
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}