package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Получение порта из окружения
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Говорим маршрутизатору использовать папку templates для всех путей с префиксом /templates/
	fs := http.FileServer(http.Dir("templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))

	fs = http.FileServer(http.Dir("src/"))
	http.Handle("/src/", http.StripPrefix("/src/", fs))

	fs = http.FileServer(http.Dir("css/"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	fs = http.FileServer(http.Dir("html/"))
	http.Handle("/html/", http.StripPrefix("/html/", fs))

	// Настройка маршрутизации
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/dir", DirHandler)

	// Запуск сервера
	fmt.Printf(fmt.Sprintf("Server run on http://127.0.0.1:%s\n", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
