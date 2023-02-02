package models

import "time"

type Estado struct {
	Id                 int64  `gorm:"primaryKey;autoIncrement:true"`
	Comorbilidad       string `gorm:"type:varchar(50)"`
	Nivel_comorbilidad float64
	Embarazo           bool
}

type Persona struct {
	Cedula    int64  `gorm:"primaryKey;autoIncrement:false"`
	Nombres   string `gorm:"type:varchar(40)"`
	Apellidos string `gorm:"type:varchar(50)"`
	Telefono  int64
	Direccion string `gorm:"type:varchar(50)"`
	Edad      int    `gorm:"type:tinyint"`
	EstadoID  int64
	Estado    Estado `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Login struct {
	Correo    string `gorm:"primaryKey;type:varchar(50)"`
	Password  string `gorm:"type:varchar(100)"`
	Rol       string `gorm:"type:varchar(50)"`
	PersonaID int64
	Persona   Persona `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Cita struct {
	Id         int64     `gorm:"primaryKey;autoIncrement:true"`
	Fecha_hora time.Time `gorm:"type:timestamp"`
	Tipo       string    `gorm:"type:varchar(20)"`
	PersonaID  int64
	Persona    Persona `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Agenda struct {
	Id        int64     `gorm:"primaryKey;autoIncrement:true"`
	Fecha     time.Time `gorm:"type:timestamp"`
	Jornada   string    `gorm:"type:varchar(15)"`
	Tipo      string    `gorm:"type:varchar(20)"`
	PersonaID int64
	Persona   Persona `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Autorizacion struct {
	Id                int64  `gorm:"primaryKey;autoIncrement:true"`
	Nombre            string `gorm:"type:varchar(20)"`
	Tipo              string `gorm:"type:varchar(20)"`
	Url               string `gorm:"type:varchar(200)"`
	Estado            bool
	Fecha_creacion    time.Time `gorm:"type:timestamp"`
	Fecha_vencimiento time.Time `gorm:"type:timestamp"`
	PersonaID         int64
	Persona           Persona `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
