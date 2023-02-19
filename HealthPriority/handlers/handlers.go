package handlers

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
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
			"name":    "Parámetros incorrectos",
			"message": "Los parámetros no coinciden",
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
		bytes, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		datajson = bytes
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

func AgendarMedicoG(w http.ResponseWriter, r *http.Request) {
	var RequestJson dto.Agenda
	conn := connection.GetConnection()
	err := json.NewDecoder(r.Body).Decode(&RequestJson)
	if err != nil {
		panic(err)
	}
	var datajson []byte
	var httpCode int = 200
	var data map[string]interface{}
	if RequestJson == (dto.Agenda{}) {
		data = map[string]interface{}{
			"name":    "Parámetros incorrectos",
			"message": "Los parámetros no coinciden",
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
			"name":    "Agendamiento exitoso",
			"message": "¡Enhorabuena!",
			"code":    200,
			"succes":  true,
		}
		bytes, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		datajson = bytes
	}

	persona := dao.PersonaById(conn, RequestJson.Cedula)

	// Convertir cadena a tiempo

	fecha, err := time.Parse(time.RFC3339, RequestJson.Fecha+"T12:00:00.511Z")
	if err != nil {
		fmt.Println(err)
		return
	}
	agenda := models.Agenda{
		Fecha:     fecha,
		Jornada:   RequestJson.Jornada,
		Tipo:      RequestJson.Tipo,
		PersonaID: persona.Cedula,
		Persona:   persona,
	}

	dao.CrearAgenda(conn, agenda)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(datajson)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var RequestJson dto.Login
	var token *jwt.Token
	var datajson []byte
	var httpCode int
	var data map[string]interface{}
	conn := connection.GetConnection()
	err := json.NewDecoder(r.Body).Decode(&RequestJson)
	if err != nil {
		panic(err)
	}

	sha := []byte(RequestJson.Password)
	hash := sha256.Sum256(sha)
	hashStr := fmt.Sprintf("%x", hash)
	usuario := dao.VerificarUsuario(conn, RequestJson.User, hashStr)

	if usuario == (models.Login{}) {
		data = map[string]interface{}{
			"name":    "usuario no encontrado, error al autenticar",
			"message": "Usuario o contraseña incorrecta, error al autenticar",
			"code":    http.StatusForbidden,
			"succes":  false,
		}
		bytes, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		httpCode = http.StatusForbidden
		datajson = bytes
	} else {
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       usuario.PersonaID,
			"username": usuario.Correo,
			"role":     usuario.Rol,
		})
		tokenString, err := token.SignedString([]byte("cb97baeaab7da33a8c6ca9b9165261ce05e9554533bcbab9389489f9c8d9f1d6"))
		if err != nil {
			panic(err)
		}
		data = map[string]interface{}{
			"name":    "Autenticación exitosa",
			"message": tokenString,
			"code":    http.StatusOK,
			"succes":  true,
		}
		bytes, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		httpCode = http.StatusOK
		datajson = bytes

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(datajson)
}

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if true {
			next.ServeHTTP(w, r)
			return
		}

		authorization := r.Header.Get("Authorization")
		fmt.Println("header", authorization)
		stringArray := strings.Split(authorization, " ")
		fmt.Println("tokem", stringArray)
		tokenString := stringArray[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("cb97baeaab7da33a8c6ca9b9165261ce05e9554533bcbab9389489f9c8d9f1d6"), nil
		})
		if err == nil && token.Valid {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
		}
	})
}

