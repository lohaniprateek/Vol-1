package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dir := "contribution"

		// read all files in directory
		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Fprintf(w, "Error reading directory: %v", err)
			os.Exit(1)
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			filePath := filepath.Join(dir, file.Name())

			data, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Fprintf(w, "Error reading file %s: %v\n", filePath, err)
				continue
			}

			content := string(data)

			if strings.Contains(content, "@") {
				start := strings.Index(content, "@")
				end := strings.Index(content[start:], " ")
				end = start + end
				uname := content[start:end]
				fmt.Fprintf(w, "Hello from %s, vol-1 Maine Vol-1 likha hi nhi hai\n", uname)
			} else {
				fmt.Fprintf(w, "not found in %s\n", file.Name())
			}
		}
	})

	http.ListenAndServe(":8080", nil)
}
