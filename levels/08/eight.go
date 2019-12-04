package eight

import (
	"strconv"
)

func Eight() {
	inputLower := 271973
	inputUpper := 785961
	counter := 0

	for i := inputLower; i <= inputUpper; i++ {
		if checkPassword(i) {
			if checkStrictPair(i) {
				counter++
			}
		}
	}
	println(counter)

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

func checkStrictPair(password int) bool {
	passwordSplit := strconv.Itoa(password)
	count := 1
	outcome := false
	current := passwordSplit[0]

	for i := 1; i < len(passwordSplit); i++ {
		if passwordSplit[i] != current {
			if count == 2 {
				outcome = true
			}
			current = passwordSplit[i]
			count = 1
		} else {
			count++
		}
	}

	if count == 2 {
		outcome = true
	}
	return outcome
}
