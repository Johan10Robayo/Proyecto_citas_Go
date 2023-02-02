package connection

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?"+
		"charset:utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database...")
	}
	fmt.Println("Conectado exitosamente a la bd")

	return db
}
