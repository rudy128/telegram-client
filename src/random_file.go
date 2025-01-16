package src

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func getRandomFileFromDir(dir string) (string, error) {
	// Open the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("could not read directory: %v", err)
	}

	// Filter out directories, keeping only files
	var fileList []string
	for _, file := range files {
		if !file.IsDir() {
			fileList = append(fileList, file.Name())
		}
	}

	// If there are no files, return an error
	if len(fileList) == 0 {
		return "", fmt.Errorf("no files found in the directory")
	}

	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// Select a random index
	randomIndex := rand.Intn(len(fileList))

	// Construct the full path of the random file
	randomFilePath := filepath.Join(dir, fileList[randomIndex])

	// Return the full file path
	return randomFilePath, nil
}
