package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

// This will be the image service
/*
The image service will read a number from a file 'image-service.txt'

The image service will read that file and overwrite the file with a path to an image
The UI will then read that file and print the path found in that file
*/

/*
Image service's responsibilities
1. Read from 'image-service.txt'
2. Use the number read from that file to select an image to return to the UI
3. Write the file path for that image to 'image-service.txt'

*/

func main() {

	imageFilePath := "image-service.txt"
	randomImagePath := ""
	var randomNumber int64 = -1

	// read from file
	imageFileData, err := os.ReadFile(imageFilePath)
	if err != nil {
		log.Printf("Error reading file: %v", err)
	}
	imageFileString := string(imageFileData)
	var _, fileStringErr = os.Stat(imageFileString)
	for fileStringErr != nil {
		log.Printf("Image file's string indicates file that doesn't exist %s. Error: %v \n", imageFileString, fileStringErr)
		time.Sleep(5 * time.Second)
		imageFileData, err := os.ReadFile(imageFilePath)
		if err != nil {
			log.Printf("Error reading image service file: %v", err)
		}
		imageFileString = string(imageFileData)
		randomNumber, err = strconv.ParseInt(imageFileString, 10, 64)
		if err != nil {
			log.Printf("Error parsing random number: %v", err)
		}
	}

	_ = randomNumber

	//image service communication
	buffer := []byte(randomImagePath)
	err = os.WriteFile(imageFilePath, buffer, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
