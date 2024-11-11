package groupie

import (
	"net/http"
	"os"
	"strings"
)

func StaticHandle(w http.ResponseWriter, r *http.Request) {
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	_, err := os.Stat("." + r.URL.Path)
	if strings.HasSuffix(r.URL.Path, "/") || err != nil {
		HandleError(w,r , 403)
		return
	}
	fs.ServeHTTP(w, r)
}
