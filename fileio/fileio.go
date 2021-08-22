package fileio

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func ReadFile(fileName string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("error occurred %s", err)
	}

	filePath = filepath.Join(currentDir, "..", "data", fileName)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error occurred %s", err)
	}

	fmt.Printf("data string\n%s", data)
}