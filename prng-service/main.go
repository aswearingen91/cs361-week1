package main

import (
	"log"
	"os"
	"time"
	"math/rand"
	"strconv"
)

// This will be the random number program
/*
It reads the word run from prng-service.txt
The PRNG service will overwrite the contents of that file with a pseudo-random number
The UI service will then read that random number from the prng-service.txt file
*/

/*
Random number programs's responsibilities
1. It reads the word run from prng-service.txt
2. Overwrite the contents of that file with a pseudo-random number
*/

func main() {
	for true {

		run()
	}
}

func run() {
	prngFilePath := "prng-service.txt"
	randomNumber := 1

	prngFileData, err := os.ReadFile(prngFilePath)
	if err != nil {
		log.Printf("Error reading file: %v", err)
	}
	for err != nil {
		prngFileData, err = os.ReadFile(prngFilePath)
		// Chill out
		time.Sleep(5 * time.Second)
	}

	prngFileString := string(prngFileData)

	for prngFileString != "run" {
		log.Print("Waiting for prng-service.txt to contain 'run' \n.")

		// Chill out
		time.Sleep(5 * time.Second)

		prngFileData, err = os.ReadFile(prngFilePath)
		if err != nil {
			log.Printf("Error reading file: %v", err)
		}

		prngFileString = string(prngFileData)

	}
	// generate random number
	rand.Seed(time.Now().UnixNano())

	randomNumber = rand.Intn(1024)

	// write random number

	data := []byte(strconv.Itoa(randomNumber))

	err = os.WriteFile(prngFilePath, data, 0644)
	if err != nil {
		log.Print(err)
	}
	log.Print("Success! Wrote random number to image-service	.txt")

}
