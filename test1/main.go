package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func maxPathSum(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += max(triangle[i+1][j], triangle[i+1][j+1])
		}
	}

	return triangle[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func readJSONFromURL(url string) ([][]int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data [][]int
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func main() {
	// Test case 1: Input as a static triangle
	triangle := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}
	fmt.Println("Output for test case 1:", maxPathSum(triangle)) // Expected: 237

	// Test case 2: Input from a JSON URL
	url := "https://raw.githubusercontent.com/7-solutions/backend-challenge/main/files/hard.json"
	triangleFromFile, err := readJSONFromURL(url)
	if err != nil {
		fmt.Println("Error reading JSON from URL:", err)
		return
	}
	fmt.Println("Output for test case 2:", maxPathSum(triangleFromFile)) // Expected: 7273
}
