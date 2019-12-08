package thirteen

import (
	"fmt"
)

func Thirteen() {
	input := setInput()
	inputParam := [2]int{0, 0}

	params := []int{
		0, 1, 2, 3, 4,
		1, 0, 2, 3, 4,
		2, 0, 1, 3, 4,
		0, 2, 1, 3, 4,
		1, 2, 0, 3, 4,
		2, 1, 0, 3, 4,
		2, 1, 3, 0, 4,
		1, 2, 3, 0, 4,
		3, 2, 1, 0, 4,
		2, 3, 1, 0, 4,
		1, 3, 2, 0, 4,
		3, 1, 2, 0, 4,
		3, 0, 2, 1, 4,
		0, 3, 2, 1, 4,
		2, 3, 0, 1, 4,
		3, 2, 0, 1, 4,
		0, 2, 3, 1, 4,
		2, 0, 3, 1, 4,
		1, 0, 3, 2, 4,
		0, 1, 3, 2, 4,
		3, 1, 0, 2, 4,
		1, 3, 0, 2, 4,
		0, 3, 1, 2, 4,
		3, 0, 1, 2, 4,
		4, 0, 1, 2, 3,
		0, 4, 1, 2, 3,
		1, 4, 0, 2, 3,
		4, 1, 0, 2, 3,
		0, 1, 4, 2, 3,
		1, 0, 4, 2, 3,
		1, 0, 2, 4, 3,
		0, 1, 2, 4, 3,
		2, 1, 0, 4, 3,
		1, 2, 0, 4, 3,
		0, 2, 1, 4, 3,
		2, 0, 1, 4, 3,
		2, 4, 1, 0, 3,
		4, 2, 1, 0, 3,
		1, 2, 4, 0, 3,
		2, 1, 4, 0, 3,
		4, 1, 2, 0, 3,
		1, 4, 2, 0, 3,
		0, 4, 2, 1, 3,
		4, 0, 2, 1, 3,
		2, 0, 4, 1, 3,
		0, 2, 4, 1, 3,
		4, 2, 0, 1, 3,
		2, 4, 0, 1, 3,
		3, 4, 0, 1, 2,
		4, 3, 0, 1, 2,
		0, 3, 4, 1, 2,
		3, 0, 4, 1, 2,
		4, 0, 3, 1, 2,
		0, 4, 3, 1, 2,
		0, 4, 1, 3, 2,
		4, 0, 1, 3, 2,
		1, 0, 4, 3, 2,
		0, 1, 4, 3, 2,
		4, 1, 0, 3, 2,
		1, 4, 0, 3, 2,
		1, 3, 0, 4, 2,
		3, 1, 0, 4, 2,
		0, 1, 3, 4, 2,
		1, 0, 3, 4, 2,
		3, 0, 1, 4, 2,
		0, 3, 1, 4, 2,
		4, 3, 1, 0, 2,
		3, 4, 1, 0, 2,
		1, 4, 3, 0, 2,
		4, 1, 3, 0, 2,
		3, 1, 4, 0, 2,
		1, 3, 4, 0, 2,
		2, 3, 4, 0, 1,
		3, 2, 4, 0, 1,
		4, 2, 3, 0, 1,
		2, 4, 3, 0, 1,
		3, 4, 2, 0, 1,
		4, 3, 2, 0, 1,
		4, 3, 0, 2, 1,
		3, 4, 0, 2, 1,
		0, 4, 3, 2, 1,
		4, 0, 3, 2, 1,
		3, 0, 4, 2, 1,
		0, 3, 4, 2, 1,
		0, 2, 4, 3, 1,
		2, 0, 4, 3, 1,
		4, 0, 2, 3, 1,
		0, 4, 2, 3, 1,
		2, 4, 0, 3, 1,
		4, 2, 0, 3, 1,
		3, 2, 0, 4, 1,
		2, 3, 0, 4, 1,
		0, 3, 2, 4, 1,
		3, 0, 2, 4, 1,
		2, 0, 3, 4, 1,
		0, 2, 3, 4, 1,
		1, 2, 3, 4, 0,
		2, 1, 3, 4, 0,
		3, 1, 2, 4, 0,
		1, 3, 2, 4, 0,
		2, 3, 1, 4, 0,
		3, 2, 1, 4, 0,
		3, 2, 4, 1, 0,
		2, 3, 4, 1, 0,
		4, 3, 2, 1, 0,
		3, 4, 2, 1, 0,
		2, 4, 3, 1, 0,
		4, 2, 3, 1, 0,
		4, 1, 3, 2, 0,
		1, 4, 3, 2, 0,
		3, 4, 1, 2, 0,
		4, 3, 1, 2, 0,
		1, 3, 4, 2, 0,
		3, 1, 4, 2, 0,
		2, 1, 4, 3, 0,
		1, 2, 4, 3, 0,
		4, 2, 1, 3, 0,
		2, 4, 1, 3, 0,
		1, 4, 2, 3, 0,
		4, 1, 2, 3, 0,
	}

	maxAmp := 0

	for i := 0; i < 120; i++ {
		inputParam = [2]int{0, 0}
		for j := 0; j < 5; j++ {
			inputParam[0] = params[i*5+j]
			inputParam[1] = getOutput(input, inputParam)
			if inputParam[1] > maxAmp {
				maxAmp = inputParam[1]
			}

			input = setInput()
		}
	}

	fmt.Printf("Maximal result: %v", maxAmp)

}

