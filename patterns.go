package main

func InitPatterns() []Pattern {
	patterns := []Pattern{}

	patterns = append(patterns, generatePattern([3]int{1, 1, 1}, 0))
	patterns = append(patterns, generatePattern([3]int{1, 1, 0}, 1))
	patterns = append(patterns, generatePattern([3]int{1, 0, 1}, 1))
	patterns = append(patterns, generatePattern([3]int{1, 0, 0}, 0))
	patterns = append(patterns, generatePattern([3]int{0, 1, 1}, 1))
	patterns = append(patterns, generatePattern([3]int{0, 1, 0}, 1))
	patterns = append(patterns, generatePattern([3]int{0, 0, 1}, 1))
	patterns = append(patterns, generatePattern([3]int{0, 0, 0}, 0))

	return patterns
}

func generatePattern(input [3]int, output int) Pattern {
	p := Pattern{}
	p.input = input
	p.output = output
	return p
}
