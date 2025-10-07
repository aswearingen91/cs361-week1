package main

import (
	"log"
	"os"
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
4. Read from that file a path to an image
5. Output the image path to console.

*/

func main() {

	prngFilePath := "prng-service.txt"
	imageFilePath := "image-service.txt"

	//prng communication
	data := []byte("run")
	err := os.WriteFile(prngFilePath, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	prngNumber, err := os.ReadFile(prngFilePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	//image service communication
	//data := []byte("run")
	err = os.WriteFile(imageFilePath, prngNumber, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
