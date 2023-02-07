package dao

import (
	"gorm.io/gorm"
	"healthpriority.com/models"
)

func CrearPersona(conn *gorm.DB, persona models.Persona) {
	conn.Create(&persona)
}

func CrearAgenda(conn *gorm.DB, agenda models.Agenda) {
	conn.Create(&agenda)
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
func PersonaById(conn *gorm.DB, id int) (persona models.Persona) {
	conn.First(&persona, id)
	return persona
}
