package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type ejemploEstructura struct {
	Nombre string
}

func main() {
	http.HandleFunc("/", funcionEjemplo)
	http.HandleFunc("/respuesta", probarServicio)
	fmt.Println("Frontend escuchando en 7000")
	http.ListenAndServe(":7000", nil)
}

func funcionEjemplo(x http.ResponseWriter, r *http.Request) {

	u := template.Must(template.ParseFiles("plantilla.html"))
	u.Execute(x, nil)
}

func probarServicio(x http.ResponseWriter, s *http.Request) {

	cuerpo := s.FormValue("entrada")

	datosJson := map[string]string{"Nombre": cuerpo}
	valorJson, _ := json.Marshal(datosJson)
	respuesta, error := http.Post("http://localhost:2000/ejemploServicio/", "application/json", bytes.NewBuffer(valorJson))
	fmt.Println(error)
	datos, _ := ioutil.ReadAll(respuesta.Body)
	var a ejemploEstructura
	json.Unmarshal(datos, &a)
	u, _ := template.ParseFiles("plantillaRespuesta.html")

	u.Execute(x, a)

}
