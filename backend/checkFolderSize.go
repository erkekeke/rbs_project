package backend

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

// checkFolderSize() Установление размера директории.
func CheckFolderSize(ctx context.Context, wg *sync.WaitGroup, sizeCh chan int64, root string, unsortedFiles *[]File, deepIndex int) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		fmt.Println("Пользователь вышел. Прекращение сканирования")
		sizeCh <- 0
		return
	default:
	}

	files, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Println("Неверный адрес директории")
		sizeCh <- 0
		return
	}

	var size int64

	// Обход по файлам
	for _, file := range files {
		if !file.IsDir() {
			*unsortedFiles = append(*unsortedFiles, File{file.Name(), file.Size(), "file", deepIndex})
			size += file.Size()
		} else {
			newRoot := fmt.Sprintf("%s/%s/", root, file.Name())

			innerSizeCh := make(chan int64)
			wg.Add(1)
			go CheckFolderSize(ctx, wg, innerSizeCh, newRoot, unsortedFiles, deepIndex+1)

			subDirSize := <-innerSizeCh
			size += subDirSize
		}
	}

	pathArray := strings.Split(root, "/")
	*unsortedFiles = append(*unsortedFiles, File{pathArray[len(pathArray)-2], size, "folder", deepIndex - 1})

	sizeCh <- size
}
