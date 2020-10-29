package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Post struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

var posts []Post

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		log.Fatal(err)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000000))
	posts = append(posts, post)
	err := json.NewEncoder(w).Encode(&post)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range posts {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(&Post{})
	if err != nil {
		log.Fatal(err)
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			var post Post
			_ = json.NewDecoder(r.Body).Decode(&post)
			post.ID = params["id"]
			posts = append(posts, post)
			err := json.NewEncoder(w).Encode(&post)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		log.Fatal(err)
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		log.Fatal(err)
	}
}