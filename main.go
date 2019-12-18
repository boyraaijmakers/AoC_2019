package main

import (
	twentyone "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/21"
	twentytwo "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/22"
	twentythree "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/23"
	twentyfour "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/24"
	twentyfive "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/25"
	twentysix "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/26"
	twentyseven "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/27"
	twentyeight "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/28"
	twentynine "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/29"
	thirty "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/30"
	thirtyone "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/31"
	thirtytwo "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/32"
	thirtythree "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/33"
	thirtyfour "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/34"
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
	seventeen "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/17"
	eighteen "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/18"
	nineteen "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/19"
	twenty "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/20"
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
	case 17:
		seventeen.Seventeen()
	case 18:
		eighteen.Eighteen()
	case 19:
		nineteen.Nineteen()
	case 20:
		twenty.Twenty()
	case 21:
		twentyone.Twentyone()
	case 22:
		twentytwo.Twentytwo()
	case 23:
		twentythree.Twentythree()
	case 24:
		twentyfour.Twentyfour()
	case 25:
		twentyfive.Twentyfive()
	case 26:
		twentysix.Twentysix()
	case 27:
		twentyseven.Twentyseven()
	case 28:
		twentyeight.Twentyeight()
	case 29:
		twentynine.Twentynine()
	case 30:
		thirty.Thirty()
	case 31:
		thirtyone.Thirtyone()
	case 32:
		thirtytwo.Thirtytwo()
	case 33:
		thirtythree.Thirtythree()
	case 34:
		thirtyfour.Thirtyfour()

	}
}
