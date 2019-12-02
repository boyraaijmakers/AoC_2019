package three

import "fmt"

func Three() {
	input := []int{
		1, 0, 0, 3,
		1, 1, 2, 3,
		1, 3, 4, 3,
		1, 5, 0, 3,
		2, 1, 10, 19,
		1, 19, 5, 23,
		2, 23, 6, 27,
		1, 27, 5, 31,
		2, 6, 31, 35,
		1, 5, 35, 39,
		2, 39, 9, 43,
		1, 43, 5, 47,
		1, 10, 47, 51,
		1, 51, 6, 55,
		1, 55, 10, 59,
		1, 59, 6, 63,
		2, 13, 63, 67,
		1, 9, 67, 71,
		2, 6, 71, 75,
		1, 5, 75, 79,
		1, 9, 79, 83,
		2, 6, 83, 87,
		1, 5, 87, 91,
		2, 6, 91, 95,
		2, 95, 9, 99,
		1, 99, 6, 103,
		1, 103, 13, 107,
		2, 13, 107, 111,
		2, 111, 10, 115,
		1, 115, 6, 119,
		1, 6, 119, 123,
		2, 6, 123, 127,
		1, 127, 5, 131,
		2, 131, 6, 135,
		1, 135, 2, 139,
		1, 139, 9, 0,
		99, 2, 14, 0, 0}

	input[1] = 12
	input[2] = 2

	fmt.Printf("Answer: %v", getOutput(input))

}

func getOutput(input []int) int {
	pointer := 0
	for {
		program := input[pointer]
		first := input[pointer+1]
		second := input[pointer+2]
		third := input[pointer+3]
		var newVal int = 0

		if first >= len(input) || second >= len(input) || third >= len(input) {
			return 0
		}

		if program == 99 {
			break
		} else if program == 1 {
			newVal = input[first] + input[second]
		} else if program == 2 {
			newVal = input[first] * input[second]
		} else {
			return 0
		}
		input[third] = newVal
		pointer += 4
	}
	return input[0]
}
