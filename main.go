// this package simply takes a sequence of numbers and converts them
// to their ascii equivalent
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func asciiToString(input []string) string {
	output := make([]byte, len(input))
	for idx, c := range input {
		val, err := strconv.Atoi(c)
		if err != nil {
			fmt.Printf("error - %s\n", err)
		}
		output[idx] = byte(val)
	}
	fmt.Printf("output: %v\n", output)
	return string(output)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ASCII conversion tool")

	for {
		// Get input
		fmt.Print("Enter sequence (:quit to exit): ")
		text, err := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if err != nil {
			fmt.Printf("Encountered error: %v\nExiting\n", err)
		}
		if text == ":quit" {
			break
		}

		// Parse input and get ASCII string
		sequence := strings.Fields(text)
		result := asciiToString(sequence)
		fmt.Printf("Result: %s\n", result)
	}
}
