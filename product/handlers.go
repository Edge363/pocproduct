package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"encoding/json"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	productid := mux.Vars(r)["productid"]
	svc := Localsvc

	product, err := getProduct(&productid, svc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	if product.Id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	productjson, err := json.Marshal(product)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(productjson)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	svc := Localsvc

	decoder := json.NewDecoder(r.Body)
	product := new(Product)
	err := decoder.Decode(&product)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	err = putProduct(product, svc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	productid := mux.Vars(r)["productid"]
	svc := Localsvc
	decoder := json.NewDecoder(r.Body)
	product := new(Product)
	err := decoder.Decode(&product)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	err = postProduct(&productid, product, svc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	productid := mux.Vars(r)["productid"]
	svc := Localsvc
	err := deleteProduct(&productid, svc)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}