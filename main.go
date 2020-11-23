package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Nota ...
type Nota struct {
	Titulo      string    `json:"titulo"`
	Descripcion string    `json:"descripcion"`
	CreadaElDia time.Time `json:"creada_el_dia"`
}

var datosNotas = make(map[string]Nota)
var id int

// Main ..
func main() {
	gorillarouter := mux.NewRouter().StrictSlash(false)

	gorillarouter.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	gorillarouter.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	gorillarouter.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	gorillarouter.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8083",
		Handler:        gorillarouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Escuchando en localhost puerto 8083 ...")
	server.ListenAndServe()

}

// GetNoteHandler ...
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notas []Nota
	for _, valor := range datosNotas {
		notas = append(notas, valor)
	}
	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(notas)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)

	w.Write(j)

}

// PostNoteHandler ...
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {

	var nota Nota
	err := json.NewDecoder(r.Body).Decode(&nota)
	if err != nil {
		panic(err)
	}
	nota.CreadaElDia = time.Now()
	id++
	k := strconv.Itoa(id)
	datosNotas[k] = nota

	w.Header().Set("Content-Type", "application/json")

	j, err := json.Marshal(nota)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)

	w.Write(j)
}

// PutNoteHandler ...
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	k := vars["id"]

	var notaUpdate Nota
	err := json.NewDecoder(r.Body).Decode(&notaUpdate)
	if err != nil {
		panic(err)
	}

	if nota, ok := datosNotas[k]; ok {

		notaUpdate.CreadaElDia = nota.CreadaElDia

		delete(datosNotas, k)

		datosNotas[k] = notaUpdate
	} else {

	}

	w.WriteHeader(http.StatusNoContent)

}

// DeleteNoteHandler ...
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	k := vars["id"]

	if _, ok := datosNotas[k]; ok {
		delete(datosNotas, k)
	} else {

	}

	w.WriteHeader(http.StatusNoContent)

}
