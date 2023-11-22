package backend

import (
	"fmt"
	"net/http"
	"os"
)

func StartServer() {
	// Получение порта из окружения
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Говорим маршрутизатору использовать папки для всех путей с префиксами этих папок
	directories := []string{"web", "src", "css", "html", "backend", "dist", "gifs"}

	for _, dir := range directories {
		fs := http.FileServer(http.Dir(dir))
		path := fmt.Sprintf("/%s/", dir)
		http.Handle(path, http.StripPrefix(path, fs))
	}

	// Настройка маршрутизации
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/dir", DirHandler)

	// Запуск сервера
	fmt.Printf(fmt.Sprintf("Server run on http://127.0.0.1:%s\n", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}