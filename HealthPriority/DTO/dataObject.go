package dto

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type PersonaLogin struct {
	Nombres      string
	Apellidos    string
	Cedula       int64
	Telefono     int64
	Direccion    string
	Edad         int
	Embarazo     bool
	Comorbilidad string
	Correo       string
	Password     string
}

func LeerJson() (dataJson map[string]float64) {
	fileContent, err := os.Open("DTO/comorbilidades.json")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	json.Unmarshal([]byte(byteResult), &dataJson)
	return dataJson
}

func ObtenerNivel(comorbilidad string) (nivel float64) {
	dataJson := LeerJson()
	for clave, valor := range dataJson {
		if clave == comorbilidad {
			nivel = valor
			break
		}

	}
	return nivel
}
