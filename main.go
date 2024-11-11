package main

import (
	"fmt"
	"log"
	"net/http"

	x "groupie/Operations"
)

func main() {
	http.HandleFunc("/", x.HomeHandler)
	http.HandleFunc("/Artist/{id}", x.ArtistDetailHandler)
	http.HandleFunc("/searchBar", x.SearchBar)
	http.HandleFunc("/filter", x.FilterHandler)
	http.HandleFunc("/static/", x.StaticHandle)
	http.HandleFunc("/Geoloca", x.Geolocalisation)

	fmt.Println("http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
