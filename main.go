package main

import (
	"fmt"
	"os"
	"strconv"

	one "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/01"
	two "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/02"
	three "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/03"
	four "bitbucket.org/aimdeloittenl/aoc2019_boyraaijmakers/levels/04"
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
	}
}
