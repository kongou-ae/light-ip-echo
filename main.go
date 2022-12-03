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
	Host string `json:"Host"`
	XAzureFDID string `json:"XAzureFDID"`
}

func returnInfo(w http.ResponseWriter, r *http.Request) {

	result := result{
		RemoteAddr:    r.RemoteAddr,
		RequestURI:    r.RequestURI,
		XForwardedFor: r.Header.Get("X-FORWARDED-FOR"),
		Host: r.Header.Get("Host"),
		XAzureFDID: r.Header.Get("X-Azure-FDID"),
	}

	log.Println("RemoteAddr:" + r.RemoteAddr + " | RequestURI:" + r.RequestURI + " | X-FORWARDED-FOR:" + r.Header.Get("X-FORWARDED-FOR") + " | Host:" + r.Header.Get("Host") + " | XAzureFDID:" + r.Header.Get("XAzureFDID"))
	w.Header().Set("Content-Type", "application/json")

	res, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	w.Write(res)
}

func main() {
	http.HandleFunc("/", returnInfo)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
