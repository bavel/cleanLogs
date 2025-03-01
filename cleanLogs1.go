package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func deleteOldFiles(dir string, daysOld int) error {
	// Получаем текущее время
	now := time.Now()

	// Проходим по всем файлам в директории
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Проверяем, является ли файл .zip или .log
		ext := filepath.Ext(path)
		if ext == ".zip" || ext == ".log" {
			// Проверяем, старше ли файл указанного количества дней
			if now.Sub(info.ModTime()).Hours() > float64(daysOld*24) {
				fmt.Printf("Удаление файла: %s\n", info.Name())
				os.Remove(path)
			}
		}
		return nil
	})

	return err
}

func main() {
	// Получаем путь к AppData\Roaming текущего пользователя
	appDataDir := os.Getenv("APPDATA")
	if appDataDir == "" {
		fmt.Println("Ошибка: не удалось найти каталог AppData.")
		return
	}

	// Указываем путь к каталогу с логами
	logDir := filepath.Join(appDataDir, "iiko", "CashServer", "Logs")

	// Количество дней, после которых файлы считаются старыми
	daysOld := 30

	// Удаляем старые файлы
	err := deleteOldFiles(logDir, daysOld)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Удаление файлов старше %d дней завершено.\n", daysOld)
	}
}
