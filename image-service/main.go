package main

import (
	"log"
	"os"
	"io/ioutil"
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

	for true {

		run()
	}
}

func run() {

	imageFilePath := "image-service.txt"
	imageDirectoryPath := "images"
	randomImagePath := ""
	var randomNumber int = 1

	// read from file
	imageFileData, err := os.ReadFile(imageFilePath)
	if err != nil {
		log.Printf("Error reading file: %v", err)
	}
	imageFileString := string(imageFileData)
	temp, randomNumberErr := strconv.ParseInt(imageFileString, 10, 64)
	for randomNumberErr != nil {
		log.Printf("Could not convert %s to integer. Error: %v \n", imageFileString, randomNumberErr)
		time.Sleep(5 * time.Second)
		imageFileData, err := os.ReadFile(imageFilePath)
		if err != nil {
			log.Printf("Error reading image service file: %v", err)
		}
		imageFileString = string(imageFileData)
		temp, randomNumberErr = strconv.ParseInt(imageFileString, 10, 64)

		if err != nil {
			log.Printf("Error parsing random number: %v", err)
		}
	}
	randomNumber = int(temp)

	_ = randomNumber
	imagePaths, err := ioutil.ReadDir(imageDirectoryPath)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}
	var randomIndex int
	randomIndex = randomNumber % (len(imagePaths) - 1)
	randomImagePath = imagePaths[randomIndex].Name()

	//image service communication
	buffer := []byte(randomImagePath)
	err = os.WriteFile(imageFilePath, buffer, 0644)
	if err != nil {
		log.Print(err)
	}
	log.Print("Success! Wrote file path to image-service.txt")
}
