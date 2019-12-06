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
	debug := flag.Bool("d", false, "Enable debugging statements")

	flag.Parse()

	if *debug {
		fmt.Printf("Input file: %v\n", *inputFile)
	}

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

		totalFuel = totalFuel + calculateFuel(moduleMass)
	}
	if *debug {
		fmt.Printf("Total fuel (delta not included): %v\n", totalFuel)
	}

	deltaFuel := totalFuel
	for deltaFuel > 0 {
		deltaFuel = calculateFuel(float64(deltaFuel))
		if *debug {
			fmt.Printf("Delta fuel: %v\n", deltaFuel)
		}
		if deltaFuel > 0 {
			totalFuel = totalFuel + deltaFuel
		}
	}

	fmt.Printf("Total fuel (delta included):     %v\n", totalFuel)
}

func calculateFuel(fuel float64) int64 {
	return int64(math.Floor(fuel/3) - 2)
}
