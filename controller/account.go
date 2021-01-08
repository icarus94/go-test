package controller

import (
	"fmt"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login hit")
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "register hit")
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "logut hit")
}
