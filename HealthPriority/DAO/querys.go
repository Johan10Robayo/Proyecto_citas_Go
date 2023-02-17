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
func PersonaById(conn *gorm.DB, id int64) (persona models.Persona) {
	conn.First(&persona, id)
	return persona
}

func CrearAutorizacion(conn *gorm.DB, autorizacion models.Autorizacion) {
	conn.Create(&autorizacion)
}
func GetAutorizaciones(conn *gorm.DB, id int64) (autorizaciones []models.Autorizacion) {
	conn.Where("persona_id = ?", id).Find(&autorizaciones)
	return autorizaciones
}
func VerificarUsuario(conn *gorm.DB, correo string, password string) (usuario models.Login) {
	conn.Where("correo = ? AND password >= ?", correo, password).Find(&usuario)
	return usuario
}
