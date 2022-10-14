package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := processInput()

	numOfTestCases, err := strconv.Atoi(input[0])
	if err != nil {
		log.Fatal(err)
	}

	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	wr := bufio.NewWriter(outputFile)
	defer wr.Flush()

	curLine := 1
	for i := 1; i <= numOfTestCases; i++ {
		info := strings.Split(input[curLine], " ")
		capacity, err := strconv.Atoi(info[1])
		if err != nil {
			log.Fatal(err)
		}

		styles := strings.Split(input[curLine+1], " ")
		result := isPossibleToArrange(styles, capacity)
		lineOutput := fmt.Sprintf("Case #%d: %s\n", i, result)
		_, err = wr.WriteString(lineOutput)
		if err != nil {
			log.Fatal(err)
		}

		curLine += 2
	}
}

func processInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}

func isPossibleToArrange(styles []string, capacity int) string {
	if capacity*2 < len(styles) {
		return "NO"
	}

	counts := make(map[string]int)
	for _, style := range styles {
		counts[style]++
		if counts[style] > 2 {
			return "NO"
		}
	}

	return "YES"
}
