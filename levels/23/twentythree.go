package twentythree

import (
	"fmt"
	"math"
)

func Twentythree() {

	planets := [4][3]int64{
		{14, 15, -2},
		{17, -3, 4},
		{6, 12, -13},
		{-2, 10, -8},
	}

	var velocities [4][3]int64

	for i:=0; i < 1000; i++ {
		doMove(&planets, &velocities)
		/*for _, line := range velocities {
			fmt.Printf("%v\n", line)
		}*/
		println()

	}



	var totalEnergy int64 = 0

	for i, planet := range planets {
		var totalPot int64 = 0
		var totalKin int64 = 0
		for j, _ := range planet {
			totalPot += int64(math.Abs(float64(planets[i][j])))
			totalKin += int64(math.Abs(float64(velocities[i][j])))
		}

		totalEnergy += totalPot * totalKin
	}

	fmt.Printf("Output: %v", totalEnergy)

}

func doMove(planets *[4][3]int64, velocities *[4][3]int64) {
	for i, planetA := range planets {
		for _, planetB := range planets {
			if planetA == planetB {
				continue
			}

			for coord, _ := range planetA {
				if planetA[coord] < planetB[coord] {
					velocities[i][coord] += 1
				} else if planetA[coord] > planetB[coord] {
					velocities[i][coord] -= 1
				}
			}
		}
	}

	for i, planet := range planets {
		for j, _ := range planet {
			planets[i][j] += velocities[i][j]
		}
	}
}
