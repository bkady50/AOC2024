package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isRule(rules []int, number int) bool {
	for i := range(len(rules)) {
		if number == rules[i] {
			return true
		}
	}

	return false
}

func correctUpdate2(update []string, rules map[int][]int) []string {
	index := 1

	for index < len(update) {
		cur_number, err := strconv.Atoi(update[index])

		if err != nil {
			panic(err)
		}

		cur_rules := rules[cur_number]

		swapped := false

		i := index - 1
		for i >= 0 {
			// if update[i] is in cur_rules swap and set index to i
			compare, err := strconv.Atoi(update[i])

			if err != nil {
				panic(err)
			}

			swapped = isRule(cur_rules, compare)

			if swapped {
				//fmt.Printf("update[%d]=%s, update[%d]=%s\n", index, update[index], i, update[i])
				tmp := update[index]
				update[index] = update[i]
				update[i] = tmp
				index = i
				break
			}
			i--
		}
		if !swapped {
			index++
		}

	}

	//fmt.Println(update)
	return update
}

func correctUpdate(update []string, rules map[int][]int) []string {
	for i := range(len(update)) {
		i_number, err := strconv.Atoi(update[i])
		if err != nil {
			panic(err)
		}

		cur_rules := rules[i_number]

		for j := range(len(update)) {
			j_number, err := strconv.Atoi(update[j])

			if err != nil {
				panic(err)
			}

			for k := range(len(cur_rules)) {
				if j_number == cur_rules[k] && i < j {
					fmt.Println(update)
					//fmt.Printf("update[%d]=%d, update[%d]=%d\n", i, i_number, j, j_number)


					temp := i_number

					update[i] = strconv.Itoa(j_number)
					update[j] = strconv.Itoa(temp)
					fmt.Println(update)
				}
			}
		}
	} 
	return update
}

// take current update, index into update, rules, previous numbers
func isBad(prev_numbers []int, cur_rules []int) bool {
	for i := range(prev_numbers) {
		for j := range(cur_rules) {
			if prev_numbers[i] == cur_rules[j] {
				return true
			}
		}
	}
	return false
}

func addMid(update []string) int {
	mid := len(update) / 2

	val, err := strconv.Atoi(update[mid])

	if err != nil {
		panic(err)
	}

	//fmt.Printf("string: %s adding: %d\n", update, val)

	return val
}

func main () {
	file, err := os.Open("input.txt")
    if err != nil {
		panic(err)
	}
	defer file.Close()

	rules := make(map[int][]int)
	var updates [][]string

	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			rule := strings.Split(line, "|")

			key, err := strconv.Atoi(rule[0])
			if err != nil {
				panic(err)
			}
			
			val, err := strconv.Atoi(rule[1])
			if err != nil {
				panic(err)
			}

			_, prs := rules[key]

			if !prs {
				rules[key] = make([]int, 0)
				rules[key] = append(rules[key], val)
			} else {
				rules[key] = append(rules[key], val)
			}
		} else if strings.Contains(line, ",") {
			update := strings.Split(line, ",")
			updates = append(updates, update)
			//fmt.Println(update)
		}
    }

	//fmt.Println(rules)

	count := 0

	// look through each update
	for i := range(len(updates)) {
		var prev_numbers []int
		cur_update := updates[i]

		// look through all the numbers in an update
		for j := range(len(cur_update)) {
			cur_number, err := strconv.Atoi(cur_update[j])
			cur_rules := rules[cur_number]

			if err != nil {
				panic(err)
			}

			// take current update, index into update, rules, previous number
			if len(prev_numbers) > 0 {
				if isBad(prev_numbers, cur_rules) {
					count += addMid(correctUpdate2(cur_update, rules))
					break
				}
			} 

			prev_numbers = append(prev_numbers, cur_number)
		}
	}

	fmt.Println(count)

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}