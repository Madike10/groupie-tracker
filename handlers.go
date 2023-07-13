package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type ApiGroupie struct {
	ID           int      `json:"id"`
	IMAGE        string   `json:"image"`
	NAME         string   `json:"name"`
	MEMBERS      []string `json:"members"`
	CREATIONDATE int      `json:"creationDate"`
	FIRSTALBUM   string   `json:"firstAlbum"`
	LOCATIONS    string   `json:"locations"`
	CONCERTDATES string   `json:"concertDates"`
	RELATIONS    string   `json:"relations"`
}

type NewApiGroupie struct {
	ID           int
	IMAGE        string
	NAME         string
	MEMBERS      []string
	CREATIONDATE int
	FIRSTALBUM   string
	LOCATIONS    map[string]string
	CONCERTDATES string
	RELATIONS    string
}

var Groupie []ApiGroupie

var newGroupie NewApiGroupie

// Parcours de chaque élément de Groupie

func GetApi(url string) {
	responseBody, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer responseBody.Body.Close()
	// A ce niveau nous prenons les donées en json que l'on stocke au niveau de ApiGroupies
	//qui est une instance de ApiGroupies
	err = json.NewDecoder(responseBody.Body).Decode(&Groupie)

	if err != nil {
		fmt.Println(err)
	}
}
func HomePage(resp http.ResponseWriter, req *http.Request) {
	template, _ := template.ParseFiles("./templates/index.html")
	template.Execute(resp, Groupie)

}
func ArtistPage(resp http.ResponseWriter, req *http.Request) {
	template, _ := template.ParseFiles("./templates/artist.html")
	val := req.URL.Query().Get("id")
	fmt.Println(val)
	for _, v := range Groupie {
		a, _ := strconv.Atoi(val)
		if v.ID == a {

			newGroupie.NAME = v.NAME
			newGroupie.MEMBERS = v.MEMBERS
			newGroupie.IMAGE = v.IMAGE
			newGroupie.CONCERTDATES = v.CONCERTDATES
			newGroupie.FIRSTALBUM = v.FIRSTALBUM
			newGroupie.CREATIONDATE = v.CREATIONDATE
			newGroupie.LOCATIONS = v.LOCATIONS
			newGroupie.RELATIONS = v.RELATIONS

		}
	}
	template.Execute(resp, newGroupie)
}
