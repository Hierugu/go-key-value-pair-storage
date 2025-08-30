package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var store = make(map[string]string)

var ErrorNoSuchKey = errors.New("No such key")

func Put(key, value string) error {
	store[key] = value
	return nil
}

func Get(key string) (string, error) {
	value, ok := store[key]
	if !ok {
		return "", ErrorNoSuchKey
	}
	return value, nil
}

func Delete(key string) error {
	delete(store, key)
	return nil
}

// Удоавлетворяет типу HandlerFunc
func helloMuxHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello gorilla/mux!\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloMuxHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
