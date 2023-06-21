package main

import (
	"encoding/json"
	"net/http"
)

type Buku struct {
	ID        int    `json:"id"`
	Judul     string `json:"judul"`
	Pengarang string `json:"pengarang"`
}

var perpustakan []Buku

func GetBooks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(perpustakan)
}

func CreateBooks(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var params Buku
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	perpustakan = append(perpustakan, params)
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(perpustakan)
}

func main() {
	mux := http.NewServeMux()
	perpustakan = []Buku{
		{ID: 1, Judul: "Halo Book", Pengarang: "Siti"},
	}
	mux.HandleFunc("/books", GetBooks)
	mux.HandleFunc("/create-books", CreateBooks)
	http.ListenAndServe(":8000", mux)
}
