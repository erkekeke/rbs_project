package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sort"
	"sync"
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

	// Создание контекста для выхода при сканировании
	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()

	// Инициализация WaitGroup
	var wg sync.WaitGroup

	wg.Add(1)
	go CheckFolderSize(ctx, &wg, sizeCh, root, &unsortedFiles, 0)
	defer close(sizeCh)
	wg.Wait()

	if len(unsortedFiles) > 0 {
		unsortedFiles = unsortedFiles[:len(unsortedFiles)-1]
	}
	// Сортировка в зависимости от её типа(desc, asc)
	if sortType == "asc" {
		sort.Sort(FilesArray(unsortedFiles))
	} else if sortType == "desc" {
		sort.Sort(sort.Reverse(FilesArray(unsortedFiles)))
	} else {
		fmt.Println("Неверный тип сортировки. Используйте 'asc' или 'desc'.")
		os.Exit(1)
	}

	// var resp = Response{500, "Произошла ошибка на сервере", nil}
	// Конвертация данных в формат json и их вывод в клиент

	jsonData, err := json.Marshal(unsortedFiles)

	if err != nil {
		res.Write(nil)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.Write(jsonData)
}

// indexHandler() Обрабатывает рут: /
func IndexHandler(res http.ResponseWriter, req *http.Request) {

	tpl, err := template.ParseFiles("templates/html/index.html")
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
