package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

// This will be the UI program
/*
It will call a PRNG service by writing the word "run" to a file called
prng-service.txt
The PRNG service will overwrite the contents of that file
with a pseudo-random number
The UI service will then read that random number from the prng-service.txt file
The UI will then write that number to a file 'image-service.txt'

The image service will read that file and overwrite the file with a path to an image
The UI will then read that file and print the path found in that file
*/

/*
UI's responsibilities
1. Write the word run to prng-service.txt
2. Read a number from that file
3. Write that number to a file called 'image-service.txt'
TODO: 4. Read from that file a path to an image
TODO 5. Output the image path to console.

*/

func main() {

	prngFilePath := "prng-service.txt"
	imageFilePath := "image-service.txt"

	//prng communication
	data := []byte("run")
	// Might need to change file mode to 0666
	// Should be fine since all programs will be running as
	// my user
	err := os.WriteFile(prngFilePath, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// should not proceed if contents are not a number
	prngFileData, err := os.ReadFile(prngFilePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	prngFileString := string(prngFileData)
	prngNumber, err := strconv.ParseInt(prngFileString, 10, 64)
	log.Print("Failed to parse prng number")
	for err != nil {
		log.Print("Waiting for prng number to be parseable.")
		// Chill out
		time.Sleep(5 * time.Second)
		prngFileData, err = os.ReadFile(prngFilePath)
		if err != nil {
			log.Printf("Error reading file: %v", err)
		}

		prngFileString = string(prngFileData)
		log.Printf("Prng file contents: %s", prngFileString)
		prngNumber, err = strconv.ParseInt(prngFileString, 10, 64)
		if err != nil {
			log.Printf("Error parsing prng number: %v", err)
		}
	}

	//image service communication
	//data := []byte("run")
	buffer := make([]byte, prngNumber)
	err = os.WriteFile(imageFilePath, buffer, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
