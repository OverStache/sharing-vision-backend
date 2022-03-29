package handlers

import (
	"encoding/json"
	"fmt"
	"goCRUD/connection"
	"goCRUD/structs"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wilkommen!")
}

func Create(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)
	var request structs.Users
	json.Unmarshal(payloads, &request)

	hash, _ := HashPassword(request.Password)
	var user = structs.Users{Name: request.Name, Age: request.Age, Password: hash}
	connection.DB.Create(&user)

	var stock, mm, bond float32
	x := 55 - user.Age
	if x >= 30 {
		stock = 72.5
		bond = 21.5
		mm = 100 - (stock + bond)
	} else if x >= 20 {
		stock = 54.5
		bond = 25.5
		mm = 100 - (stock + bond)
	} else if x < 20 {
		stock = 34.5
		bond = 45.5
		mm = 100 - (stock + bond)
	}

	var risk_profile = structs.Risk_profiles{Id_user: user.ID, Stock: stock, Bond: bond, MM: mm}
	json.Unmarshal(payloads, &risk_profile)
	connection.DB.Create(&risk_profile)

	res := structs.Result{Code: 200, Data: user, Message: "Success create user"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("take")
	offset := r.URL.Query().Get("page")

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

func Login(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)
	var request structs.Login
	json.Unmarshal(payloads, &request)

	var user structs.Users
	connection.DB.First(&user, "name = ?", request.Name)

	match := CheckPasswordHash(request.Password, user.Password)

	var res structs.Result
	if match {
		res = structs.Result{Code: 200, Data: match, Message: "Login Success"}
	} else {
		res = structs.Result{Code: 400, Data: match, Message: "Login Failed"}
	}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
