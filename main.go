package main

import (
	"fmt"
	"net/http"
)

func main() {
	a := "https://groupietrackers.herokuapp.com/api/artists"
	GetApi(a)
    locations := "https://groupietrackers.herokuapp.com/api/locations"
	GetApi(locations)
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/artist/", ArtistPage)

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// res := json.NewDecoder(responseBody.Body)
// result := res.Decode(&apiGroupie)
// if result != nil {
// 	fmt.Println("error")
// }
// fmt.Print(apiGroupie)
