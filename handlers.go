package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

var apiGroupie []ApiGroupie

func GetApi(url string) {
	responseBody, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer responseBody.Body.Close()
	// A ce niveau nous prenons les donées en json que l'on stocke au niveau de ApiGroupies
	//qui est une instance de ApiGroupies
	err = json.NewDecoder(responseBody.Body).Decode(&apiGroupie)

	if err != nil {
		fmt.Println(err)
	}

	// traite des donner pour l'accueil

	// for _, data := range apiGroupie {
	// 	fmt.Println("ID", data.ID)
	// 	fmt.Println("IMAGES", data.IMAGE)
	// 	fmt.Println("NAMES", data.NAME)
	// fmt.Println("MEMBERS", data.MEMBERS)
	// 	fmt.Println("CRATIONDATE", data.CREATIONDATE)
	// 	fmt.Println("FIRSTALBUM", data.FIRSTALBUM)
	// 	fmt.Println("LOCATIONS", data.LOCATIONS)
	// 	fmt.Println("CONCERTDATE", data.CONCERTDATES)
	// 	fmt.Println("RELATIONS", data.RELATIONS)

	// }
}
func HomePage(resp http.ResponseWriter, req *http.Request) {
	template, _ := template.ParseFiles("./templates/index.html")
	template.Execute(resp, apiGroupie)

}
func ArtistPage(resp http.ResponseWriter, req *http.Request) {
	template, _ := template.ParseFiles("./templates/artist.html")

	// for _, v := range apiGroupie{
	// 	a , _ := strconv.Atoi(id)
	// 	if v.ID == a {
	// 		v.NAME = N
	// 	}
	// }

	template.Execute(resp, apiGroupie)
}

// res, err := ioutil.ReadAll(responseBody.Body)
// if err != nil {
// 	log.Fatal(err)
// }
// for _, v := range res {
// 	fmt.Print(string(v))
// }
