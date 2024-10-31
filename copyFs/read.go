package main

import (
	"fmt"
	"os"
)

func getAllFileNames(filepath string) {
	files, err := os.ReadDir(filepath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			// Recursively call for subdirectories
			getAllFileNames(filepath + file.Name() + "/") // Ensure correct path with '/'
			continue // Skip further processing for directories
		}

		// Read file content
		data, err := os.ReadFile(filepath + file.Name())
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue // Skip to next file on error
		}

		// Print the filename and its content
		fmt.Printf("Data from %s:\n%s\n", file.Name(), string(data))
		fmt.Println("-------------------------") // Separator for clarity
	}
}

func ReadFiles() {
	source := "/home/lemon/Desktop/Go/AllGO"
	getAllFileNames(source) // Start reading from the source directory
}

func main() {
	ReadFiles()
}
