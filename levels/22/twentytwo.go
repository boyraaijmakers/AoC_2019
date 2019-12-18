package twentytwo

import "fmt"

func Twentytwo() {

	input := []int64{
		3, 8,
		1005, 8, 352,
		1106, 0, 11,
		0, 0, 0,
		104, 1,
		104, 0,
		3, 8,
		102, -1, 8, 10,
		101, 1, 10, 10,
		4, 10,
		108, 1, 8, 10,
		4, 10,
		102, 1, 8, 28,
		1,
		1003, 20, 10, 2,
		106, 11, 10, 2,
		1107, 1, 10, 1,
		1001, 14, 10, 3, 8,
		1002, 8, -1, 10,
		1001, 10, 1, 10,
		4, 10,
		1008, 8, 0, 10,
		4, 10,
		1002, 8, 1, 67, 2,
		1009, 7, 10, 3, 8,
		1002, 8, -1, 10,
		1001, 10, 1, 10,
		4, 10,
		108, 0, 8, 10,
		4, 10,
		101, 0, 8, 92, 1,
		105, 9, 10,
		1006, 0, 89, 1,
		108, 9, 10,
		3, 8,
		1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1002, 8, 1, 126, 1, 1101, 14, 10, 1, 1005, 3, 10, 1006, 0, 29, 1006, 0, 91, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 108, 1, 8, 10, 4, 10, 1002, 8, 1, 161, 1, 1, 6, 10, 1006, 0, 65, 2, 106, 13, 10, 1006, 0, 36, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 102, 1, 8, 198, 1, 105, 15, 10, 1, 1004, 0, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 101, 0, 8, 228, 2, 1006, 8, 10, 2, 1001, 16, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 1001, 8, 0, 257, 1006, 0, 19, 2, 6, 10, 10, 2, 4, 13, 10, 2, 1002, 4, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1002, 8, 1, 295, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 108, 0, 8, 10, 4, 10, 101, 0, 8, 316, 2, 101, 6, 10, 1006, 0, 84, 2, 1004, 13, 10, 1, 1109, 3, 10, 101, 1, 9, 9, 1007, 9, 1046, 10, 1005, 10, 15, 99, 109, 674, 104, 0, 104, 1, 21101, 387365315340, 0, 1, 21102, 369, 1, 0, 1105, 1, 473, 21101, 666685514536, 0, 1, 21102, 380, 1, 0, 1106, 0, 473, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 21102, 1, 46266346536, 1, 21102, 427, 1, 0, 1105, 1, 473, 21101, 235152829659, 0, 1, 21101, 438, 0, 0, 1105, 1, 473, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 0, 21102, 838337188620, 1, 1, 21101, 461, 0, 0, 1105, 1, 473, 21102, 988753429268, 1, 1, 21102, 1, 472, 0, 1106, 0, 473, 99, 109, 2, 22101, 0, -1, 1, 21101, 40, 0, 2, 21101, 504, 0, 3, 21102, 494, 1, 0, 1106, 0, 537, 109, -2, 2105, 1, 0, 0, 1, 0, 0, 1, 109, 2, 3, 10, 204, -1, 1001, 499, 500, 515, 4, 0, 1001, 499, 1, 499, 108, 4, 499, 10, 1006, 10, 531, 1101, 0, 0, 499, 109, -2, 2106, 0, 0, 0, 109, 4, 2101, 0, -1, 536, 1207, -3, 0, 10, 1006, 10, 554, 21102, 1, 0, -3, 21202, -3, 1, 1, 21201, -2, 0, 2, 21102, 1, 1, 3, 21101, 573, 0, 0, 1105, 1, 578, 109, -4, 2105, 1, 0, 109, 5, 1207, -3, 1, 10, 1006, 10, 601, 2207, -4, -2, 10, 1006, 10, 601, 21201, -4, 0, -4, 1105, 1, 669, 22101, 0, -4, 1, 21201, -3, -1, 2, 21202, -2, 2, 3, 21101, 620, 0, 0, 1106, 0, 578, 22102, 1, 1, -4, 21101, 0, 1, -1, 2207, -4, -2, 10, 1006, 10, 639, 21101, 0, 0, -1, 22202, -2, -1, -2, 2107, 0, -3, 10, 1006, 10, 661, 22101, 0, -1, 1, 21102, 661, 1, 0, 106, 0, 536, 21202, -2, -1, -2, 22201, -4, -2, -4, 109, -5, 2106, 0, 0,
	}

	for i := 0; i < 1000; i++ {
		input = append(input, 0)
	}

	var hullColor [200][200]int64
	var hullPainted [200][200]int
	robotPosition := [2]int16{100, 100}
	robotDirectionInt := 0
	robotDirection := [2]int16{0, 1}
	var pointer int64 = 0
	var relativeBase int64 = 0
	var outputOne int64 = 0
	var outputTwo int64 = 0

	steps := 0

	for {
		println( hullColor[robotPosition[1]][robotPosition[0]])

		outputOne, outputTwo, pointer, relativeBase = getOutput(input, hullColor[robotPosition[1]][robotPosition[0]], pointer, relativeBase)

		fmt.Printf("%v, %v, %v\n", outputOne, outputTwo, pointer)

		if hullPainted[robotPosition[1]][robotPosition[0]] == 0 {
			hullPainted[robotPosition[1]][robotPosition[0]] = 1
		}
		hullColor[robotPosition[1]][robotPosition[0]] = outputOne


		if outputTwo == 0 {
			robotDirectionInt = (robotDirectionInt - 1)
			if robotDirectionInt == -1 {
				robotDirectionInt = 3
			}
		} else if outputTwo == 1 {
			robotDirectionInt = (robotDirectionInt + 1) % 4
		} else if outputTwo == -1 {
			break
		}

		switch robotDirectionInt {
		case 0:
			robotDirection = [2]int16{0, -1}
		case 1:
			robotDirection = [2]int16{1, 0}
		case 2:
			robotDirection = [2]int16{0, 1}
		case 3:
			robotDirection = [2]int16{-1, 0}
		}


		fmt.Printf("%v\n", robotDirection)

		steps++

		if steps > 50 {
			//break
		}

		robotPosition[0] += robotDirection[0]
		robotPosition[1] += robotDirection[1]

		fmt.Printf("%v\n", robotPosition)

		/*for _, row := range hullPainted {
			fmt.Printf("%v\n", row)
		}*/
		//fmt.Printf("%v\n", pointer)
		//fmt.Printf("%v\n", input)
	}

	countPainted := 0
	for _, i := range hullPainted {
		for _, j := range i {
			countPainted += j
		}
	}
	//fmt.Printf("Result: %v", hullPainted)
	fmt.Printf("Result: %v", countPainted)

	for _, row := range hullColor {
		fmt.Printf("%v\n", row)
	}

}

