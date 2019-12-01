package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CalculateSumFuel() {
	masses := readFile("./day1/input-star1")
	requiredFuel := 0
	requiredFuelPart2 := 0
	for _, mass := range masses {
		requiredFuel += calculateRequiredFuel(mass)
		requiredFuelPart2 += calculateRequiredFuelPartTwo(mass)
	}
	fmt.Printf("Part1 : Requires %d fuel.\n", requiredFuel)
	fmt.Printf("Part2 : Requires %d fuel.\n", requiredFuelPart2)
}

func calculateRequiredFuel(mass int) int {
	return (mass / 3) - 2
}

func calculateRequiredFuelPartTwo(mass int) int {
	requiredFuel := 0
	for calculateRequiredFuel(mass) > 0 {
		mass = calculateRequiredFuel(mass)
		requiredFuel += mass
	}
	return requiredFuel
}

func readFile(input string) []int {
	file, err := os.Open(input)
	if err != nil {
		fmt.Printf("fail to read file " + input + " and got: " + err.Error())
	}
	defer file.Close()

	totalmasses := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("fail to convert string to int and got: " + err.Error())
		}
		totalmasses = append(totalmasses, mass)
	}

	return totalmasses
}
