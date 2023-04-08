package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	entity "kayn.ooo/api/src/Entity"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:password@tcp(localhost)/kayn.ooo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate the schema
	DB.AutoMigrate(&entity.User{})
}
