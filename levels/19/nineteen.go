package nineteen

import (
	"fmt"
	"math"
)

func Nineteen() {

	input := []string{
		".#..#..#..#...#..#...###....##.#....",
		".#.........#.#....#...........####.#",
		"#..##.##.#....#...#.#....#..........",
		"......###..#.#...............#.....#",
		"......#......#....#..##....##.......",
		"....................#..............#",
		"..#....##...#.....#..#..........#..#",
		"..#.#.....#..#..#..#.#....#.###.##.#",
		".........##.#..#.......#.........#..",
		".##..#..##....#.#...#.#.####.....#..",
		".##....#.#....#.......#......##....#",
		"..#...#.#...##......#####..#......#.",
		"##..#...#.....#...###..#..........#.",
		"......##..#.##..#.....#.......##..#.",
		"#..##..#..#.....#.#.####........#.#.",
		"#......#..........###...#..#....##..",
		".......#...#....#.##.#..##......#...",
		".............##.......#.#.#..#...##.",
		"..#..##...#...............#..#......",
		"##....#...#.#....#..#.....##..##....",
		".#...##...........#..#..............",
		".............#....###...#.##....#.#.",
		"#..#.#..#...#....#.....#............",
		"....#.###....##....##...............",
		"....#..........#..#..#.......#.#....",
		"#..#....##.....#............#..#....",
		"...##.............#...#.....#..###..",
		"...#.......#........###.##..#..##.##",
		".#.##.#...##..#.#........#.....#....",
		"#......#....#......#....###.#.....#.",
		"......#.##......#...#.#.##.##...#...",
		"..#...#.#........#....#...........#.",
		"......#.##..#..#.....#......##..#...",
		"..##.........#......#..##.#.#.......",
		".#....#..#....###..#....##..........",
		"..............#....##...#.####...##.",
	}

	var comets [][2]int

	for i, item := range input {
		for j, coord := range item {
			if string(coord) == "#" {
				comets = append(comets, [2]int{j, i})
			}
		}
	}

	maxCount := 0
	maxCountCoords := [2]int{0, 0}

	for _, station := range comets {
		stationCanSee := 0

		for _, comet := range comets {
			if station == comet {
				continue
			}

			if canSee(comets, station, comet) {
				stationCanSee++
			}

		}

		//fmt.Printf("Location %v can see %v comets\n", station, stationCanSee)

		if stationCanSee > maxCount {
			maxCount = stationCanSee
			maxCountCoords = station
		}
	}

	fmt.Printf("The maximum value is: %v\n", maxCount)
	fmt.Printf("Found at location: %v\n", maxCountCoords)

}

func canSee(comets [][2]int, station [2]int, comet [2]int) bool {
	seeable := true

	for _, cometEl := range comets {
		if cometEl == comet || cometEl == station {
			continue
		}

		sumDist := distance(station, cometEl) + distance(cometEl, comet)
		totDist := distance(station, comet)

		//fmt.Printf("sumDist: %v, totDist: %v\n", sumDist, totDist)

		if math.Abs(sumDist-totDist) < 0.0000001 {
			//println("True")
			seeable = false
		}
	}
	return seeable
}

func distance(a [2]int, b [2]int) float64 {
	return math.Sqrt(math.Pow(float64(a[0]-b[0]), 2) + math.Pow(float64(a[1]-b[1]), 2))
}