func getOutput(input []int, inputParam [2]int) int {
	pointer := 0
	output := 0
	firstInput := true
	for {
		program := input[pointer]

		opCode := program % 10
		paramModeOne := program%1000 - opCode
		paramModeTwo := program%10000 - opCode - paramModeOne

		if program == 99 {
			println("Diagnostic")
			break
		}
		if opCode == 3 {
			if firstInput {
				input[input[pointer+1]] = inputParam[0]
				firstInput = false
			} else {
				input[input[pointer+1]] = inputParam[1]
			}
			pointer += 2
		} else if opCode == 4 {
			if paramModeOne == 0 {
				output = input[input[pointer+1]]
				fmt.Printf("Test: %v \n", input[input[pointer+1]])
			} else {
				output = input[pointer+1]
				fmt.Printf("Test: %v \n", input[pointer+1])
			}

			pointer += 2
		} else if opCode == 5 {
			condition := false
			if paramModeOne == 0 {
				condition = input[input[pointer+1]] != 0
			} else {
				condition = input[pointer+1] != 0
			}
			if condition {
				if paramModeTwo == 0 {
					pointer = input[input[pointer+2]]
				} else {
					pointer = input[pointer+2]
				}
			} else {
				pointer += 3
			}
		} else if opCode == 6 {
			condition := false
			if paramModeOne == 0 {
				condition = input[input[pointer+1]] == 0
			} else {
				condition = input[pointer+1] == 0
			}
			if condition {
				if paramModeTwo == 0 {
					pointer = input[input[pointer+2]]
				} else {
					pointer = input[pointer+2]
				}
			} else {
				pointer += 3
			}
		} else if opCode == 7 {
			var valOne int
			var valTwo int

			if paramModeOne == 0 {
				valOne = input[input[pointer+1]]
			} else {
				valOne = input[pointer+1]
			}
			if paramModeTwo == 0 {
				valTwo = input[input[pointer+2]]
			} else {
				valTwo = input[pointer+2]
			}

			if valOne < valTwo {
				input[input[pointer+3]] = 1
			} else {
				input[input[pointer+3]] = 0
			}

			pointer += 4

		} else if opCode == 8 {

			var valOne int
			var valTwo int

			if paramModeOne == 0 {
				valOne = input[input[pointer+1]]
			} else {
				valOne = input[pointer+1]
			}
			if paramModeTwo == 0 {
				valTwo = input[input[pointer+2]]
			} else {
				valTwo = input[pointer+2]
			}

			if valOne == valTwo {
				input[input[pointer+3]] = 1
			} else {
				input[input[pointer+3]] = 0
			}

			pointer += 4

		} else {

			/*println("Program")
			println(program)
			println(opCode)
			println(paramModeOne)
			println(paramModeTwo)*/

			paramOne := 0
			paramTwo := 0

			if paramModeOne == 0 {
				paramOne = input[input[pointer+1]]
			} else {
				paramOne = input[pointer+1]
			}

			if paramModeTwo == 0 {
				paramTwo = input[input[pointer+2]]
			} else {
				paramTwo = input[pointer+2]
			}

			if opCode == 1 {
				input[input[pointer+3]] = paramOne + paramTwo
			} else if opCode == 2 {
				input[input[pointer+3]] = paramOne * paramTwo
			} else {
				println("Error!")
				break
			}

			pointer += 4
		}
	}

	return output
}

