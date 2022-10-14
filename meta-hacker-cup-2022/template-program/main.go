package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"coding-competition/meta-hacker-cup-2022/template-program/helper"
)

var cache map[[4]int]int

func main() {
	cache = make(map[[4]int]int)

loop:
	for {
		fmt.Println("\nInput command (c: build cache, s: solve problem, x: exit):")
		key := helper.ReadKeyboardInput()
		switch key {
		case 'c':
			start := time.Now()
			fmt.Println("building cache start")
			buildCache()
			fmt.Println("build cache done")
			fmt.Println(time.Since(start))
		case 's':
			start := time.Now()
			fmt.Println("solving problem start")
			solveProblem()
			fmt.Println("solve problem done")
			fmt.Println(time.Since(start))
		case 'x':
			fmt.Println("stopping program")
			break loop
		}
	}
}

func buildCache() {

}

func solveProblem() {
	input := helper.SplitInputStrByLine("meta-hacker-cup-2022/template-program/input.txt")

	numOfTestCases, err := strconv.Atoi(input[0])
	if err != nil {
		log.Fatal(err)
	}

	outputFile, err := os.Create("meta-hacker-cup-2022/template-program/output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	wr := bufio.NewWriter(outputFile)
	defer wr.Flush()

	ch := make(chan helper.Result, numOfTestCases)

	curLine := 1
	for testCase := 1; testCase <= numOfTestCases; testCase++ {
		testCaseInput := input[curLine]
		curLine++

		go helper.StartWork(ch, solveTestCase, testCase, testCaseInput)
	}

	outputStr := helper.WaitForAllResult(ch, numOfTestCases)

	for _, s := range outputStr {
		_, err = wr.WriteString(s)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func solveTestCase(testCaseInput string) string {
	fmt.Println("testCaseInput", testCaseInput)
	return "output"
}
