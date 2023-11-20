package backend

// Структура для файлов в родительской директории
type File struct {
	FileName  string `json:"fileName"`
	FileSize  int64  `json:"fileSize"`
	FileType  string `json:"fileType"`
	DeepIndex int    `json:"deepIndex"`
}
type FilesArray []File

// Структура для ответа сервера
type Response struct {
	Status    int64
	ErrorText string
	Data      []File
}

// функции для интерфейса sort
func (object FilesArray) Len() int {
	return len(object)
}
func (object FilesArray) Less(i, j int) bool {
	return object[i].FileSize < object[j].FileSize
}
func (object FilesArray) Swap(i, j int) {
	object[i], object[j] = object[j], object[i]
}
