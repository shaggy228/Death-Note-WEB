package main

import (
	"fmt"
	"net/http"
	"note/noteHttp"
)

func main() {
	// Обслуживаем статические файлы из папки static
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Сервер запущен на http://localhost:9091/note")
	http.HandleFunc("/note", noteHttp.Note)
	http.ListenAndServe(":9091", nil)
}
