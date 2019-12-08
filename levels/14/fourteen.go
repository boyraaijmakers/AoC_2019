package fourteen

import (
	"fmt"
)

func Fourteen() {
	input := [][]int{setInput(), setInput(), setInput(), setInput(), setInput()}
	inputParam := [10]int{0, 0}
	result := [2]int{0, 0}

	params := []int{
		5, 6, 7, 8, 9,
		6, 5, 7, 8, 9,
		7, 5, 6, 8, 9,
		5, 7, 6, 8, 9,
		6, 7, 5, 8, 9,
		7, 6, 5, 8, 9,
		7, 6, 8, 5, 9,
		6, 7, 8, 5, 9,
		8, 7, 6, 5, 9,
		7, 8, 6, 5, 9,
		6, 8, 7, 5, 9,
		8, 6, 7, 5, 9,
		8, 5, 7, 6, 9,
		5, 8, 7, 6, 9,
		7, 8, 5, 6, 9,
		8, 7, 5, 6, 9,
		5, 7, 8, 6, 9,
		7, 5, 8, 6, 9,
		6, 5, 8, 7, 9,
		5, 6, 8, 7, 9,
		8, 6, 5, 7, 9,
		6, 8, 5, 7, 9,
		5, 8, 6, 7, 9,
		8, 5, 6, 7, 9,
		9, 5, 6, 7, 8,
		5, 9, 6, 7, 8,
		6, 9, 5, 7, 8,
		9, 6, 5, 7, 8,
		5, 6, 9, 7, 8,
		6, 5, 9, 7, 8,
		6, 5, 7, 9, 8,
		5, 6, 7, 9, 8,
		7, 6, 5, 9, 8,
		6, 7, 5, 9, 8,
		5, 7, 6, 9, 8,
		7, 5, 6, 9, 8,
		7, 9, 6, 5, 8,
		9, 7, 6, 5, 8,
		6, 7, 9, 5, 8,
		7, 6, 9, 5, 8,
		9, 6, 7, 5, 8,
		6, 9, 7, 5, 8,
		5, 9, 7, 6, 8,
		9, 5, 7, 6, 8,
		7, 5, 9, 6, 8,
		5, 7, 9, 6, 8,
		9, 7, 5, 6, 8,
		7, 9, 5, 6, 8,
		8, 9, 5, 6, 7,
		9, 8, 5, 6, 7,
		5, 8, 9, 6, 7,
		8, 5, 9, 6, 7,
		9, 5, 8, 6, 7,
		5, 9, 8, 6, 7,
		5, 9, 6, 8, 7,
		9, 5, 6, 8, 7,
		6, 5, 9, 8, 7,
		5, 6, 9, 8, 7,
		9, 6, 5, 8, 7,
		6, 9, 5, 8, 7,
		6, 8, 5, 9, 7,
		8, 6, 5, 9, 7,
		5, 6, 8, 9, 7,
		6, 5, 8, 9, 7,
		8, 5, 6, 9, 7,
		5, 8, 6, 9, 7,
		9, 8, 6, 5, 7,
		8, 9, 6, 5, 7,
		6, 9, 8, 5, 7,
		9, 6, 8, 5, 7,
		8, 6, 9, 5, 7,
		6, 8, 9, 5, 7,
		7, 8, 9, 5, 6,
		8, 7, 9, 5, 6,
		9, 7, 8, 5, 6,
		7, 9, 8, 5, 6,
		8, 9, 7, 5, 6,
		9, 8, 7, 5, 6,
		9, 8, 5, 7, 6,
		8, 9, 5, 7, 6,
		5, 9, 8, 7, 6,
		9, 5, 8, 7, 6,
		8, 5, 9, 7, 6,
		5, 8, 9, 7, 6,
		5, 7, 9, 8, 6,
		7, 5, 9, 8, 6,
		9, 5, 7, 8, 6,
		5, 9, 7, 8, 6,
		7, 9, 5, 8, 6,
		9, 7, 5, 8, 6,
		8, 7, 5, 9, 6,
		7, 8, 5, 9, 6,
		5, 8, 7, 9, 6,
		8, 5, 7, 9, 6,
		7, 5, 8, 9, 6,
		5, 7, 8, 9, 6,
		6, 7, 8, 9, 5,
		7, 6, 8, 9, 5,
		8, 6, 7, 9, 5,
		6, 8, 7, 9, 5,
		7, 8, 6, 9, 5,
		8, 7, 6, 9, 5,
		8, 7, 9, 6, 5,
		7, 8, 9, 6, 5,
		9, 8, 7, 6, 5,
		8, 9, 7, 6, 5,
		7, 9, 8, 6, 5,
		9, 7, 8, 6, 5,
		9, 6, 8, 7, 5,
		6, 9, 8, 7, 5,
		8, 9, 6, 7, 5,
		9, 8, 6, 7, 5,
		6, 8, 9, 7, 5,
		8, 6, 9, 7, 5,
		7, 6, 9, 8, 5,
		6, 7, 9, 8, 5,
		9, 7, 6, 8, 5,
		7, 9, 6, 8, 5,
		6, 9, 7, 8, 5,
		9, 6, 7, 8, 5,
	}

	maxAmp := 0

	for i := 0; i < 120; i++ {
		inputParam = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		pointers := [5]int{0, 0, 0, 0, 0}
		for {
			for j := 0; j < 5; j++ {
				inputParam[j*2] = params[i*5+j]
				result = getOutput(input[j], [2]int{inputParam[2*j], inputParam[2*j+1]}, pointers[j])
				if j == 4 {
					inputParam[1] = result[0]
				} else {
					inputParam[j*2+3] = result[0]
				}
				pointers[j] = result[1]
				if inputParam[j*2+1] > maxAmp {
					maxAmp = inputParam[j*2+1]
				}
			}
			if result[1] == -1 {
				break
			}
		}
		input = [][]int{setInput(), setInput(), setInput(), setInput(), setInput()}
	}

	fmt.Printf("Maximal result: %v", maxAmp)
}

func getOutput(input []int, inputParam [2]int, pointer int) [2]int {
	firstInput := false
	if pointer == 0 {
		firstInput = true
	}

	output := 0
	for {
		program := input[pointer]

		opCode := program % 10
		paramModeOne := program%1000 - opCode
		paramModeTwo := program%10000 - opCode - paramModeOne

		if program == 99 {
			//println("Diagnostic")
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
				//fmt.Printf("Test: %v \n", input[input[pointer+1]])
			} else {
				output = input[pointer+1]
				//fmt.Printf("Test: %v \n", input[pointer+1])
			}

			pointer += 2

			return [2]int{output, pointer}
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

	return [2]int{output, -1}
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
