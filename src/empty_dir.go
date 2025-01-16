package src

import (
	"fmt"
	"os"
	"path/filepath"
)

func EmptyDir(dir string) error {
	// Read the contents of the directory
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("could not read directory: %v", err)
	}

	// Iterate over each entry in the directory
	for _, entry := range entries {
		// Construct the full path for the entry
		fullPath := filepath.Join(dir, entry.Name())

		// Remove the file or directory
		if entry.IsDir() {
			// For subdirectories, recursively remove them
			err := os.RemoveAll(fullPath)
			if err != nil {
				return fmt.Errorf("failed to remove directory %s: %v", fullPath, err)
			}
		} else {
			// For files, just remove them
			err := os.Remove(fullPath)
			if err != nil {
				return fmt.Errorf("failed to remove file %s: %v", fullPath, err)
			}
		}
	}

	return nil
}
