package main

import (
	"strings"
	"fmt"
	"os"
	"strconv"
)

func addMul(before string, mul string, cur_op bool) (float64, bool) {
	fmt.Println(mul)

	last_do := strings.LastIndex(before, "do()")
	last_dont := strings.LastIndex(before, "don't()")

	//fmt.Printf("last do: %d, last dont: %d\n", last_do, last_dont)

	if last_dont > last_do {
		//fmt.Printf("skipped\n")
		return 0, false
	} else if last_do > last_dont {
		cur_op = true
	}

	if !cur_op {
		return 0, false
	}

	pre, found := strings.CutPrefix(mul, "(")

	if !found {
		return 0, cur_op
	}

	//fmt.Println(pre)

	suf, _, found := strings.Cut(pre, ")")

	if !found {
		return 0, cur_op
	}

	//fmt.Println(suf)

	split := strings.Split(suf, ",")

	//fmt.Println(split)

	if len(split) != 2 {
		return 0, cur_op
	}

	first, err := strconv.ParseFloat(split[0], 64)
	if err != nil {
        return 0, cur_op
    }

	second, err := strconv.ParseFloat(split[1], 64)
	if err != nil {
        return 0, cur_op
    }

	//fmt.Println(first * second)

	return first * second, cur_op
} 

func main() {
    data, err := os.ReadFile("../input.txt")
	if err != nil {
        panic(err)
    }
    //fmt.Print(string(data))

	program := string(data)

	count := 0.0
	add := 0.0

	cur_op := true

	before, after, found := strings.Cut(program, "mul")

	//fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", "string", "mul", before, after, found)

	for found {
		add, cur_op = addMul(before, after, cur_op)
		count += add
		before, after, found = strings.Cut(after, "mul")
	}

	fmt.Println(count)
}