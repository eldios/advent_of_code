package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	inputFile := flag.String("i", "input.txt", "Input file with new line separated values")

	flag.Parse()

	fmt.Println(*inputFile)

	// Open file and create scanner on top of it
	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	// Scan for next line
	totalFuel := int64(0)
	for scanner.Scan() {
		// False on error or EOF. Check error
		err = scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		moduleMass, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}

		requiredFuel := int64(math.Floor(moduleMass/3) - 2)

		totalFuel = totalFuel + requiredFuel
	}

	fmt.Println(totalFuel)
}
