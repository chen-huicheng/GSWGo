package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	InitLog()
	InitWorker()
	http.HandleFunc("/post/json", JsonHandler)
	fmt.Printf("please post user to http://10.37.141.193:8000/post/json\n")
	http.ListenAndServe("0.0.0.0:8000", nil)
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read the body into a string for json decoding
	var jsonReq = &JsonReq{}
	err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&jsonReq)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("post json response ok,%v", jsonReq)
	for _, user := range jsonReq.Users {
		work := Job{User: user}
		JobQueue <- work
	}

	w.WriteHeader(http.StatusOK)
}
