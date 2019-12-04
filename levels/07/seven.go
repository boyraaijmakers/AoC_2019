package seven

import (
	"strconv"
)

func Seven() {
	inputLower := 271973
	inputUpper := 785961
	counter := 0

	for i := inputLower; i <= inputUpper; i++ {
		if checkPassword(i) {
			counter++
		}
	}
	print(counter)

}

func checkPassword(password int) bool {
	passwordSplit := strconv.Itoa(password)

	incrementCond := true
	similarityCond := false

	for i := 1; i < len(passwordSplit); i++ {
		if passwordSplit[i] < passwordSplit[i-1] {
			incrementCond = false
			break
		}
		if passwordSplit[i] == passwordSplit[i-1] {
			similarityCond = true
		}
	}
	return incrementCond && similarityCond
}
