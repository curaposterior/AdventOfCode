package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func convert_list_string_to_list_int(list []string) ([]int, error) {
	list_int := make([]int, len(list))

	for i, s := range list {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		list_int[i] = num
	}
	return list_int, nil
}

func read_data(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	matrix := make([][]int, 0)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		row, err := convert_list_string_to_list_int(line)
		if err != nil {
			log.Fatal(err)
		}
		matrix = append(matrix, row)
	}
	for err := scanner.Err(); err != nil; {
		log.Fatal(err)
	}
	return matrix
}

func check_difference_condition(list []int) bool {
	var subs int
	for i := 0; i < len(list)-1; i++ {
		subs = list[i] - list[i+1]
		if math.Abs(float64(subs)) > 3 {
			return false
		}
		if math.Abs(float64(subs)) < 1 {
			return false
		}
	}
	return true
}

func is_increasing(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] >= list[i+1] {
			return false
		}
	}
	return true
}

func is_decreasing(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] <= list[i+1] {
			return false
		}
	}
	return true
}

func is_safe(list []int) bool {
	increasing := is_increasing(list)
	decreasing := is_decreasing(list)
	if increasing || decreasing {
		diff := check_difference_condition(list)
		if diff {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func is_safe_2(list []int) bool {
	if is_safe(list) {
		return true
	}

	for i := 0; i < len(list); i++ {
		// this is kinda sketchy, but works
		modifiedList := make([]int, 0, len(list)-1)
		modifiedList = append(modifiedList, list[:i]...)
		modifiedList = append(modifiedList, list[i+1:]...)

		fmt.Println(modifiedList)
		if is_increasing(modifiedList) || is_decreasing(modifiedList) {
			if check_difference_condition(modifiedList) {
				return true
			}
		}
	}
	return false
}

func part_1(filename string) int {
	number_of_safe := 0

	list_of_levels := read_data(filename)
	for _, element := range list_of_levels {
		if is_safe(element) {
			number_of_safe++
		}
	}

	return number_of_safe
}

func part_2(filename string) int {
	number_of_safe := 0

	list_of_levels := read_data(filename)
	for _, element := range list_of_levels {
		if is_safe_2(element) {
			number_of_safe += 1
		}
	}

	return number_of_safe
}

func main() {
	fmt.Println("PART ONE:")
	fmt.Print("Test: ", part_1("input_help.txt"), ", real: ", part_1("input_puzzle.txt"))
	fmt.Println("\nPART TWO:")
	fmt.Print("Test: ", part_2("input_help.txt"), ", real: ", part_2("input_puzzle.txt"))
}
