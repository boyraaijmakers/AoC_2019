package main

import (
	"fmt"
	"os"
	"strconv"

	one "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/01"
	two "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/02"
	three "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/03"
	four "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/04"
	five "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/05"
	six "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/06"
	seven "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/07"
	eight "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/08"
	nine "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/09"
	ten "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/10"
	eleven "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/11"
	twelve "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/12"
	thirteen "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/13"
	fourteen "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/14"
	fifteen "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/15"
	sixteen "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/16"
)

func main() {
	level, _ := strconv.Atoi(os.Args[1])

	fmt.Printf("Looking for level %v\n", level)

	switch level {
	case 1:
		one.One()
	case 2:
		two.Two()
	case 3:
		three.Three()
	case 4:
		four.Four()
	case 5:
		five.Five()
	case 6:
		six.Six()
	case 7:
		seven.Seven()
	case 8:
		eight.Eight()
	case 9:
		nine.Nine()
	case 10:
		ten.Ten()
	case 11:
		eleven.Eleven()
	case 12:
		twelve.Twelve()
	case 13:
		thirteen.Thirteen()
	case 14:
		fourteen.Fourteen()
	case 15:
		fifteen.Fifteen()
	case 16:
		sixteen.Sixteen()
	}
}
