package groupie

import (
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		HandleError(w,r , 405)
		return
	}

	if r.URL.Path != "/" {
		HandleError(w,r , 404)
		return
	}

	artists, err := FetchArtists()
	if err != nil {
		HandleError(w,r , 500)
		return
	}

	Locations, err := FetchLocations()
	if err != nil {
		HandleError(w,r , 500)
		return
	}

	// Parse and execute the template for the homepage (artist images)
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		HandleError(w,r , 500)
		return
	}

	seen := make(map[string]bool)
	var locati []string
	for _, loca := range Locations {
		for _, r := range loca.Locations {
			if !seen[r] {
				locati = append(locati, r)
				seen[r] = true
			}
		}
	}

	data := struct {
		Artist   []Artist
		Location []string
	}{
		Artist:   artists,
		Location: locati,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		HandleError(w,r , 500)
		return
	}
}
