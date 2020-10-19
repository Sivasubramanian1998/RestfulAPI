package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Contact struct
type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

// Init contacts var as a slice Contact struct
var contacts []Contact

// Get all contacts
func getContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

// Get single contact
func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Looping through contacts and find one with the id from the params
	for _, item := range contacts {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Contact{})
}

// Add new contact
func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contact)
}

// Delete contact
func deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range contacts {
		if item.Name == params["name"] {
			contacts = append(contacts[:idx], contacts[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(contacts)
}

// Update contact
func updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range contacts {
		if item.Name == params["name"] {
			contacts = append(contacts[:idx], contacts[idx+1:]...)
			var contact Contact
			_ = json.NewDecoder(r.Body).Decode(&contact)
			contact.Name = params["name"]
			contacts = append(contacts, contact)
			json.NewEncoder(w).Encode(contact)
			return
		}
	}
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data
	contacts = append(contacts, Contact{Name: "akash", Phone: "98231-12435", Email: "akash@mail.com"})
	contacts = append(contacts, Contact{Name: "parthiban", Phone: "96321-54321", Email: "parthiban@mail.com"})
	contacts = append(contacts, Contact{Name: "akhil", Phone: "97123-12345", Email: "akhil@mail.com"})

	// Route handles & endpoints
	r.HandleFunc("/contacts", getContacts).Methods("GET")
	r.HandleFunc("/contacts/{name}", getContact).Methods("GET")
	r.HandleFunc("/contacts", createContact).Methods("POST")
	r.HandleFunc("/contacts/{name}", updateContact).Methods("PUT")
	r.HandleFunc("/contacts/{name}", deleteContact).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":3000", r))
}
