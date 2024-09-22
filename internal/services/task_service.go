package services

import (
	"io"
	"net/http"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gotten Task!"))
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	w.Write([]byte(body))
}