package main

import (
	"fmt"
	"github.com/iamlucasvieira/aoc/utils"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// distance Calculates the distance a boat will travel in a given time at a given speed.
func distance(time, speed int) int {
	return (time - speed) * speed
}

// findSpeed Finds the speed a boat must travel to cover a given distance in a given time.
// Solves: d = -s^2 + ts
func findSpeed(time, distance int) (float64, float64, error) {
	speed1, speed2, err := findRoots(-1, float64(time), float64(-distance))
	if err != nil {
		return 0, 0, err
	}
	return speed1, speed2, nil
}

// findRoots Finds the real roots of a quadratic equation.
func findRoots(a, b, c float64) (float64, float64, error) {
	delta := b*b - 4*a*c

	// If delta is negative, there are no real roots.
	if delta < 0 {
		return 0, 0, fmt.Errorf("no real roots")
	} else if delta == 0 {
		// If delta is zero, there is one real root.
		root := -b / (2 * a)
		return root, root, nil
	} else {
		// If delta is positive, there are two real roots.
		return (-b + math.Sqrt(delta)) / (2 * a), (-b - math.Sqrt(delta)) / (2 * a), nil
	}
}

// bestSpeeds Finds the best speeds for a boat to travel a higher distance in a given time.
func bestSpeeds(time, distance int) ([]int, error) {
	speed1, speed2, err := findSpeed(time, distance)
	if err != nil {
		return nil, err
	}

	minSpeed := math.Min(speed1, speed2)
	maxSpeed := math.Max(speed1, speed2)

	if minSpeed < 0 {
		minSpeed = 0
	}

	// Find all integer speeds between minSpeed and maxSpeed
	var speeds []int

	for i := int(minSpeed) + 1; i < int(math.Ceil(maxSpeed)); i++ {
		speeds = append(speeds, i)
	}

	return speeds, nil
}

// totalBestSpeeds Finds the total number of best speeds for a boat to travel a higher distance in a given time.
func totalBestSpeeds(time, distance int) (int, error) {
	speeds, err := bestSpeeds(time, distance)
	if err != nil {
		return 0, err
	}
	return len(speeds), nil
}

type race struct {
	time, distance int
}

// parseData Parses the data from the input file.
func parseData(data []string) []race {
	var times, distances []int

	// Define regex for parsing numbers
	numberRegex := regexp.MustCompile(`\d+`)
	for _, line := range data {

		// Create list with integers found
		var listNumbers []int
		for _, number := range numberRegex.FindAllString(line, -1) {
			numberInt, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			listNumbers = append(listNumbers, numberInt)
		}
		if strings.HasPrefix(line, "Time:") {
			times = listNumbers
		} else if strings.HasPrefix(line, "Distance:") {
			distances = listNumbers
		}
	}

	var allRaces = make([]race, len(times))

	for idx := range times {
		allRaces[idx] = race{times[idx], distances[idx]}
	}

	return allRaces
}

// buildSingleRace appends all integers to form a single time and distance
func buildSingleRace(races []race) race {

	var time, distance string

	for _, r := range races {
		// convert int to string
		time += strconv.Itoa(r.time)
		distance += strconv.Itoa(r.distance)
	}

	// convert string to int
	timeInt, err := strconv.Atoi(time)
	if err != nil {
		panic(err)
	}

	distanceInt, err := strconv.Atoi(distance)
	if err != nil {
		panic(err)
	}

	return race{timeInt, distanceInt}
}

// productNumberBestSpeeds returns the product of the number of best speeds in each race
func productNumberBestSpeeds(races []race) int {
	var product = 1

	for _, r := range races {
		speeds, err := bestSpeeds(r.time, r.distance)
		if err != nil {
			panic(err)
		}
		product *= len(speeds)
	}

	return product
}

func part1() {
	fmt.Println("Part 1:")
	data := parseData(utils.ReadFile("input.txt"))
	product := productNumberBestSpeeds(data)
	fmt.Printf("Product: %d\n", product)
}

func part2() {
	fmt.Println("Part 2:")
	data := parseData(utils.ReadFile("input2.txt"))
	singleRace := buildSingleRace(data)
	total, err := totalBestSpeeds(singleRace.time, singleRace.distance)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of best speeds: %d\n", total)
}

func main() {
	part1()
	part2()
}
