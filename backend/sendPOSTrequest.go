package backend

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func SendPOSTRequest(postData File) {
	client := &http.Client{}
	phpScriptURL := "http://localhost:80/setStat.php"

	jsonData, err := json.Marshal(postData)
	if err != nil {
		log.Println("Файл: sendPOSTrequest.go. Строка:17. Ошибка при маршалинге данных для PHP-запроса")
	}

	reqPHP, err := http.NewRequest("POST", phpScriptURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Файл: sendPOSTrequest.go. Строка:22. Ошибка при создании HTTP-запроса")
		return
	}
	reqPHP.Header.Set(
		"Content-type",
		"application/json",
	)

	resp, err := client.Do(reqPHP)
	if err != nil {
		log.Println("Файл: sendPOSTrequest.go. Строка:32. Ошибка при выполнении HTTP-запроса")
	}
	defer resp.Body.Close()

	// Ответ PHP
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Файл: sendPOSTrequest.go. Строка:39. Ошибка при конвертации ответа от PHP в string")
		return
	}
	respBody := string(body)
	log.Println("Файл: sendPOSTrequest.go. Строка:43. PHP-скрипт вернул: ", respBody)
}
