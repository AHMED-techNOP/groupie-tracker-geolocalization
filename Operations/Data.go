package groupie




type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}


type ArtistsPageData struct {
	Artists []Artist
}

var Artists []Artist // Artists should be a slice of Artist


type ArtistDetailPageData struct {
	Artist    Artist
	Locations []string
	Dates     []string
	Relations map[string][]string
	Latitude float64
	Longitude float64
}