func getOutput(input []int64, inputParam int64, pointer int64, relativeBase int64) (int64, int64, int64, int64) {
	//pointer = 0
	firstOutput := true
	var outputOne int64 = 0
	var outputTwo int64 = 0

	for {
		if inputParam == 1 || inputParam == 0 {
			fmt.Printf("%v", input[pointer:pointer+4])
		}
		program := input[pointer]

		opCode := program % 10
		paramModeOne := program%1000 - opCode
		paramModeTwo := program%10000 - opCode - paramModeOne
		paramModeThree := program - paramModeTwo - paramModeOne - opCode

		if program == 99 {
			println("Diagnostic")
			return -1, -1, -1, -1
		}
		if opCode == 3 {
			if paramModeOne == 0 {
				input[input[pointer+1]] = inputParam
			} else if paramModeOne == 100 {
				input[pointer+1] = inputParam
			} else if paramModeOne == 200 {
				input[relativeBase+input[pointer+1]] = inputParam
			}

			pointer += 2
		} else if opCode == 4 {
			var output int64
			if paramModeOne == 0 {
				//fmt.Printf("Test: %v \n", input[input[pointer+1]])
				output = input[input[pointer+1]]
			} else if paramModeOne == 100 {
				//fmt.Printf("Test: %v \n", input[pointer+1])
				output = input[pointer+1]
			} else if paramModeOne == 200 {
				//fmt.Printf("Test: %v \n", input[relativeBase+input[pointer+1]])
				output = input[relativeBase+input[pointer+1]]
			} else {
				println("Error! 4")
			}

			if firstOutput {
				outputOne = output
				firstOutput = false
			} else {
				outputTwo = output
				return outputOne, outputTwo, pointer + 2, relativeBase
			}

			pointer += 2
		} else if opCode == 5 {
			condition := false
			if paramModeOne == 0 {
				condition = input[input[pointer+1]] != 0
			} else if paramModeOne == 100 {
				condition = input[pointer+1] != 0
			} else if paramModeOne == 200 {
				condition = input[relativeBase+input[pointer+1]] != 0
			}
			if condition {
				if paramModeTwo == 0 {
					pointer = input[input[pointer+2]]
				} else if paramModeTwo == 1000 {
					pointer = input[pointer+2]
				} else if paramModeTwo == 2000 {
					pointer = input[relativeBase+input[pointer+2]]
				}
			} else {
				pointer += 3
			}
		} else if opCode == 6 {
			condition := false
			if paramModeOne == 0 {
				condition = input[input[pointer+1]] == 0
			} else if paramModeOne == 100 {
				condition = input[pointer+1] == 0
			} else if paramModeOne == 200 {
				condition = input[relativeBase+input[pointer+1]] == 0
			}
			if condition {
				if paramModeTwo == 0 {
					pointer = input[input[pointer+2]]
				} else if paramModeTwo == 1000 {
					pointer = input[pointer+2]
				} else if paramModeTwo == 2000 {
					pointer = input[relativeBase+input[pointer+2]]
				}
			} else {
				pointer += 3
			}
		} else if opCode == 7 {
			var valOne int64
			var valTwo int64

			if paramModeOne == 0 {
				valOne = input[input[pointer+1]]
			} else if paramModeOne == 100 {
				valOne = input[pointer+1]
			} else if paramModeOne == 200 {
				valOne = input[relativeBase+input[pointer+1]]
			}
			if paramModeTwo == 0 {
				valTwo = input[input[pointer+2]]
			} else if paramModeTwo == 1000 {
				valTwo = input[pointer+2]
			} else if paramModeTwo == 2000 {
				valTwo = input[relativeBase+input[pointer+2]]
			}

			var opRes int64 = 0
			if valOne < valTwo {
				opRes = 1
			}

			if paramModeThree == 0 {
				input[input[pointer+3]] = opRes
			} else if paramModeThree == 10000 {
				input[pointer+3] = opRes
			} else if paramModeThree == 20000 {
				input[relativeBase+input[pointer+3]] = opRes
			}

			pointer += 4

		} else if opCode == 8 {

			var valOne int64
			var valTwo int64

			if paramModeOne == 0 {
				valOne = input[input[pointer+1]]
			} else if paramModeOne == 100 {
				valOne = input[pointer+1]
			} else if paramModeOne == 200 {
				valOne = input[relativeBase+input[pointer+1]]
			}
			if paramModeTwo == 0 {
				valTwo = input[input[pointer+2]]
			} else if paramModeTwo == 1000 {
				valTwo = input[pointer+2]
			} else if paramModeTwo == 2000 {
				valTwo = input[relativeBase+input[pointer+2]]
			}

			var opRes int64 = 0

			if valOne == valTwo {
				opRes = 1
			}

			if paramModeThree == 0 {
				input[input[pointer+3]] = opRes
			} else if paramModeThree == 10000 {
				input[pointer+3] = opRes
			} else if paramModeThree == 20000 {
				input[relativeBase+input[pointer+3]] = opRes
			}

			pointer += 4

		} else if opCode == 9 {
			if paramModeOne == 0 {
				relativeBase += input[input[pointer+1]]
			} else if paramModeOne == 100 {
				relativeBase += input[pointer+1]
			} else if paramModeOne == 200 {
				relativeBase += input[relativeBase+input[pointer+1]]
			}
			pointer += 2
		} else {

			var paramOne int64 = 0
			var paramTwo int64 = 0

			if paramModeOne == 0 {
				paramOne = input[input[pointer+1]]
			} else if paramModeOne == 100 {
				paramOne = input[pointer+1]
			} else if paramModeOne == 200 {
				paramOne = input[relativeBase+input[pointer+1]]
			}

			if paramModeTwo == 0 {
				paramTwo = input[input[pointer+2]]
			} else if paramModeTwo == 1000 {
				paramTwo = input[pointer+2]
			} else if paramModeTwo == 2000 {
				paramTwo = input[relativeBase+input[pointer+2]]
			}

			var opRes int64 = 0

			if opCode == 1 {
				opRes = paramOne + paramTwo
			} else if opCode == 2 {
				opRes = paramOne * paramTwo
			} else {
				println("Error! 1/2")
				break
			}

			if paramModeThree == 0 {
				input[input[pointer+3]] = opRes
			} else if paramModeThree == 10000 {
				input[pointer+3] = opRes
			} else if paramModeThree == 20000 {
				input[relativeBase+input[pointer+3]] = opRes
			}

			pointer += 4
		}
	}
	return -1, -1, -1, -1
}