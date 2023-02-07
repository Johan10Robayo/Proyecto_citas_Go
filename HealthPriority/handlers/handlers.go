package handlers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"

	dao "healthpriority.com/DAO"
	dto "healthpriority.com/DTO"
	"healthpriority.com/connection"
	"healthpriority.com/models"
)

func RegistroCliente(w http.ResponseWriter, r *http.Request) {
	var RequestJson dto.PersonaLogin
	conn := connection.GetConnection()
	err := json.NewDecoder(r.Body).Decode(&RequestJson)
	if err != nil {
		panic(err)
	}
	var datajson []byte
	var httpCode int = 200
	var data map[string]interface{}
	if RequestJson == (dto.PersonaLogin{}) {
		data = map[string]interface{}{
			"name":    "parametros incorrectos",
			"message": "los parametros no coinciden",
			"code":    500,
			"succes":  false,
		}

		bytes, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		datajson = bytes
		httpCode = 500
	} else {
		data = map[string]interface{}{
			"name":    "Registro exitoso",
			"message": "¡Enhorabuena!",
			"code":    200,
			"succes":  true,
		}
	}

	nivel := dto.ObtenerNivel(RequestJson.Comorbilidad)

	estado := models.Estado{
		Comorbilidad:       RequestJson.Comorbilidad,
		Nivel_comorbilidad: nivel,
		Embarazo:           RequestJson.Embarazo,
	}

	dao.CrearEstado(conn, estado)

	estadoIns := dao.ObtenerEstado(conn)

	persona := models.Persona{
		Cedula:    RequestJson.Cedula,
		Nombres:   RequestJson.Nombres,
		Apellidos: RequestJson.Apellidos,
		Telefono:  RequestJson.Telefono,
		Direccion: RequestJson.Direccion,
		Edad:      RequestJson.Edad,
		EstadoID:  estadoIns.Id,
		Estado:    estadoIns,
	}

	dao.CrearPersona(conn, persona)

	data2 := []byte(RequestJson.Password)
	hash := sha256.Sum256(data2)
	hashStr := fmt.Sprintf("%x", hash)

	cliente := models.Login{
		Correo:    RequestJson.Correo,
		Password:  hashStr,
		Rol:       "CLIENTE",
		PersonaID: persona.Cedula,
		Persona:   persona,
	}

	dao.CrearCliente(conn, cliente)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(datajson)
}
