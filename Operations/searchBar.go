package groupie

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

type Art struct{
	Artist []Artist
	Nart []Artist
	Location []Location
}

func SearchBar(w http.ResponseWriter, r *http.Request) {
	search := strings.ToLower(r.FormValue("search"))
	if len(search) == 0 {
		HandleError(w,r , 400)
		return
	}

	artist, err := FetchArtists()
	if err != nil {
		HandleError(w,r , 500)
		return
	}

	locations, err := FetchLocations()
	if err != nil {
		HandleError(w,r , 500)
		return
	}

	var NewArt []Artist

	found := map[int]bool{}
	for _, artist := range artist {
		if strings.Contains(strings.ToLower(artist.Name), search) {
			NewArt = append(NewArt, artist)
			found[artist.ID] = true
		}
	}

	for _, artist := range artist {
		for _, Members := range artist.Members {
			if strings.Contains(strings.ToLower(Members), search) {
				if !found[artist.ID] {
					NewArt = append(NewArt, artist)
					found[artist.ID] = true
				}
			}
		}
	}

	for _, artist := range artist {
		if strings.Contains(fmt.Sprintf("%v", artist.CreationDate), search) {
			if !found[artist.ID] {
				NewArt = append(NewArt, artist)
				found[artist.ID] = true
			}
		}
	}

	for _, artist := range artist {
		if strings.Contains(strings.ToLower(artist.FirstAlbum), search) {
			if !found[artist.ID] {
				NewArt = append(NewArt, artist)
				found[artist.ID] = true
			}
		}
	}

	for _, artist := range artist {
		for _, location := range locations {
			for _, loc := range location.Locations {
				if strings.Contains(strings.ToLower(loc), search) {
					if artist.ID == location.ID && !found[artist.ID] {
						NewArt = append(NewArt, artist)
						found[artist.ID] = true
					}
				}
			}
		}
	}

	if len(NewArt) == 0 {
		HandleError(w,r , 404)
		return
	}

	tmpl, err := template.ParseFiles("templates/search-bar.html")
	if err != nil {
		HandleError(w,r , 500)
		return
	}

	data := Art{
		Artist: artist,
		Nart: NewArt,
		Location: locations,
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		HandleError(w,r , 500)
		return
	}
}
