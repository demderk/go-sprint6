package handlers

import (
	"fmt"
	serverLogs "go-sprint6/internal/logs"
	"go-sprint6/internal/service"
	"io"
	"net/http"
	"os"
	"time"
)

func GetMain(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	indexPage, err := os.ReadFile("index.html")
	if err != nil {
		serverLogs.Main.Print(fmt.Errorf("can't open index.html file: %w", err))
		http.Error(response, "File not found", http.StatusNotFound)
		return
	}

	response.Header().Add("Content-Type", "text/html; charset=utf-8")
	_, err = response.Write(indexPage)
	if err != nil {
		serverLogs.Main.Print(fmt.Errorf("response write caught an error: %w", err))
		http.Error(response, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func PostUpload(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	request.ParseMultipartForm(10 << 20)

	file, _, err := request.FormFile("myFile")
	if err != nil {
		serverLogs.Main.Println(fmt.Errorf("form file raised an error: %w", err))
		http.Error(response, "Bad Request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		serverLogs.Main.Println(fmt.Errorf("can't read file: %w", err))
		http.Error(response, "Bad Request", http.StatusBadRequest)
		return
	}

	convertedText := service.ConvertText(string(fileBytes))
	if err := writeLocalResult(convertedText); err != nil {
		serverLogs.Main.Print(fmt.Errorf("can't write local result: %w", err))
		http.Error(response, "Internal server error", http.StatusInternalServerError)
		return
	}

	_, err = response.Write([]byte(convertedText))
	if err != nil {
		serverLogs.Main.Print(fmt.Errorf("response write caught an error: %w", err))
		http.Error(response, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func writeLocalResult(text string) error {
	resultFile, err := makeFile()
	if err != nil {
		return fmt.Errorf("can't make result file: %w", err)
	}

	if _, err := io.WriteString(resultFile, text); err != nil {
		return fmt.Errorf("can't write to result file: %w", err)
	}

	return nil
}

func makeFile() (*os.File, error) {
	if _, err := os.Stat("convertedFiles"); os.IsNotExist(err) {
		serverLogs.Main.Printf("creating directory convertedFiles")
		if err := os.Mkdir("convertedFiles", 0755); err != nil {
			err := fmt.Errorf("can't create directory convertedFiles: %w", err)
			return nil, err
		}
	}

	root, err := os.OpenRoot("convertedFiles")
	if err != nil {
		return nil, fmt.Errorf("can't open directory convertedFiles: %w", err)
	}

	fileName := fmt.Sprintf("%v.txt", time.Now().UTC().String())

	file, err := root.Create(fileName)
	if err != nil {
		return nil, fmt.Errorf("can't create file %q: %w", fileName, err)
	}

	return file, nil
}