func Autorizacion(w http.ResponseWriter, r *http.Request) {
	var RequestJson dto.DatosAutorizacion
	conn := connection.GetConnection()
	err := json.NewDecoder(r.Body).Decode(&RequestJson)
	if err != nil {
		panic(err)
	}
	var datajson []byte
	var httpCode int = 200
	var data map[string]interface{}
	//fmt.Println(RequestJson)
	if RequestJson == (dto.DatosAutorizacion{}) {
		data = map[string]interface{}{
			"name":    "Parámetros incorrectos",
			"message": "Los parámetros no coinciden",
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

		imageBytes, err := base64.StdEncoding.DecodeString(RequestJson.Imagen)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Especificar la ruta de destino del archivo
		filePath := "C:/Users/57322/Documents/Universidad Johan/Semestre 10/Go/Proyecto_final/Proyecto_citas_Go/HealthPriority/statics/archivos/documentos/" + RequestJson.NombreImagen

		err3 := ioutil.WriteFile(filePath, imageBytes, 0644)

		if err3 != nil {
			fmt.Println(err3)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fechaV := time.Now().AddDate(0, 0, 30)
		currentTime := time.Now()
		person := dao.PersonaById(conn, RequestJson.Person_id)
		autorizacion := models.Autorizacion{
			Nombre:            RequestJson.Nombre,
			Tipo:              RequestJson.Tipo,
			Url:               filePath,
			Estado:            RequestJson.Estado,
			Fecha_creacion:    currentTime,
			Fecha_vencimiento: fechaV,
			PersonaID:         RequestJson.Person_id,
			Persona:           person,
		}
		dao.CrearAutorizacion(conn, autorizacion)

		fmt.Fprintf(w, "Imagen subida exitosamente: %s", RequestJson.NombreImagen)
		httpCode = http.StatusOK
		data = map[string]interface{}{
			"name":    "Autorización cargada",
			"message": "¡Enhorabuena!",
			"code":    200,
			"succes":  true,
		}
		bytes, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		datajson = bytes

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(datajson)
}

func GetAutorizaciones(w http.ResponseWriter, r *http.Request) {
	var RequestJson dto.PersonaId
	conn := connection.GetConnection()
	err := json.NewDecoder(r.Body).Decode(&RequestJson)
	if err != nil {
		panic(err)
	}
	var datajson []byte
	var httpCode int = 200
	var data map[string]interface{}
	//fmt.Println(RequestJson)
	if RequestJson == (dto.PersonaId{}) {
		data = map[string]interface{}{
			"name":    "Parámetros incorrectos",
			"message": "Los parámetros no coinciden",
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
		fmt.Println(RequestJson.Persona_id)
		autorizaciones := dao.GetAutorizaciones(conn, RequestJson.Persona_id)

		httpCode = http.StatusOK
		data = map[string]interface{}{
			"name":    "Autorizaciones obtenidas",
			"message": "¡Enhorabuena!",
			"code":    200,
			"succes":  true,
		}
		bytes, err := json.Marshal(autorizaciones)
		if err != nil {
			panic(err)
		}
		datajson = bytes

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(datajson)
}

func AgendarEsp(w http.ResponseWriter, r *http.Request) {
	var RequestJson dto.Agenda
	conn := connection.GetConnection()
	err := json.NewDecoder(r.Body).Decode(&RequestJson)
	if err != nil {
		panic(err)
	}
	var datajson []byte
	var httpCode int = 200
	var data map[string]interface{}
	if RequestJson == (dto.Agenda{}) {
		data = map[string]interface{}{
			"name":    "Parámetros incorrectos",
			"message": "Los parámetros no coinciden",
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
			"name":    "Agendamiento exitoso",
			"message": "¡Enhorabuena!",
			"code":    200,
			"succes":  true,
		}
		bytes, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		datajson = bytes
	}

	persona := dao.PersonaById(conn, RequestJson.Cedula)

	// Convertir cadena a tiempo

	fecha, err := time.Parse(time.RFC3339, RequestJson.Fecha+"T12:00:00.511Z")
	if err != nil {
		fmt.Println(err)
		return
	}
	agenda := models.Agenda{
		Fecha:     fecha,
		Jornada:   RequestJson.Jornada,
		Tipo:      RequestJson.Tipo,
		PersonaID: persona.Cedula,
		Persona:   persona,
	}

	dao.CrearAgenda(conn, agenda)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(datajson)
}

func GetAgendas(w http.ResponseWriter, r *http.Request) {
	var RequestJson dto.PersonaId
	conn := connection.GetConnection()
	err := json.NewDecoder(r.Body).Decode(&RequestJson)
	if err != nil {
		panic(err)
	}
	var datajson []byte
	var httpCode int = 200
	var data map[string]interface{}
	//fmt.Println(RequestJson)
	if RequestJson == (dto.PersonaId{}) {
		data = map[string]interface{}{
			"name":    "Parámetros incorrectos",
			"message": "Los parámetros no coinciden",
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
		fmt.Println(RequestJson.Persona_id)
		agendas := dao.GetAgenda(conn, RequestJson.Persona_id)

		httpCode = http.StatusOK
		data = map[string]interface{}{
			"name":    "Agendas obtenidas",
			"message": "¡Enhorabuena!",
			"code":    200,
			"succes":  true,
		}
		bytes, err := json.Marshal(agendas)
		if err != nil {
			panic(err)
		}
		datajson = bytes

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(datajson)
}

func InfoPersonal(w http.ResponseWriter, r *http.Request) {
	var RequestJson dto.PersonaId
	conn := connection.GetConnection()
	err := json.NewDecoder(r.Body).Decode(&RequestJson)
	if err != nil {
		panic(err)
	}
	var datajson []byte
	var httpCode int = 200
	var data map[string]interface{}
	//fmt.Println(RequestJson)
	if RequestJson == (dto.PersonaId{}) {
		data = map[string]interface{}{
			"name":    "Parámetros incorrectos",
			"message": "Los parámetros no coinciden",
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
		fmt.Println(RequestJson.Persona_id)

		httpCode = http.StatusOK
		info := dao.PersonaById(conn, RequestJson.Persona_id)

		bytes, err := json.Marshal(info)
		if err != nil {
			panic(err)
		}
		datajson = bytes

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(datajson)
}
