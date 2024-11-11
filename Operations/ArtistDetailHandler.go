package groupie

import (
	"fmt"
	"net/http"
	"text/template"
)

func ArtistDetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		HandleError(w,r , 405)
		return
	}

	artistID := r.PathValue("id")

	artists, err := FetchArtists()
	if err != nil {
		HandleError(w,r , 500)
		return
	}
	locations, err := FetchLocations()
	if err != nil {
		HandleError(w,r , 500)
		return
	}
	dates, err := FetchDates()
	if err != nil {
		HandleError(w,r , 500)
		return
	}
	relations, err := FetchRelations()
	if err != nil {
		HandleError(w,r , 500)
		return
	}

	// Find the selected artist by ID
	found := false
	var selectedArtist Artist
	for _, artist := range artists {
		if artistID == fmt.Sprintf("%d", artist.ID) {
			selectedArtist = artist
			found = true
			break
		}
	}

	// If no artist is found
	if !found {
		HandleError(w,r , 404)
		return
	}

	// Find the corresponding locations for this artist
	var artistLocations []string
	for _, loc := range locations {
		if loc.ID == selectedArtist.ID {
			artistLocations = loc.Locations
			break
		}
	}

	// Find the corresponding dates for this artist
	var artistDates []string
	for _, date := range dates {
		if date.ID == selectedArtist.ID {
			artistDates = date.Dates
			break
		}
	}

	// Find the corresponding relations (dates and locations) for this artist
	var artistRelations map[string][]string
	for _, relation := range relations {
		if relation.ID == selectedArtist.ID {
			artistRelations = relation.Relations
			break
		}
	}
	
	// latitude, longitude, err := Geocode("Germany Mainz")

	// Prepare the data to pass to the template
	data := ArtistDetailPageData{
		Artist:    selectedArtist,
		Locations: artistLocations,
		Dates:     artistDates,
		Relations: artistRelations,
	}

	// Parse and render the artist details template
	tmpl, err := template.ParseFiles("templates/artist_detail.html")
	if err != nil {
		HandleError(w,r , 500)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		HandleError(w,r , 500)
		return
	}
}
