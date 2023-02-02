package models

import (
	"fmt"

	"gorm.io/gorm"
)

func Migrate(conn *gorm.DB) {
	if !(conn.Migrator().HasTable(&Estado{})) {
		conn.AutoMigrate(&Estado{})
		fmt.Println("La tabla estado ha sido creada")
	}

	if !(conn.Migrator().HasTable(&Persona{})) {
		conn.AutoMigrate(&Persona{})
		fmt.Println("La tabla persona ha sido creada")
	}

	if !(conn.Migrator().HasTable(&Login{})) {
		conn.AutoMigrate(&Login{})
		fmt.Println("La tabla login ha sido creada")
	}

	if !(conn.Migrator().HasTable(&Cita{})) {
		conn.AutoMigrate(&Cita{})
		fmt.Println("La tabla cita ha sido creada")
	}

	if !(conn.Migrator().HasTable(&Agenda{})) {
		conn.AutoMigrate(&Agenda{})
		fmt.Println("La tabla agenda ha sido creada")
	}

	if !(conn.Migrator().HasTable(&Autorizacion{})) {
		conn.AutoMigrate(&Autorizacion{})
		fmt.Println("La tabla autorización ha sido creada")
	}
}
