package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sort"
	"sync"
	"path"
)

// DirHandler() Обрабатывает рут: /dir
func DirHandler(res http.ResponseWriter, req *http.Request) {

	// Установка заголовка
	res.Header().Set(
		"Content-type",
		"application/json",
	)

	// Получение параметров из url
	reqParams := req.URL.Query()
	root := reqParams.Get("root")
	sortType := reqParams.Get("sort")

	fmt.Println("sort: ", sortType)
	fmt.Println("root: ", root)

	if sortType == "" {
		sortType = "asc"
	}

	// Объявление среза файлов и многопоточный запуск функции
	sizeCh := make(chan int64, 1)
	var unsortedFiles []File
	response := Response{500, "Ошибка при конвертации массива в json-формат", unsortedFiles}

	// Создание контекста для выхода при сканировании
	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()

	// Инициализация WaitGroup
	var wg sync.WaitGroup

	// Запуск горутины
	wg.Add(1)

	go CheckFolderSize(ctx, &wg, sizeCh, root, &response.Data, 0)
	defer close(sizeCh)

	wg.Wait()

	if len(response.Data) > 0 {
		response.Data = response.Data[:len(response.Data)-1]
	}
	// Сортировка в зависимости от её типа(desc, asc)
	if sortType == "asc" {
		sort.Sort(FilesArray(response.Data))
	} else if sortType == "desc" {
		sort.Sort(sort.Reverse(FilesArray(response.Data)))
	} else {
		fmt.Println("Неверный тип сортировки. Используйте 'asc' или 'desc'.")
		os.Exit(1)
	}

	jsonData, err := json.Marshal(response.Data)

	if err != nil {
		fmt.Println("Ошибка при конвертации данных в формат json")
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = res.Write(jsonData)
	if err != nil {
		fmt.Println("Ошибка при отправке json файлов")
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

// indexHandler() Обрабатывает рут: /
func IndexHandler(res http.ResponseWriter, req *http.Request) {

	// htmlPath := path.Join("web", "html", "index.html")
	htmlPath := path.Join("dist", "index.html")
	tpl, err := template.ParseFiles(htmlPath)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Header().Set(
		"Context-Type",
		"text/html",
	)

	tpl.Execute(res, nil)
}
