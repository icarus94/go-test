package controller

import (
	"fmt"
	"net/http"
)

func addContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login hit")
}

func editContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "edit hit")
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete contact")
}

func listContacts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "list contacts")
}
