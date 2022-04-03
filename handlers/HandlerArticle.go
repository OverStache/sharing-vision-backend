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

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)
	var request structs.Posts
	json.Unmarshal(payloads, &request)

	connection.DB.Create(&request)

	res := structs.Result{Code: 200, Data: request, Message: "Success create post"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var post structs.Posts
	connection.DB.First(&post, "id = ?", id)

	res := structs.Result{Code: 200, Data: post, Message: "Success get post"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit := vars["limit"]
	offset := vars["offset"]

	posts := []structs.Posts{}

	connection.DB.
		Limit(limit).
		Offset(offset).
		Find(&posts)

	res := structs.Result{Code: 200, Data: posts, Message: "Success get posts"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)
	var request structs.Request
	json.Unmarshal(payloads, &request)

	var post structs.Posts
	connection.DB.First(&post, "id = ?", id)
	post.Title = request.Title
	post.Content = request.Content
	post.Category = request.Category
	post.Status = request.Status
	connection.DB.Save(&post)

	res := structs.Result{Code: 200, Data: post, Message: "Success update post"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var post structs.Posts
	connection.DB.First(&post, "id = ?", id).
		Model(&post).Update("status", "trash")

	res := structs.Result{Code: 200, Data: post, Message: "Success trash post"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}
