package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// note: output.txt not submitted, this program might be incorrect
func main() {
	start := time.Now()

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
		totalTrees, err := strconv.Atoi(input[curLine])
		if err != nil {
			log.Fatal(err)
		}

		curLine++

		trees := make([][]int, totalTrees)
		for row := range trees {
			treeCoord := strings.Split(input[curLine], " ")
			x, _ := strconv.Atoi(treeCoord[0])
			y, _ := strconv.Atoi(treeCoord[1])
			trees[row] = []int{x, y}
			curLine++
		}

		totalWells, err := strconv.Atoi(input[curLine])
		if err != nil {
			log.Fatal(err)
		}

		curLine++

		wells := make([][]int, totalWells)
		for row := range wells {
			wellCoord := strings.Split(input[curLine], " ")
			x, _ := strconv.Atoi(wellCoord[0])
			y, _ := strconv.Atoi(wellCoord[1])
			wells[row] = []int{x, y}
			curLine++
		}

		result := solve(trees, wells)
		lineOutput := fmt.Sprintf("Case #%d: %v\n", i, result)
		_, err = wr.WriteString(lineOutput)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(lineOutput)
		//fmt.Println(time.Since(start))
	}

	fmt.Println(time.Since(start))
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

func solve(trees, wells [][]int) int {
	//fmt.Println("l", len(trees), len(wells))
	var sum int

	xTreeCount := make(map[int]int)
	yTreeCount := make(map[int]int)
	xWellCount := make(map[int]int)
	yWellCount := make(map[int]int)

	for _, tree := range trees {
		xTreeCount[tree[0]]++
		yTreeCount[tree[1]]++
	}

	for _, well := range wells {
		xWellCount[well[0]]++
		yWellCount[well[1]]++
	}

	//fmt.Println(xTreeCount)
	//fmt.Println(yTreeCount)
	//fmt.Println(xWellCount)
	//fmt.Println(yWellCount)

	for xWell, wellCount := range xWellCount {
		for xTree, treeCount := range xTreeCount {
			sum += ((xWell - xTree) * (xWell - xTree) * wellCount * treeCount) % 1000000007
		}
	}

	for yWell, wellCount := range yWellCount {
		for yTree, treeCount := range yTreeCount {
			sum += ((yWell - yTree) * (yWell - yTree) * wellCount * treeCount) % 1000000007
		}
	}

	return sum % 1000000007
}
