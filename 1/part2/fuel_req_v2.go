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

		if *debug {
			fmt.Printf("Module Mass: %v\n", moduleMass)
		}
		totalFuel = totalFuel + calculateFuel(int64(moduleMass), debug)
	}
	if *debug {
		fmt.Printf("Total fuel (delta not included): %v\n", totalFuel)
	}

	fmt.Printf("Total fuel (delta included):     %v\n", totalFuel)
}

func calculateFuel(fuel int64, debug *bool) int64 {
	deltaFuel := int64(math.Floor(float64(fuel)/3) - 2)
	if *debug {
		fmt.Printf("%v -> %v\n", fuel, deltaFuel)
	}
	if deltaFuel > 0 {
		deltaFuel = deltaFuel + calculateFuel(deltaFuel, debug)
		return deltaFuel
	} else {
		return 0
	}
}
