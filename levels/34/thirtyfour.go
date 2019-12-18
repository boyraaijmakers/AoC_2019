package thirtyfour

import "fmt"

func Thirtyfour() {


	var inputParams [][]rune
	var inputParamsAscii [][]int64

	inputParams = append(inputParams, []rune{'A','A','B' ,'C','A','C','B','C','A','B'})
	inputParams = append(inputParams, []rune{'L','4','L','0','L','6'})
	inputParams = append(inputParams, []rune{'L','6','L','4','R','8','R','8'})
	inputParams = append(inputParams, []rune{'L','6','R','8','L','0','L','8','L','8'})

	for _, input := range inputParams {
		var asciiRow []int64
		for _, el := range input {
			if el == '0' {
				asciiRow = append(asciiRow, 49)
				asciiRow = append(asciiRow, 48)
			} else {
				asciiRow = append(asciiRow, int64(el))
			}
			asciiRow = append(asciiRow,44)
		}
		asciiRow[len(asciiRow) - 1] = 10
		inputParamsAscii = append(inputParamsAscii, asciiRow)
	}

	inputParamsAscii = append(inputParamsAscii, []int64{110, 10})

	fmt.Printf("Messed up the code so no output anymore :(")

}

