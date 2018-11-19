package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const dataSourceName = "capstonefinal:capstonefinal@tcp(capstonefinal.culhn7yxy1fu.us-east-1.rds.amazonaws.com:3306)/capstone"
//
const driverName = "mysql"

type Response struct {
	Success string `json:"success"`
}

type User struct {
	ID              	string `json:"id"`
	Email           	string `json:"email"`
	FirstName      		string `json:"firstName"`
	LastName 			string `json:"lastName"`
	IsBeingTutored      string `json:"isBeingTutored"`
	Team				string `json:"team"`
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	results, err := db.Query("SELECT * FROM user WHERE email = " + "'" + params["email"] + "'")
	if err != nil {
		panic(err.Error())
	}

	var user User
	for results.Next() {
		err = results.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.IsBeingTutored, &user.Team)
		if err != nil {
			panic(err.Error())
		}
	}

	jsonArray := [1]User{user}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonArray)
	results.Close()
	db.Close()
	println("Good to Go!")
}


func insertUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	results, err := db.Query("INSERT INTO `user` (email, firstName, lastName, isBeingTutored, team)" +
		" VALUES (" + " '" + user.Email + "', " + " '" + user.FirstName + "', " + " '" + user.LastName + "', " +
		" '" + user.IsBeingTutored + "', " + " '" + user.Team + "' )")
	if err != nil {
		panic(err.Error())
	}

	results.Close()
	db.Close()
	println("Good to Go!")
}


func updateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	results, err := db.Query("UPDATE `user` SET `email` = '" + user.Email + "', `firstName` = '" + user.FirstName + "', `lastName` = '" + user.LastName +
		"', `isBeingTutored` = '" + user.IsBeingTutored + "' WHERE `email` = '" + user.Email + "'")
	if err != nil {
		panic(err.Error())
	}

	var response Response
	response.Success = "true"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	results.Close()
	db.Close()
	println("Good to Go!")
}


func deleteUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	results, err := db.Query("DELETE  FROM  `user` WHERE `email` = '" + user.Email + "'")
	if err != nil {
		panic(err.Error())
	}

	var response Response
	response.Success = "true"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	results.Close()
	db.Close()
	println("Good to Go!")
}

func main() {

	// Init router, this things handles the routes...
	r := mux.NewRouter()

	// Route handles & endpoints, way easier then I thought. If it has been tested with POSTMAN I have a GOOD next to it.
	r.HandleFunc("/api/user/{email}", getUser).Methods("GET")       // GOOD
	r.HandleFunc("/api/user/{data}", insertUser).Methods("POST")    // GOOD
	r.HandleFunc("/api/user/{email}", updateUser).Methods("PUT")    // GOOD
	r.HandleFunc("/api/user/{email}", deleteUser).Methods("DELETE") // GOOD

	// Don't forget, this creates Goroutines with channels for u
	log.Fatal(http.ListenAndServe(":8080", r))
}