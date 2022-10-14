package helper

import (
	"fmt"
)

type Result struct {
	TestCaseNumber int
	Val            string
}

func StartWork(ch chan Result, fn func(string) string, testCaseNumber int, testCaseInput string) {
	fmt.Println("starting test", testCaseNumber)
	result := fn(testCaseInput)
	fmt.Println("finished test", testCaseNumber)
	ch <- Result{testCaseNumber, fmt.Sprintf("Case #%d: %v\n", testCaseNumber, result)}
}

func WaitForAllResult(ch chan Result, totalTestCases int) []string {
	var count int
	outputStr := make([]string, totalTestCases)

loop:
	for {
		select {
		case result := <-ch:
			count++
			outputStr[result.TestCaseNumber-1] = result.Val

			if count == totalTestCases {
				fmt.Println("all done", count)
				break loop
			}
		}
	}

	return outputStr
}