func setInput() []int {
	return []int{
		3, 8,
		1001, 8, 10, 8,
		105, 1,

		0, 0, 21, 30, 55, 76, 97, 114, 195, 276, 357, 438, 99999, 3, 9,

		102, 3, 9, 9,
		4, 9,
		99,
		3, 9,
		1002, 9, 3, 9,
		1001, 9, 5, 9,
		1002, 9, 2, 9,
		1001, 9, 2, 9,
		102, 2, 9, 9,
		4, 9,
		99,
		3, 9,
		1002, 9, 5, 9,
		1001, 9, 2, 9,
		102, 5, 9, 9,
		1001, 9, 4, 9,
		4, 9,
		99,
		3, 9,
		1001, 9, 4, 9,
		102, 5, 9, 9,
		101, 4, 9, 9,
		1002, 9, 4, 9,
		4, 9,
		99,
		3, 9,
		101, 2, 9, 9,
		102, 4, 9, 9,
		1001, 9, 5, 9,
		4, 9,
		99,
		3, 9,
		1002, 9, 2, 9,
		4, 9,
		3, 9,
		102, 2, 9, 9,
		4, 9,
		3, 9,
		102, 2, 9, 9,
		4, 9,
		3, 9,
		1001, 9, 1, 9,
		4, 9,
		3, 9,
		102, 2, 9, 9,
		4, 9,
		3, 9,
		1002, 9, 2, 9,
		4, 9,
		3, 9,
		102, 2, 9, 9,
		4, 9,
		3, 9,
		102, 2, 9, 9,
		4, 9,
		3, 9,
		1001, 9, 2, 9,
		4, 9,
		3, 9,
		101, 1, 9, 9,
		4, 9,
		99,
		3, 9,
		1002, 9, 2, 9,
		4, 9,
		3, 9,
		1001, 9, 2, 9,
		4, 9,
		3, 9,
		101, 1, 9, 9,
		4, 9,
		3, 9,
		1002, 9, 2, 9,
		4, 9,
		3, 9,
		1001, 9, 2, 9,
		4, 9,
		3, 9,
		1001, 9, 2, 9,
		4, 9,
		3, 9,
		1002, 9, 2, 9,
		4, 9,
		3, 9,
		1001, 9, 2, 9,
		4, 9,
		3, 9,
		101, 2, 9, 9,
		4, 9,
		3, 9,
		102, 2, 9, 9,
		4, 9,
		99,
		3, 9,
		101, 1, 9, 9,
		4, 9,
		3, 9,
		1002, 9, 2, 9,
		4, 9,
		3, 9,
		1001, 9, 1, 9,
		4, 9,
		3, 9,
		1002, 9, 2, 9,
		4, 9,
		3, 9,
		1001, 9, 1, 9,
		4, 9,
		3, 9,
		1001, 9, 1, 9,
		4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99}

}
