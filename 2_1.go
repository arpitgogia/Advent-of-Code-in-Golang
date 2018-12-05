package main

import (
    "fmt";
    "os";
    "bufio";
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
	twoCount := 0
    threeCount := 0
    for scanner.Scan() {
        text := string(scanner.Text())
        charMap := make(map[string]int)
        for _, ch := range text {
            if charMap[string(ch)] == 0 {
                charMap[string(ch)] = strings.Count(text, string(ch))
            }
        }
        didCheckTwo := false
        didCheckThree := false
        for _, value := range charMap {
        	if value == 2 && didCheckTwo == false {
                twoCount = twoCount + 1
                didCheckTwo = true
            } else if value == 3 && didCheckThree == false {
                threeCount = threeCount + 1
                didCheckThree = true
            }
        }
    }
    fmt.Println(twoCount * threeCount)
}
