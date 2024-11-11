package groupie

import (
	"encoding/json"
	"io"
	"net/http"
)

var (
	LocationURL = "https://groupietrackers.herokuapp.com/api/locations"
	ArtistURL = "https://groupietrackers.herokuapp.com/api/artists"
	DatesURL =  "https://groupietrackers.herokuapp.com/api/dates"
	RelationURL = "https://groupietrackers.herokuapp.com/api/relation"
)

func FetchArtists() ([]Artist, error) {
	resp, err := http.Get(ArtistURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var artists []Artist
	err = json.Unmarshal(body, &artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func FetchLocations() ([]Location, error) {
	resp, err := http.Get(LocationURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data struct {
		Index []Location `json:"index"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data.Index, nil
}

func FetchDates() ([]Date, error) {
	resp, err := http.Get(DatesURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data struct {
		Index []Date `json:"index"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	} 
	return data.Index, nil
}

func FetchRelations() ([]Relation, error) {
	resp, err := http.Get(RelationURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data struct {
		Index []Relation `json:"index"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data.Index, nil
}