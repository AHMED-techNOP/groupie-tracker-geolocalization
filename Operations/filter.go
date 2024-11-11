package groupie

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func ToInt(s string) (int, error) {
	if s == "" {
		return 0, nil
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return n, nil
}

var (
	FoundArtists []Artist
	fond         bool
)

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/filter" {
		HandleError(w,r , 404)
		return
	}

	if r.Method != "GET" {
		HandleError(w,r , 405)
		return
	}

	// Validate and extract parameters
	minalbum, maxalbum, mincreation, maxcreation, members, alocation, bigger, err := validateAndExtractParams(r)
	if err != nil || bigger == 1 {
		HandleError(w,r , 400)
		return
	}

	if mincreation == 0 && maxcreation == 0 || minalbum == 0 && maxalbum == 0 {
		HandleError(w,r , 400)
		return
	}

	// Perform filtering
	Filter(w, r, members, minalbum, maxalbum, mincreation, maxcreation, alocation)
	if !fond {
		HandleError(w,r , 404)
		return
	}
	// Load and execute template
	temp, err := template.ParseFiles("templates/filter.html")
	if err != nil {
		HandleError(w,r , 500)
		return
	}
	if err := temp.Execute(w, FoundArtists); err != nil {
		HandleError(w,r , 500)
		return
	}
}

func Filter(w http.ResponseWriter, r *http.Request, members []string, minalbum int, maxalbum int, creatmin int, creatmax int, alocation string) {
	// var location Location2
	found := []Artist{}
	artists, err := FetchArtists()
	if err != nil {
		HandleError(w,r , 500)
		return
	}
	location, err := FetchLocations()
	if err != nil {
		HandleError(w,r , 500)
		return
	}
	fond = false

	for _, artist := range artists {
		if CreationDate(creatmin, creatmax, artist.CreationDate) &&
			Album(minalbum, maxalbum, artist.FirstAlbum) &&
			Members(members, len(artist.Members)) && sLocation(location, alocation, artist.ID) {
			found = append(found, artist)
		}
	}
	FoundArtists = found
}

func CreationDate(creatmin int, creatmax int, creatDate int) bool {
	return creatDate >= creatmin && creatDate <= creatmax
}

func Album(minalbum int, maxalbum int, firstdate string) bool {
	year, _ := strconv.Atoi(firstdate[6:])
	return year >= minalbum && year <= maxalbum
}

func Members(nmembers []string, num int) bool {
	if nmembers == nil {
		return true
	}

	for _, v := range nmembers {
		// _,err :=ToInt(v)
		if strconv.Itoa(num) == v {
			return true
		}
	}
	return false
}

func sLocation(slocation []Location, alocation string, id int) bool {
	if alocation == "" {
		fond = true
		return true
	}

	for _, loca := range slocation {
		for _, l := range loca.Locations {
			ll := strings.ReplaceAll(alocation, ", ", "-")
			if strings.Contains(strings.ToLower(l), ll) && loca.ID == id {
				fond = true
				return true
			}
		}
	}

	return false
}

func validateMembers(members []string) bool {
	for _, r := range members {
		num, err := strconv.Atoi(r)
		if err != nil || num < 1 || num > 8 {
			return false
		}
	}
	return true
}

func validateAndExtractParams(r *http.Request) (int, int, int, int, []string, string, int, error) {
	minalbum, err := ToInt(r.FormValue("firstAlbumMin"))
	if err != nil {
		return 0, 0, 0, 0, nil, "", 0, err
	}
	maxalbum, err := ToInt(r.FormValue("firstAlbumMax"))
	if err != nil {
		return 0, 0, 0, 0, nil, "", 0, err
	}
	if minalbum > maxalbum {
		return 0, 0, 0, 0, nil, "", 1, err
	}

	mincreation, err := ToInt(r.FormValue("creationDateMin"))
	if err != nil {
		return 0, 0, 0, 0, nil, "", 0, err
	}
	maxcreation, err := ToInt(r.FormValue("creationDateMax"))
	if err != nil {
		return 0, 0, 0, 0, nil, "", 0, err
	}
	if mincreation > maxcreation {
		return 0, 0, 0, 0, nil, "", 1, err
	}

	members := r.Form["member"]
	if !validateMembers(members) {
		return 0, 0, 0, 0, nil, "", 1, err
	}

	alocation := strings.ToLower(r.FormValue("location"))

	return minalbum, maxalbum, mincreation, maxcreation, members, alocation, 0, nil
}
