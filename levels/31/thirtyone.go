package thirtyone

import (
	"strconv"
)

func Thirtyone() {

	input := "59756772370948995765943195844952640015210703313486295362653878290009098923609769261473534009395188480864325959786470084762607666312503091505466258796062230652769633818282653497853018108281567627899722548602257463608530331299936274116326038606007040084159138769832784921878333830514041948066594667152593945159170816779820264758715101494739244533095696039336070510975612190417391067896410262310835830006544632083421447385542256916141256383813360662952845638955872442636455511906111157861890394133454959320174572270568292972621253460895625862616228998147301670850340831993043617316938748361984714845874270986989103792418940945322846146634931990046966552"

	for i:=0;i<100;i++ {
		input = getNextPhase(input)
	}

	println(input[0:8])
}

func getNextPhase(input string) string {
	var nextPhase []int
	for i:=1;i<=len(input);i++ {
		rowSum := 0
		for index, multi := range getPattern(i, len(input)) {
			val, _ := strconv.Atoi(string(input[index]))
			rowSum += val * multi
		}
		if rowSum < 0 {
			rowSum = -rowSum
		}
		nextPhase = append(nextPhase, rowSum % 10)
	}

	result := ""
	for _, el := range nextPhase {
		result += strconv.Itoa(el)
	}
	return result
}

func getPattern(row int, length int) []int {
	var output []int
	var currentCount int = 1
	options := []int{0, 1, 0, -1}
	currentOption := 0
	pointer := 0
	for {
		for currentCount < row {
			output = append(output, options[currentOption])
			currentCount++
			pointer++
			if pointer >= length {
				return output
			}
		}
		currentCount = 0
		currentOption = (currentOption + 1) % 4
	}
	return output
}