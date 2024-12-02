package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Mins(values []int) (int, int) {
	if len(values) == 0 {
		return 0, -1
	}

	value := values[0]
	to_del := 0
	for index, v := range values {
		if v < value {
			value = v
			to_del = index
		}
	}
	return value, to_del
}

func ParseInputData(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	left := make([]int, 0)
	right := make([]int, 0)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		num1, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}
		num2, err := strconv.Atoi(line[1])
		if err != nil {
			log.Fatal(err)
		}

		left = append(left, num1)
		right = append(right, num2)
	}

	for err := scanner.Err(); err != nil; {
		log.Fatal(err)
	}
	return left, right
}

func part_1(filename string) int {
	// this is definitely not efficient, but it does the job
	// (+ go is pretty fast so I guess it's not that bad)
	left, right := ParseInputData(filename)
	final_sum := 0
	for range left {
		left_min, first_del_index := Mins(left)
		right_min, second_del_index := Mins(right)

		var substraction int
		if left_min > right_min {
			substraction = left_min - right_min
		} else {
			substraction = right_min - left_min
		}
		final_sum += substraction

		left = append(left[:first_del_index], left[first_del_index+1:]...)
		right = append(right[:second_del_index], right[second_del_index+1:]...)
	}
	return final_sum
}

func similarity_score(number int, list []int) int {
	occurences := 0
	for i := range list {
		if list[i] == number {
			occurences += 1
		}
	}
	return occurences * number
}

func part_2(filename string) int {
	// this can probably be done faster with a lookup table or something
	left, right := ParseInputData(filename)
	final_sum := 0
	for _, item := range left {
		final_sum += similarity_score(item, right)
	}
	return final_sum
}

func main() {
	fmt.Println("PART ONE:")
	fmt.Print("Test: ", part_1("input_help.txt"), ", real: ", part_1("input_puzzle.txt"))
	fmt.Println("\nPART TWO:")
	fmt.Print("Test: ", part_2("input_help.txt"), ", real: ", part_2("input_puzzle.txt"))
}
