package internal

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
)

func GenerateRandomChareters(archive string, characters string) {

	fileName := archive
	size, err := strconv.Atoi(characters)
	if err != nil {
		fmt.Println("Error converting size to integer:", err)
		return
	}

	// Open file for writing
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Generate random bytes
	randomBytes := make([]byte, size)
	_, err = rand.Read(randomBytes)
	if err != nil {
		fmt.Println("Error generating random bytes:", err)
		return
	}

	// Write random bytes to file
	_, err = file.Write(randomBytes)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Successfully wrote", size, "random bytes to", fileName)
}
