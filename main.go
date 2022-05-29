package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type result struct {
	RemoteAddr    string `json:"RemoteAddr"`
	RequestURI    string `json:"RequestURI"`
	XForwardedFor string `json:"Xforwardedfor"`
}

func returnInfo(w http.ResponseWriter, r *http.Request) {

	result := result{
		RemoteAddr:    r.RemoteAddr,
		RequestURI:    r.RequestURI,
		XForwardedFor: r.Header.Get("X-FORWARDED-FOR"),
	}

	w.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	w.Write(res)
}

func main() {
	http.HandleFunc("/", returnInfo)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
