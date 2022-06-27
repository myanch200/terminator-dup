package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Write the folder name: ")
	var filePath string

	fmt.Scan(&filePath)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	filePath = homeDir + "/" + filePath

	files, err := os.ReadDir(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fileNames := make(map[string]string)

	for _, f := range files {

		if fileNames[f.Name()] == "" {
			fileNames[f.Name()] = f.Name()
		} else {
			fmt.Println(f.Name() + " is a duplicate")
		}

	}
	fmt.Println("The files in the folder are: ")
	for _, f := range fileNames {
		fmt.Println(f)
	}

}
