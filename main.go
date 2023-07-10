package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type api struct {
	ID           int    `json:"id"`
	IMAGES       string `json:images`
	NAMES        string `json:names`
	MEMBERS      string `json:members`
	CREATIONDATE string `json:creationDate`
	FISTERALBUM  string `json:fisterAlbum`
	LOCATIONS    string `json:locations`
	CONCERTDATES string `json:concertDates`
	RELATIONS    string `json:relations`
}

func GetApi(url string) {
	responseBody, err := http.Get(url)
	if err != nil {
		fmt.Println("sdfrudiu")
	}
	defer responseBody.Body.Close()

	res, err := ioutil.ReadAll(responseBody.Body)
	if err != nil {

		fmt.Println("sdfrudiu")
	}

	for _, v := range res {
		fmt.Print(string(v))
	}

}

func main() {
	a := "https://groupietrackers.herokuapp.com/api/artists"
	GetApi(a)
}
