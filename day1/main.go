package main
//test
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// apple
const (
	lowerBound = 0
	upperBound = 99
	startPos   = 50
)

type dialMotion struct {
	direction string
	value     int
}

func main() {
	seq := readFileGetSeq()

	var counter int
	currPos := startPos
	for i := range len(seq) {
		if seq[i].direction == "L" {
			currPos -= seq[i].value
		} else {
			currPos += seq[i].value
		}
	TurnOver:
		if currPos >= upperBound+100 {
			currPos -= 100
			counter++
			goto TurnOver
		} else if currPos <= lowerBound-100 {
			currPos += 100
			counter++
			goto TurnOver
		}
		if currPos < lowerBound {
			currPos += 100
			if currPos == 0 {
				counter++
				continue
			}
			counter++
		} else if currPos > upperBound {
			currPos -= 100
			if currPos == 0 {
				counter++
				continue
			}
			if currPos == 0 {
				counter++
			}
		}
	}
	fmt.Println("Here is the count", counter)
}

func readFileGetSeq() []dialMotion {
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("Failed to read input.txt")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var seq []dialMotion
	for scanner.Scan() {
		text := scanner.Text()
		suffix, found := strings.CutPrefix(text, "L")
		if found {
			num, err := strconv.Atoi(suffix)
			if err != nil {
				fmt.Println("Failed to parse input.txt")
				os.Exit(1)
			}
			seq = append(seq, dialMotion{
				direction: "L",
				value:     num,
			})
		} else {
			suffix, _ = strings.CutPrefix(text, "R")
			num, err := strconv.Atoi(suffix)
			if err != nil {
				fmt.Println("Failed to parse input.txt")
				os.Exit(1)
			}
			seq = append(seq, dialMotion{
				direction: "R",
				value:     num,
			})
		}
	}
	return seq
}
