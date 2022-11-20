package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type customer struct {
	ID        string
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}

// Note: short hand variable declarion doesn't work outside the function declaration
var customer1 = customer{ID: "1", Name: "Aashay", Role: "A", Email: "aashay_dhawal@abc.com", Phone: "1234567890", Contacted: false}
var customer2 = customer{ID: "2", Name: "Shilpi", Role: "A", Email: "shilpi_suman@abc.com", Phone: "0123456789", Contacted: false}
var customer3 = customer{ID: "3", Name: "Bhavya", Role: "A", Email: "bhavya_dhawal@abc.com", Phone: "9012345678", Contacted: false}

// default customers stored in CRM database
var custom_db = map[string]customer{
	"1": customer1,
	"2": customer2,
	"3": customer3,
}

func showProjectInfo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "project_info.html")
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(custom_db)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	// read the request from the user, for which we want to fetch the data
	id := mux.Vars(r)["id"]

	if _, ok := custom_db[id]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(custom_db[id])
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if _, ok := custom_db[id]; ok {
		delete(custom_db, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(custom_db)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(custom_db)
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newEntry customer
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newEntry)

	if _, ok := custom_db[newEntry.ID]; ok {
		w.WriteHeader(http.StatusConflict)
	} else {
		custom_db[newEntry.ID] = newEntry
		w.WriteHeader(http.StatusCreated)
	}
	json.NewEncoder(w).Encode(custom_db)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updateEntry customer
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updateEntry)

	id := mux.Vars(r)["id"]
	if _, ok := custom_db[id]; ok {
		custom_db[id] = updateEntry
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusConflict)
	}
	json.NewEncoder(w).Encode(custom_db)

}

func main() {

	fmt.Println("starting udacity's CRM project")
	fmt.Println()

	router := mux.NewRouter()
	http.HandleFunc("/", showProjectInfo)
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("POST")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	fmt.Println("Server is starting on port 3002...")
	log.Fatal(http.ListenAndServe(":3002", router))

}
