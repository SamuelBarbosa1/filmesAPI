package main

import (
	"encoding/json"
	"errors"
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
var idCounter int

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/filme", criarFilme).Methods("POST")
	router.HandleFunc("/filmes", listarFilmes).Methods("GET")
	router.HandleFunc("/filme/{id}", buscarFilme).Methods("GET")
	router.HandleFunc("/filme/{id}", atualizarFilme).Methods("PUT")
	router.HandleFunc("/filme/{id}", deletarFilme).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func criarFilme(w http.ResponseWriter, r *http.Request) {
	var novoFilme Filme
	err := json.NewDecoder(r.Body).Decode(&novoFilme)
	if err != nil {
		http.Error(w, "Erro ao processar dados", http.StatusBadRequest)
		return
	}

	// Validação simples
	if err := validarFilme(novoFilme); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idCounter++
	novoFilme.ID = idCounter
	filmes = append(filmes, novoFilme)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novoFilme)
}

func listarFilmes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filmes)
}

func buscarFilme(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, filme := range filmes {
		if filme.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(filme)
			return
		}
	}

	http.Error(w, "Filme não encontrado", http.StatusNotFound)
}

func atualizarFilme(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var filmeAtualizado Filme
	err = json.NewDecoder(r.Body).Decode(&filmeAtualizado)
	if err != nil {
		http.Error(w, "Erro ao processar dados", http.StatusBadRequest)
		return
	}

	// Validação para atualizar
	if err := validarFilme(filmeAtualizado); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, filme := range filmes {
		if filme.ID == id {
			filmeAtualizado.ID = filme.ID
			filmes[i] = filmeAtualizado
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(filmeAtualizado)
			return
		}
	}

	http.Error(w, "Filme não encontrado", http.StatusNotFound)
}

func deletarFilme(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for i, filme := range filmes {
		if filme.ID == id {
			filmes = append(filmes[:i], filmes[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Filme não encontrado", http.StatusNotFound)
}

// Validação simples para checar se todos os campos estão preenchidos
func validarFilme(filme Filme) error {
	if filme.Nome == "" {
		return errors.New("O campo 'nome' é obrigatório")
	}
	if filme.Ator == "" {
		return errors.New("O campo 'ator' é obrigatório")
	}
	if filme.Descricao == "" {
		return errors.New("O campo 'descricao' é obrigatório")
	}
	if filme.Ano == 0 {
		return errors.New("O campo 'ano' é obrigatório e deve ser maior que 0")
	}
	return nil
}
