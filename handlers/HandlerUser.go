package handlers

import (
	"encoding/json"
	"fmt"
	"goCRUD/connection"
	"goCRUD/structs"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wilkommen!")
}

func Create(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var user structs.Users
	json.Unmarshal(payloads, &user)
	connection.DB.Create(&user)

	var stock, mm, bond float32
	if user.Age >= 30 {
		stock = 72.5
		bond = 21.5
		mm = 100 - (stock + bond)
	} else if user.Age >= 20 {
		stock = 54.5
		bond = 25.5
		mm = 100 - (stock + bond)
	} else if user.Age < 20 {
		stock = 34.5
		bond = 45.5
		mm = 100 - (stock + bond)
	}

	var risk_profile = structs.Risk_profiles{Id_user: user.ID, Stock: stock, Bond: bond, MM: mm}
	json.Unmarshal(payloads, &risk_profile)
	connection.DB.Create(&risk_profile)

	res := structs.Result{Code: 200, Data: risk_profile, Message: "Success create user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// limit := vars["limit"]
	// offset := vars["offset"]
	limit := r.URL.Query().Get("page")
	offset := r.URL.Query().Get("take")

	users := []structs.Users{}

	connection.DB.
		Limit(limit).
		Offset(offset).
		Find(&users)

	res := structs.Result{Code: 200, Data: users, Message: "Success get users"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	var risk structs.Risk_profiles
	connection.DB.First(&risk, "id_user = ?", userId)

	res := structs.Result{Code: 200, Data: risk, Message: "Success get user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
