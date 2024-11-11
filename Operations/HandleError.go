package groupie

import (
	"net/http"
	"text/template"
)

func HandleError(w http.ResponseWriter, r *http.Request, code int) {
	message := ""

	switch code {
	case 400:
		message = "Bad Request"
	case 403:
		message = "Forbidden Page"
	case 404:
		message = "Oops! Page Not Found"
	case 405:
		message = "Method Not Allowed"
	case 500:
		message = "Internal Server Error"
	}

	w.WriteHeader(code)

	tmp, err := template.ParseFiles("Error/Err.html")
	if err != nil {
		w.WriteHeader(500)
		http.ServeFile(w, r, "Error/500.html")
		return
	}

	data := struct {
		M string
		C int
	}{
		M: message,
		C: code,
	}

	err = tmp.Execute(w, data)
	if err != nil {
		w.WriteHeader(500)
		http.ServeFile(w, r, "Error/500.html")
		return
	}
}
