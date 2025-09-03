package noteHttp

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Rules struct {
	Name        string
	Time        string
	Cause       string
	CurrentTime string
}

func Note(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println("Ошибка html:", err.Error())
		http.Error(w, "Ошибка загрузки шаблона", http.StatusInternalServerError)
		return
	}

	currentTime := time.Now()

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		cause := r.FormValue("cause")
		timeInput := r.FormValue("time")

		var timeStr string
		if timeInput == "" {
			// Если время не указано, прибавляем 40 секунд к текущему времени
			futureTime := currentTime.Add(40 * time.Second)
			timeStr = futureTime.Format("02.01.2006 15:04:05")
		} else {
			timeStr = timeInput
		}

		death := Rules{
			Name:        name,
			Time:        timeStr,
			Cause:       cause,
			CurrentTime: currentTime.Format("02.01.2006 15:04:05"),
		}
		tmpl.Execute(w, death)
		return
	}

	// Для GET запроса
	tmpl.Execute(w, Rules{CurrentTime: currentTime.Format("02.01.2006 15:04:05")})
}
