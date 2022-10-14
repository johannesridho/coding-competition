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

		countCuts, err := strconv.Atoi(info[1])
		if err != nil {
			log.Fatal(err)
		}

		curLine++

		firstDeck := input[curLine]
		curLine++
		secondDeck := input[curLine]
		curLine++

		lineResult := solve(countCuts, firstDeck, secondDeck)
		lineOutput := fmt.Sprintf("Case #%d: %s\n", i, lineResult)
		_, err = wr.WriteString(lineOutput)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(lineOutput)
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

func solve(countCuts int, firstDeck, secondDeck string) string {
	//fmt.Println(countCuts, " - ", firstDeck, " - ", secondDeck, "cont", strings.Contains(firstDeck+" "+firstDeck, secondDeck))

	if countCuts == 0 && firstDeck == secondDeck {
		return "YES"
	}

	if len(firstDeck) != len(secondDeck) || countCuts == 0 || !strings.Contains(firstDeck+" "+firstDeck, secondDeck) || (countCuts == 1 && firstDeck == secondDeck) {
		return "NO"
	}

	firstDeckArr := strings.Split(firstDeck, " ")
	//secondDeckArr := strings.Split(secondDeck, " ")

	if len(firstDeckArr) == 2 {
		if firstDeck == secondDeck {
			if countCuts%2 == 0 {
				return "YES"
			}
			return "NO"
		}

		if countCuts%2 == 0 {
			return "NO"
		}

		return "YES"
	}

	return "YES"
}
