package main

import (
	"encoding/json"
	"net/http"
	"os"
	"path"
)

func main() {
	initDb()
	server := http.Server{
		Addr: ":" + os.Getenv("PORT"),
	}
	http.HandleFunc("/publisher/", handlePublisher)
	server.ListenAndServe()
}

func handlePublisher(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = getPublisher(w, r)
	case "POST":
		err = postPublisher(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getPublisher(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	publisher, err := retrievePublisher(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(publisher, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// Create a publisher
// POST /publisher/
func postPublisher(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var p Publisher
	json.Unmarshal(body, &p)
	err = p.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
