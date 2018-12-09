package main

import (
	"bufio"
	"fmt"
	"os"
)

func charDifference(x string, y string) int {
	count := 0
	for i, _ := range x {
		if x[i] != y[i] {
			count = count + 1
		}
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	strs := []string{}
	for scanner.Scan() {
		text := string(scanner.Text())
		strs = append(strs, text)
	}

	for _, text1 := range strs {
		for _, text2 := range strs {
			difference := charDifference(text1, text2)
			if difference == 1 {
				for i, _ := range text1 {
					if text1[i] == text2[i] {
						fmt.Print(string(text1[i]))
					}
				}
				fmt.Println()
			}
		}
	}
}
