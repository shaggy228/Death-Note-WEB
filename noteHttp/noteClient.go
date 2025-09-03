package noteHttp

import (
	"fmt"
	"html/template"
	"net/http"
)

type Rules struct {
	Name string

	Time string

	Cause string
}

func Note(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println("Ошибка html:", err.Error())
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		cause := r.FormValue("cause")
		time := r.FormValue("time")
		death := Rules{Name: name, Time: time, Cause: cause}
		tmpl.Execute(w, death)
		return
	}

	tmpl.Execute(w, Rules{})
}
