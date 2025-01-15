package main

import (
	"fmt"
	"os"
)

func checkAndCreateFolder(folderpath string) {
	_, err := os.Stat(folderpath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(folderpath, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
		fmt.Println("Folder created:", folderpath)
	} else if err != nil {
		fmt.Println("Error checking folder:", err)
		return
	} else {
		fmt.Println("Folder already exists.", folderpath)
	}
	return
}
