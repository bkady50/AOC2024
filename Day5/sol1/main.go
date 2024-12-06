package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addMid(update []string) int {
	mid := len(update) / 2

	val, err := strconv.Atoi(update[mid])

	if err != nil {
		panic(err)
	}

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
		}
    }

	count := 0

	for i := range(len(updates)) {
		var prev_numbers []int
		valid := true
		cur_update := updates[i]

		for j := range(len(cur_update)) {
			cur_number, err := strconv.Atoi(cur_update[j])
			cur_rules := rules[cur_number]

			if err != nil {
				panic(err)
			}

			if len(prev_numbers) > 0 {
				for k := range(prev_numbers) {
					for l := range(cur_rules) {
						if prev_numbers[k] == cur_rules[l] {
							valid = false
						}
					}
				}
			} 

			prev_numbers = append(prev_numbers, cur_number)
		}

		if valid {
			count += addMid(cur_update)
		}
	}

	fmt.Println(count)

    if err := scanner.Err(); err != nil {
        panic(err)
    }
}