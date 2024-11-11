package groupie

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"text/template"
)

type l9nota struct {
	Id        string
	Place     string
	Latitude  float64
	Longitude float64
}

func Geolocalisation(w http.ResponseWriter, r *http.Request) {
	lmap := r.FormValue("localisation")
	latitude, longitude, err := Geocode(lmap)
	if err != nil {
		fmt.Println(err)
		return
	}
	id := r.FormValue("id")
	data := l9nota{
		Id:        id,
		Place:     lmap,
		Latitude:  latitude,
		Longitude: longitude,
	}
	tmpl, err := template.ParseFiles("templates/map.html")
	if err != nil {
		HandleError(w, r, 500)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		HandleError(w, r, 500)
		return
	}
}

type Res struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func Geocode(address string) (float64, float64, error) {
	address = strings.ReplaceAll(address, "-", ",")
	apiURL := fmt.Sprintf("https://api.geoapify.com/v1/geocode/search?text=%v&format=json&apiKey=5ba3eaa6c92d48c18cccc0c77b034bca", address)
	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}
	var data struct {
		RES []Res `json:"results"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err)
	}
	if len(data.RES) == 0 {
		return 0, 0, err
	}
	return data.RES[0].Lat, data.RES[0].Lon, nil
}
