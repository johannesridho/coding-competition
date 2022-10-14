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
		totalRow, err := strconv.Atoi(info[0])
		if err != nil {
			log.Fatal(err)
		}

		totalCol, err := strconv.Atoi(info[1])
		if err != nil {
			log.Fatal(err)
		}

		curLine++

		matrix := make([][]rune, totalRow)
		for row := range matrix {
			matrix[row] = make([]rune, totalCol)
			for idx, c := range input[curLine] {
				matrix[row][idx] = c
			}
			curLine++
		}

		result, resultMatrix := solve(matrix)
		lineOutput := fmt.Sprintf("Case #%d: %s\n", i, result)
		_, err = wr.WriteString(lineOutput)
		if err != nil {
			log.Fatal(err)
		}

		for row := range resultMatrix {
			_, err = wr.WriteString(string(resultMatrix[row]) + "\n")
			if err != nil {
				log.Fatal(err)
			}
		}
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

func solve(matrix [][]rune) (string, [][]rune) {
	var treeFound bool
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == '^' {
				treeFound = true
				break
			}
		}
		if treeFound {
			break
		}
	}

	if !treeFound {
		return "Possible", matrix
	}

	if len(matrix) <= 1 || len(matrix[0]) <= 1 {
		return "Impossible", nil
	}

	for row := range matrix {
		for col := range matrix[row] {
			matrix[row][col] = '^'
		}
	}

	return "Possible", matrix
}
