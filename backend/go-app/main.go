package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"path"
)

func main() {
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
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getPublisher(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
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
