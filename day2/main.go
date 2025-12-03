package main

import (
	"bufio"
	"fmt"
	"os"
)

type Bound struct {
	lower  int
	higher int
}

func main() {
}

func readFile() []Bound {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Fail to read file")
		os.Exit(1)
	}
	fmt.Println(file)
}
