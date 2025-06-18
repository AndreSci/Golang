package main

import (
	"errors"
	"fmt"
	"hw3/client"
	"log"
	"net/http"
	"time"
)

func requestItem(typeReq string) ([]byte, error) {
	coincapClient, err := client.NewClient(time.Second * 10)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var assets []byte

	// assets, errorAs := coincapClient.GetAssets()
	if typeReq == "GetData" {
		assets, err = coincapClient.GetData()
	} else if typeReq == "GetNames" {
		assets, err = coincapClient.GetNames()

	} else {
		err = errors.New("wrong request type")
		log.Fatal(err)
		return nil, err
	}

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return assets, nil
}

func main() {
	fmt.Println("START SERVER")

	http.HandleFunc("/GetRow", handleDataGET)
	http.HandleFunc("/GetNames", handleNamesGET)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleDataGET(w http.ResponseWriter, r *http.Request) {

	fmt.Println("REQUEST INCOME")
	switch r.Method {
	case http.MethodGet:
		getData(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func handleNamesGET(w http.ResponseWriter, r *http.Request) {

	fmt.Println("REQUEST INCOME")
	switch r.Method {
	case http.MethodGet:
		getNames(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func getData(w http.ResponseWriter, r *http.Request) {

	resp, err := requestItem("GetData")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func getNames(w http.ResponseWriter, r *http.Request) {

	resp, err := requestItem("GetNames")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}
