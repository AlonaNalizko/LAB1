package main

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"
)

type Phone struct {
    TypeID      int `json:"type_id"`
    CountryCode int `json:"country_code"`
    Operator    int `json:"operator"`
    Number      int `json:"number"`
}

type Contact struct {
    ID         int      `json:"id"`
    Username   string   `json:"username"`
    GivenName  string   `json:"given_name"`
    FamilyName string   `json:"family_name"`
    FullName   string   `json:"full_name"`
    Phone      []Phone  `json:"phone"`
    Email      []string `json:"email"`
    Birthdate  string   `json:"birthdate"`
}

type Group struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Contacts    []int  `json:"contacts"`
}

func handlerContact(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Contact{})
}

func handlerGroup(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(Group{})
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/api/v1/contact", handlerContact).Methods("POST", "GET", "PUT", "DELETE")
    r.HandleFunc("/api/v1/group", handlerGroup).Methods("POST", "GET", "PUT", "DELETE")

    http.ListenAndServe(":6080", r)
}