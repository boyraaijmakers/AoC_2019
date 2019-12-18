package thirtytwo

import (
	"fmt"
	"strconv"
)

func Thirtytwo() {

	input := "59756772370948995765943195844952640015210703313486295362653878290009098923609769261473534009395188480864325959786470084762607666312503091505466258796062230652769633818282653497853018108281567627899722548602257463608530331299936274116326038606007040084159138769832784921878333830514041948066594667152593945159170816779820264758715101494739244533095696039336070510975612190417391067896410262310835830006544632083421447385542256916141256383813360662952845638955872442636455511906111157861890394133454959320174572270568292972621253460895625862616228998147301670850340831993043617316938748361984714845874270986989103792418940945322846146634931990046966552"
	//input = "03036732577212944063491565474664"

	v := make([]int, len(input))
	for i, a := range input {
		v[i] = int(a)
	}

	reps := 10000
	phases := 100
	i, _ := strconv.Atoi(input[0:7])

	fmt.Println(toSignal(extractMessage(v, reps, phases, i)))
}

func toSignal(signal []int) string {
	b := make([]byte, len(signal))
	for i, v := range signal {
		b[i] = byte(v + '0')
	}
	return string(b)
}

func offset(b []byte) int {
	i, _ := strconv.Atoi(string(b[:7]))
	return i
}

func sliceSignal(signal []int, from, to int) []int {
	sliced := make([]int, to-from)
	for i := range sliced {
		sliced[i] = signal[(from+i)%len(signal)]
	}
	return sliced
}

func extractMessage(signal []int, reps, phases, offset int) []int {
	msg := sliceSignal(signal, offset, len(signal)*reps)

	for ; phases > 0; phases-- {
		sum := 0
		for i := len(msg) - 1; i >= 0; i-- {
			sum += msg[i]
			if sum < 0 {
				sum = -sum
			}
			msg[i] = sum % 10
		}
	}

	return msg[:8]
}