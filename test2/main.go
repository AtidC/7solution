package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Decode(encoded string) string {
	n := len(encoded) + 1
	result := make([]int, n)
	fmt.Println("input :", encoded)

	// Assign values based on rules
	for i := 0; i < len(encoded); i++ {
		switch encoded[i] {
		case 'L':
			if result[i] <= result[i+1] {
				result[i] = result[i+1] + 1
			}
		case 'R':
			if result[i+1] <= result[i] {
				result[i+1] = result[i] + 1
			}
		case '=':
			result[i+1] = result[i]
		}
		//fmt.Println("decode [%s] : %v", encoded, result)
	}

	// Validate and fix result
	for loop := 0; loop < 2; loop++ {
		for i := 0; i < len(encoded); i++ {
			switch encoded[i] {
			case 'L':
				if result[i] < result[i+1] {
					result[i] = result[i] + 2
				} else if result[i] == result[i+1] {
					result[i] = result[i] + 1
				}
			case 'R':
				if result[i] > result[i+1] {
					result[i+1] = result[i+1] + 2
				} else if result[i] == result[i+1] {
					result[i+1] = result[i+1] + 1
				}
			case '=':
				if result[i] < result[i+1] {
					result[i] = result[i+1]
				} else if result[i] > result[i+1] {
					result[i+1] = result[i]
				}
			}
			//fmt.Println("decode [%s] : %v", encoded, result)
		}
	}

	// Convert result to string
	var sequence strings.Builder
	for _, num := range result {
		sequence.WriteString(strconv.Itoa(num))
	}

	fmt.Println("output :", sequence.String())
	return sequence.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter encoded string: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	result := Decode(input)
	fmt.Println("Decoded result:", result)
}
