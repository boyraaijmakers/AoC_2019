package twentyseven

import (
	"fmt"
	"strconv"
	"strings"
)

func Twentyseven() {

	input := []string{
		"1 FJFL, 1 BPVQN => 7 CMNH",
		"6 FJFL, 2 KZJLT, 3 DZQJ => 2 NSPZ",
		"11 TPZDN => 2 TNMC",
		"1 NSPZ, 2 KQVL => 2 HPNWP",
		"3 XHDVT => 9 LRBN",
		"3 LRBN => 6 TPZDN",
		"1 KPFLZ, 1 XVXCZ => 6 WHMLV",
		"1 BDWQP, 1 JPGK, 1 MTWG => 5 GLHWQ",
		"2 BGLTP, 1 HPNWP, 2 GLHWQ, 9 CRJZ, 22 QVQJ, 3 PHGWC, 1 BDWQP => 3 LKPNB",
		"4 BDSB => 2 PNSD",
		"2 BRJDJ, 13 THQR => 2 BGLTP",
		"1 WHJKH, 2 JBTJ => 6 THQR",
		"1 JBTJ => 9 WGVP",
		"10 CTXHZ, 2 DGMN => 5 TNVC",
		"7 LCSV, 1 LKPNB, 36 CMNH, 1 JZXPH, 20 DGJPN, 3 WDWB, 69 DXJKC, 3 WHJKH, 18 XSGP, 22 CGZL, 2 BNVB, 57 PNSD => 1 FUEL",
		"13 CRCG, 8 NMQN => 1 JZXPH",
		"2 FZVS, 2 ZPFBH => 9 SRPD",
		"1 QPNTQ, 4 QVQJ, 1 XZKTG => 9 WDWB",
		"6 SXZW => 5 FJFL",
		"6 GVGZ => 6 ZPFBH",
		"1 JPGK, 3 WDFXH, 22 FJFL => 7 BDSB",
		"3 WHMLV => 4 JPGK",
		"7 CGZL, 4 LRBN => 8 MTWG",
		"11 SXZW, 33 ZTBFN => 4 XVXCZ",
		"1 FZVS, 1 TNMC, 7 JPGK => 9 FLHW",
		"2 XKFZ => 8 CGZL",
		"5 WHMLV => 8 MQRS",
		"1 QVSH, 6 TPZDN, 9 JQHCH => 2 BMNJ",
		"3 CMNH, 10 XKFZ => 2 KQVL",
		"119 ORE => 9 PSPQ",
		"1 WGVP, 18 BRJDJ => 9 PHGWC",
		"110 ORE => 6 NMQN",
		"13 NMQN, 24 XVXCZ, 9 XHDVT => 6 KQVS",
		"6 TNMC => 4 DXJKC",
		"1 XZKTG => 8 WHJKH",
		"1 KPFLZ, 1 LRBN, 7 KQVS => 9 JBTJ",
		"1 XKFZ => 8 JVGD",
		"152 ORE => 7 SXZW",
		"1 BDWQP => 5 CTXHZ",
		"2 JVGD, 8 DGMN, 2 MTWG => 6 QVQJ",
		"1 KQVL => 2 BNVB",
		"3 DZQJ, 37 MQRS => 4 CRJZ",
		"1 KQVL, 5 WDFXH => 7 BDWQP",
		"3 GVGZ => 3 BPVQN",
		"4 PSPQ, 6 ZTBFN => 1 KPFLZ",
		"34 FBTG => 9 XZKTG",
		"14 TNMC, 4 FZVS, 3 MTWG => 9 KZJLT",
		"157 ORE => 6 GVGZ",
		"5 JVGD, 11 JPGK => 5 CRCG",
		"1 SXZW, 1 NMQN => 3 XHDVT",
		"1 FBTG => 5 JQHCH",
		"3 WDFXH, 4 FZVS, 9 CGFML => 6 DZQJ",
		"5 BDWQP, 3 TNVC, 7 SRPD, 1 WDFXH, 3 JQHCH, 4 QVQJ, 5 CRCG, 4 DGMN => 7 XSGP",
		"1 KPFLZ, 3 TPZDN, 1 SRPD => 6 FBTG",
		"1 WHMLV, 3 BDSB, 2 JVGD => 9 LCSV",
		"13 XZKTG => 4 QVSH",
		"1 XHDVT => 7 XKFZ",
		"1 CMNH, 1 KQVS, 2 XVXCZ => 6 CGFML",
		"6 FLHW => 4 BRJDJ",
		"2 KQVL, 2 WGVP, 7 BMNJ, 11 KQVS, 1 HPNWP, 6 CRJZ => 4 DGJPN",
		"2 DZQJ, 1 BDSB => 2 DGMN",
		"1 XVXCZ, 4 MQRS => 3 WDFXH",
		"5 FLHW, 10 JPGK, 1 XZKTG => 4 QPNTQ",
		"2 LRBN => 9 FZVS",
		"149 ORE => 8 ZTBFN",
	}

	/*input = []string{
		"10 ORE => 10 A",
		"1 ORE => 1 B",
		"7 A, 1 B => 1 C",
		"7 A, 1 C => 1 D",
		"7 A, 1 D => 1 E",
		"7 A, 1 E => 1 FUEL",
	}

	input = []string{
		"9 ORE => 2 A",
		"8 ORE => 3 B",
		"7 ORE => 5 C",
		"3 A, 4 B => 1 AB",
		"5 B, 7 C => 1 BC",
		"4 C, 1 A => 1 CA",
		"2 AB, 3 BC, 4 CA => 1 FUEL",
	}*/

	var oreCount int64 = 0

	var reactionMap map[string][]string
	reactionMap = make(map[string][]string)

	var leftOverMap map[string]int64
	leftOverMap = make(map[string]int64)

	for _, reaction := range input {
		reactionComps := strings.Split(reaction, " => ")
		outcome := strings.Split(reactionComps[1], " ")
		reactionMap[outcome[1]] = []string{outcome[0], reactionComps[0]}
	}

	for key, _ := range leftOverMap {
		leftOverMap[key] = 0
	}

	followReaction(reactionMap, "FUEL", 1, leftOverMap, &oreCount)

	fmt.Printf("Outcome: %v", oreCount)
}

func followReaction(reactionMap map[string][]string, element string, required int64, leftOverElements map[string]int64, oreCount *int64) {

	fmt.Printf("%v\n", leftOverElements)

	if leftOverElements[element] >= required {
		leftOverElements[element] -= required
		return
	} else {
		required -= leftOverElements[element]
		leftOverElements[element] = 0
	}

	reactionComps := strings.Split(reactionMap[element][1], ", ")
	elementMultiple, _ := strconv.Atoi(reactionMap[element][0])

	fmt.Printf("Making %v of %v: %v\n", required, element, reactionMap[element])

	elementMade := 0
	for {
		for _, component := range reactionComps {
			requested := strings.Split(component, " ")
			amount, _ := strconv.Atoi(requested[0])
			element := requested[1]

			if element == "ORE" {
				*oreCount += int64(amount)
			} else {
				followReaction(reactionMap, element, int64(amount), leftOverElements, oreCount)
			}
		}
		if elementMade += elementMultiple; int64(elementMade) >= required {
			leftOverElements[element] += int64(elementMade) - required
			break
		}
	}

	fmt.Printf("oreCount %v\n", *oreCount)
}

func getMultiple(a int64, b int64) int64 {
	var multiple int64 = 0
	for b > 0 {
		multiple++
		b -= a
	}
	return multiple
}