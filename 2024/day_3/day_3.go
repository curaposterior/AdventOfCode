package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func read_data(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	data := string(file)
	return data
}

func part_1(filename string) int {
	final_sum := 0

	data := read_data(filename)
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindAllStringSubmatch(data, -1)
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		final_sum += num1 * num2
	}

	return final_sum
}

func part_2(filename string) int {
	final_sum := 0
	can_multiply := true
	data := read_data(filename)

	// Find all don't() and do() occurrences first
	dontRe := regexp.MustCompile(`don't\(\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	mulRe := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	// Get positions of all matches
	dontMatches := dontRe.FindAllStringIndex(data, -1)
	doMatches := doRe.FindAllStringIndex(data, -1)
	mulMatches := mulRe.FindAllStringSubmatchIndex(data, -1)

	// Create ordered list of all events
	type Event struct {
		position  int
		eventType string // "dont", "do", or "mul"
		match     []int
	}

	events := []Event{}

	for _, match := range dontMatches {
		events = append(events, Event{match[0], "dont", match})
	}
	for _, match := range doMatches {
		events = append(events, Event{match[0], "do", match})
	}
	for _, match := range mulMatches {
		events = append(events, Event{match[0], "mul", match})
	}

	// Sort events by position
	sort.Slice(events, func(i, j int) bool {
		return events[i].position < events[j].position
	})

	// Process events in order
	for _, event := range events {
		switch event.eventType {
		case "dont":
			can_multiply = false
		case "do":
			can_multiply = true
		case "mul":
			if can_multiply {
				// Extract numbers from the multiplication
				nums := mulRe.FindStringSubmatch(data[event.match[0]:event.match[1]])
				if len(nums) == 3 { // Full match + 2 capture groups
					num1, _ := strconv.Atoi(nums[1])
					num2, _ := strconv.Atoi(nums[2])
					final_sum += num1 * num2
				}
			}
		}
	}

	return final_sum
}

// func part_2(filename string) int {
// 	final_sum := 0
// 	can_multiply := true

// 	data := read_data(filename)

// 	for i := 0; i < len(string(data)); i++ {
// 		if i < len(data)-7 {
// 			if data[i:i+7] == "don't()" {
// 				fmt.Println(data[i : i+7])
// 				can_multiply = false
// 			}
// 			if data[i:i+2] == "do()" {
// 				fmt.Println(data[i : i+2])
// 				can_multiply = true
// 			}

// 			if data[i:i+3] == "mul" {
// 				fmt.Println(data[i : i+3])
// 				for j := 0; j < i; j++ {
// 					// check if it is valid mul(number, number) and extract the numbers in mul()
// 				}
// 			}
// 		} else {
// 			//
// 		}

// 	}

// 	return final_sum
// }

func main() {
	fmt.Println("PART ONE:")
	fmt.Println(read_data("input_help.txt"))
	fmt.Print("Test: ", part_1("input_help.txt"), ", real: ", part_1("input_puzzle.txt"))
	// part_2("input_help.txt")
	fmt.Println("\nPART TWO:")
	fmt.Print("Test: ", part_2("input_help.txt"), ", real: ", part_2("input_puzzle.txt"))
}
