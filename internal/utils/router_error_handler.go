package utils

import (
	"log"
	"net/http"
)

func HandleInternalError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	errorText := err.Error()
	l, err := w.Write([]byte(errorText))
	if err != nil {
		log.Println("error writing response:", err)
	}
	if l != len(errorText) {
		log.Println("Content length and written length does not match in response")
	}
}
