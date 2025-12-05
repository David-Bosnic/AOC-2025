package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bound struct {
	lower  int
	higher int
}

func main() {
	bound := readFile()
	var count int
	for i := range bound {
		fmt.Printf("bound[32].higher: %v\n", bound[31].higher)
		for j := bound[i].lower; j <= bound[i].higher; j++ {
			currentNumber := j
			if isInvalidString(strconv.Itoa(j)) {
				count += currentNumber
			}
		}
	}
	fmt.Printf("count: %v\n", count)
}

func readFile() []Bound {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Fail to read file")
		os.Exit(1)
	}
	var bound []Bound
	parts := strings.Split(string(file), ",")
	for i := range parts {
		lower, higher, _ := strings.Cut(parts[i], "-")
		lowerBound, _ := strconv.Atoi(lower)
		higherBound, _ := strconv.Atoi(higher)
		bound = append(bound, Bound{
			lower:  lowerBound,
			higher: higherBound,
		})
	}
	return bound
}

func isInvalidString(s string) bool {
	l := len(s)
	if l%2 != 0 {
		return false
	}
	return s[:l/2] == s[l/2:]
}
