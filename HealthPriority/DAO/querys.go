package dao

import (
	"gorm.io/gorm"
	"healthpriority.com/models"
)

func CrearPersona(conn *gorm.DB, persona models.Persona) {
	conn.Create(&persona)
}

func CrearCliente(conn *gorm.DB, login models.Login) {
	conn.Create(&login)
}

func CrearEstado(conn *gorm.DB, estado models.Estado) {
	conn.Create(&estado)
}

func ObtenerEstado(conn *gorm.DB) (estado models.Estado) {
	conn.Last(&estado)
	return estado
}
