package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
)

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func readDataset(filename string) (data []byte, err error) {
	data, err = os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func part1(data []byte) (sum int, err error) {
	lines := bytes.Split(data, []byte("\n"))

	for _, l := range lines {
		var firstDig byte
		for i := 0; i < len(l); i++ {
			if isDigit(l[i]) {
				firstDig = l[i]
				break
			}
		}

		var lastDig byte
		for i := len(l) - 1; i >= 0; i-- {
			if isDigit(l[i]) {
				lastDig = l[i]
				break
			}
		}

		number, err := strconv.Atoi(string(firstDig) + string(lastDig))
		if err != nil {
			return 0, err
		}
		sum += number
	}

	return sum, nil
}

func part2(data []byte) (sum int, err error) {
	// create array of string with digits
	line := ""
	lines := make([]string, 0)
	for i := 0; i < len(data); i++ {
		// if newline or end - append and continue
		if data[i] == '\n' || i == len(data)-1 {
			lines = append(lines, line)
			line = ""
			continue
		}

		// add digit to line
		if isDigit(data[i]) {
			line += string(data[i])
		} else if (i+2 <= len(data)-1) && data[i] == 'o' && data[i+1] == 'n' && data[i+2] == 'e' {
			line += "1"
			i += 2
		} else if (i+2 <= len(data)-1) && data[i] == 't' && data[i+1] == 'w' && data[i+2] == 'o' {
			line += "2"
			i += 2
		} else if (i+5 <= len(data)-1) && data[i] == 't' && data[i+1] == 'h' && data[i+2] == 'r' && data[i+3] == 'e' && data[i+4] == 'e' {
			line += "3"
			i += 4
		} else if (i+3 <= len(data)-1) && data[i] == 'f' && data[i+1] == 'o' && data[i+2] == 'u' && data[i+3] == 'r' {
			line += "4"
			i += 3
		} else if (i+3 <= len(data)-1) && data[i] == 'f' && data[i+1] == 'i' && data[i+2] == 'v' && data[i+3] == 'e' {
			line += "5"
			i += 3
		} else if (i+3 <= len(data)-1) && data[i] == 's' && data[i+1] == 'i' && data[i+2] == 'x' {
			line += "6"
			i += 2
		} else if (i+3 <= len(data)-1) && data[i] == 's' && data[i+1] == 'e' && data[i+2] == 'v' && data[i+3] == 'e' && data[i+4] == 'n' {
			line += "7"
			i += 4
		} else if (i+4 <= len(data)-1) && data[i] == 'e' && data[i+1] == 'i' && data[i+2] == 'g' && data[i+3] == 'h' && data[i+4] == 't' {
			line += "8"
			i += 4
		} else if (i+4 <= len(data)-1) && data[i] == 'n' && data[i+1] == 'i' && data[i+2] == 'n' && data[i+3] == 'e' {
			line += "9"
			i += 3
		}
	}

	// iterate over array of digits, take first and last digit , convert to number and add to the summ
	sum = 0
	for _, l := range lines {
		digit, err := strconv.Atoi(string(l[0]) + string(l[len(l)-1]))
		if err != nil {
			return 0, err
		}
		sum += digit
	}

	return sum, nil
}

func main() {
	data, err := readDataset("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result1, err := part1(data)
	if err != nil {
		log.Fatal(err)
	}

	result2, err := part2(data)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Part1 - Sum: %d", result1)
	log.Printf("Part1 - Sum: %d", result2)
}
