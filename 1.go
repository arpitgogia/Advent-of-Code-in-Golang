package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	answer := 0
	for scanner.Scan() {
	    text := string(scanner.Text())
	    converted, _ := strconv.Atoi(text[1:])
	    if text[0] == '+' {
	    	answer = answer + converted
	    } else {
	    	answer = answer - converted
	    }
	}
	fmt.Println(answer)
}
