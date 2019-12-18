package twentyfour

import (
	"fmt"
)

func Twentyfour() {

	planetsInit := [3][4]int64{
		{14, 17, 6, -2},
		{15, -3, 12, 10},
		{-2, 4, -13, -8},
	}

	planets := [3][4]int64{
		{14, 17, 6, -2},
		{15, -3, 12, 10},
		{-2, 4, -13, -8},
	}

	fullCycle := lcm(
			findCycle(planetsInit[0], planets[0]),
			findCycle(planetsInit[1], planets[1]),
			findCycle(planetsInit[2], planets[2]))

	fmt.Printf("A full cycle costs %v iterations", fullCycle)

}

func findCycle(coordsInit [4]int64, coords [4]int64) int64 {
	var velocs [4]int64
	var iterations int64 = 0
	for {
		iterations++
		for i, coordA := range coords {
			for _, coordB := range coords {
				if coordA != coordB {
					if velocs[i] += 1; coordA > coordB {
						velocs[i] -= 2;
					}
				}
			}
		}
		validity := true
		for i, _ := range coords {
			coords[i] += velocs[i]
			if coords[i] != coordsInit[i] || velocs[i] != 0 {
				validity = false
			}
		}
		if validity {
			break
		}
	}
	return iterations
}

func gcd(a, b int64) int64 {
	for b != 0 {
		t := b
		b, a = a % b, t
	}
	return a
}

func lcm(a, b, c int64) int64 {
	return (a * b / gcd(a, b)) * c / gcd(a * b / gcd(a, b), c)
}