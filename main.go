package main

import "fmt"

type Pattern struct {
	input  [3]int
	output int
}

var seed = [60]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}

func main() {
	patterns := InitPatterns()
	counter := 0
	computeNextGeneration(seed, patterns, counter)
	fmt.Println("Done")
}

func computeNextGeneration(input [60]int, patterns []Pattern, counter int) {

	counter++
	if counter < 100 {
		if counter > 1 {
			fmt.Printf("%v\n", makePrintableStringFromInput(input))
		}

		output := [60]int{}
		for i := 0; i < len(input)-2; i++ {
			slice := [3]int{}
			slice[0] = input[i]
			slice[1] = input[i+1]
			slice[2] = input[i+2]
			for _, pattern := range patterns {
				if pattern.input == slice {
					output[i+1] = pattern.output
				}
			}
		}
		computeNextGeneration(output, patterns, counter)
	}

}

func makePrintableStringFromInput(input [60]int) string {
	outputString := ""
	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			outputString += " "
		}
		if input[i] == 1 {
			outputString += "*"
		}
	}

	return outputString
}
