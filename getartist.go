package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var Groupie []ArtistAPI

var newGroupie NewApiGroupie

// Parcours de chaque élément de Groupie
func GetApiArtist(url string) {
	responseBody, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer responseBody.Body.Close()
	err = json.NewDecoder(responseBody.Body).Decode(&Groupie)

	if err != nil {
		fmt.Println(err)
	}
}
func HomePage(resp http.ResponseWriter, req *http.Request) {
	template, _ := template.ParseFiles("./templates/index.html")
	template.Execute(resp, Groupie)

}
