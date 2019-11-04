// encoding/json

// Marshalling - Convertion data structure in go into JSON by using something called marshalling which create a byte slice containing
// a very long string with no extraneous white space

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func main() {

	handleRequest()
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	// handleRequests()
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles) // Does the job of encoding our articles array into json string and then writing as part of our response

}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	var := mux.Vars(r)
	key :=  var["id"]

	for _, article := range Aricles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func CreatingArticle(w http.ResponseWriter, r *http.Request) {
	reqBody := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	Article = append(Articles, article)
	json.NewEncoder(w).Encode(article)

}


func handleRequest() {
	http.HandleFunc("/", rootFunc)
	http.HandleFunc("/read", returnAllArticles)

	log.Fatal(http.ListenAndServe(":9090", nil))
}

func rootFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, " Welcome to homepage")
}
