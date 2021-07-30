package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	err       error
	db        *sql.DB
	contador  int
	contactos []Contacto
	estados   []Estado
)

type Contacto struct {
	IdContacto     int    `json:"idContacto"`
	NombreCompleto string `json:"nombreCompleto"`
	Email          string `json:"email"`
	TelMovil       string `json:"telMovil"`
	Mensaje        string `json:"mensaje"`
	ID_Estado      *Estado
}
type Estado struct {
	ID_Estado int32  `json:"idEstado"`
	Nombre    string `json:"nombreEstado"`
}

func main() {
	contactos = append(contactos, Contacto{IdContacto: 1, NombreCompleto: "Itzi Cabrera Sanchez", Email: "itzi@gmail.com", TelMovil: "5512402704", Mensaje: "Hols!", ID_Estado: &Estado{ID_Estado: 1, Nombre: "Mexico"}})

	estados = append(estados, Estado{ID_Estado: 1, Nombre: "Mexico"})
	estados = append(estados, Estado{ID_Estado: 2, Nombre: "Aguascalientes"})
	estados = append(estados, Estado{ID_Estado: 3, Nombre: "Ciudad de México"})
	estados = append(estados, Estado{ID_Estado: 4, Nombre: "Cancún"})
	estados = append(estados, Estado{ID_Estado: 5, Nombre: "Baja California Sur"})
	estados = append(estados, Estado{ID_Estado: 6, Nombre: "Baja California"})
	estados = append(estados, Estado{ID_Estado: 7, Nombre: "Campeche"})
	estados = append(estados, Estado{ID_Estado: 8, Nombre: "Veracruz"})

	router := mux.NewRouter()
	router.HandleFunc("/", LandingPage).Methods("GET")
	router.HandleFunc("/contactos", GetContactosEndPoint).Methods("GET")
	router.HandleFunc("/contacto", PostContactosEndPoint).Methods("POST")

	handlerCORS := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	}).Handler(router)

	http.ListenAndServe("localhost:3306", handlerCORS)
}
func separador() {
	fmt.Println("_______________________________________")
}
func LandingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request")
	fmt.Fprintf(w, "Hola! La peticion ha sido recibida..")
	contador += 1
	fmt.Println("Vistas: ", contador)

	json.NewEncoder(w).Encode(contador)
}

func GetContactosEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Peticion Contactos")
	json.NewEncoder(w).Encode(contactos)
}

func PostContactosEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Peticion Post")

	var nuevo Contacto
	json.NewDecoder(r.Body).Decode(&nuevo)

	contactos = append(contactos, nuevo)
	json.NewEncoder(w).Encode(contactos)

	separador()

	fmt.Println("Total de Registros: ", len(contactos))
	fmt.Println("Estados:", estados)
	separador()

}
