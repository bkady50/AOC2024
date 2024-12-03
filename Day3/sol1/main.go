package main

import (
	"strings"
	"fmt"
	"os"
	"strconv"
)

func addMul(mul string) float64 {
	//fmt.Println(mul)

	pre, found := strings.CutPrefix(mul, "(")

	if !found {
		return 0
	}

	//fmt.Println(pre)

	suf, _, found := strings.Cut(pre, ")")

	if !found {
		return 0
	}

	//fmt.Println(suf)

	split := strings.Split(suf, ",")

	//fmt.Println(split)

	if len(split) != 2 {
		return 0
	}

	first, err := strconv.ParseFloat(split[0], 64)
	if err != nil {
        return 0
    }

	second, err := strconv.ParseFloat(split[1], 64)
	if err != nil {
        return 0
    }

	//fmt.Println(first * second)

	return first * second
} 

func main() {
    data, err := os.ReadFile("input.txt")
	if err != nil {
        panic(err)
    }
    //fmt.Print(string(data))

	program := string(data)

	count := 0.0

	_, after, found := strings.Cut(program, "mul")

	//fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", "string", "mul", before, after, found)

	for found {
		count += addMul(after)
		_, after, found = strings.Cut(after, "mul")
	}

	fmt.Println(count)
}