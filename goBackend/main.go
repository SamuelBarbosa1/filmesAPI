package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Filme struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Ator      string `json:"ator"`
	Descricao string `json:"descricao"`
	Ano       int    `json:"ano"`
}

var filmes []Filme
var nextID = 1

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/filme/{id}", getFilme).Methods("GET")
	router.HandleFunc("/filme", addFilme).Methods("POST")
	router.HandleFunc("/filme/{id}", updateFilme).Methods("PUT")
	router.HandleFunc("/filme/{id}", deleteFilme).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getFilme(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, filme := range filmes {
		if filme.ID == id {
			json.NewEncoder(w).Encode(filme)
			return
		}
	}

	http.Error(w, "Filme não encontrado", http.StatusNotFound)
}

func addFilme(w http.ResponseWriter, r *http.Request) {
	var filme Filme
	json.NewDecoder(r.Body).Decode(&filme)

	filme.ID = nextID
	nextID++

	filmes = append(filmes, filme)
	json.NewEncoder(w).Encode(filme)
}

func updateFilme(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var updatedFilme Filme
	json.NewDecoder(r.Body).Decode(&updatedFilme)

	for i, filme := range filmes {
		if filme.ID == id {
			filmes[i] = updatedFilme
			filmes[i].ID = id
			json.NewEncoder(w).Encode(filmes[i])
			return
		}
	}

	http.Error(w, "Filme não encontrado", http.StatusNotFound)
}

func deleteFilme(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for i, filme := range filmes {
		if filme.ID == id {
			filmes = append(filmes[:i], filmes[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Filme não encontrado", http.StatusNotFound)
}
