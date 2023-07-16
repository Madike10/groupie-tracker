package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"text/template"
)

type ArtistAPI struct {
	ID           int          `json:"id"`
	IMAGE        string       `json:"image"`
	NAME         string       `json:"name"`
	MEMBERS      []string     `json:"members"`
	CREATIONDATE int          `json:"creationDate"`
	FIRSTALBUM   string       `json:"firstAlbum"`
	LOCATIONS    LocationsAPI `json:"locations"`
	CONCERTDATES string       `json:"concertDates"`
	RELATIONS    string       `json:"relations"`
}

type LocationsAPI struct {
	Id            int      `json:"id"`
	Locations     []string `json:"locations"`
	DateLocations int      `json:"dates"`
}
type NewApiGroupie struct {
	ID           int
	IMAGE        string
	NAME         string
	MEMBERS      []string
	CREATIONDATE int
	FIRSTALBUM   string
	LOCATIONS    LocationsAPI
	CONCERTDATES string
	RELATIONS    string
}

func GetApi(url string) ([]byte, error) {
	responseBody, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer responseBody.Body.Close()

	data, err := ioutil.ReadAll(responseBody.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
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

			// Récupérer les données depuis le lien "locations"
			for _, locURL := range v.LOCATIONS.Locations {
				locationsData, err := GetApi(locURL)
				if err == nil {
					// Ajouter les données de localisation à newGroupie
					newGroupie.LOCATIONS.Locations = append(newGroupie.LOCATIONS.Locations, string(locationsData))
					fmt.Println(newGroupie.LOCATIONS.Locations)
				} else {
					fmt.Println("Erreur lors de la récupération des données de localisation:", err)
				}
			}
		}
	}
	template.Execute(resp, newGroupie)
}
