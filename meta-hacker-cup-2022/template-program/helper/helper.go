package helper

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func SplitInputStrByLine(fileName string) []string {
	file, err := os.Open(fileName)
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

func ReadKeyboardInput() rune {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	return char
}

func PowerMod(base, exponent, modulus int) int {
	if base < 1 || exponent < 0 || modulus < 1 {
		return -1
	}

	result := 1
	for exponent > 0 {
		if (exponent % 2) == 1 {
			result = (result * base) % modulus
		}
		base = (base * base) % modulus
		exponent = exponent / 2
	}
	return result
}

func MinMod(a, b, m int) int {
	return ((a % m) - (b % m) + m) % m
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Square(a int) int {
	return a * a
}
