package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article - Our struct for all articles
type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func main() {
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func homePage(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Welcome to the Articles Service"))
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(Articles)
	fmt.Println("Endpoint Hit: returnAllArticles")
}

func createNewArticle(w http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	var article Article
	json.Unmarshal(body, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(Articles)
	fmt.Println("Endpoint Hit: createNewArticle")
}

func returnSingleArticle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pathParams := mux.Vars(req)
	articleID := pathParams["id"]
	for _, article := range Articles {
		if article.Id == articleID {
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	w.Write([]byte(fmt.Sprintf(`{"error_msg":"Article with Id %s Not found"}`, articleID)))
	fmt.Println("Endpoint Hit: createNewArticle")
}

func deleteArticle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pathParams := mux.Vars(req)
	articleID := pathParams["id"]
	for index, article := range Articles {
		if article.Id == articleID {
			Articles = append(Articles[:index], Articles[index+1:]...)
			json.NewEncoder(w).Encode(Articles)
			return
		}
	}
	w.Write([]byte(fmt.Sprintf(`{"error_msg":"Article with Id %s Not found"}`, articleID)))
	fmt.Println("Endpoint Hit: createNewArticle")
}
