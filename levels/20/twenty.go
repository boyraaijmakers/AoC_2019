package twenty

import (
	"fmt"
	"math"
	"sort"
)

func Twenty() {

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

	station := [2]int{17, 22}

	var comets [][2]int
	var metrics [][4]float64

	for i, item := range input {
		for j, coord := range item {
			if string(coord) == "#" {
				if [2]int{j, i} == station {
					continue
				}
				comets = append(comets, [2]int{j, i})
				metrics = append(metrics, [4]float64{
					angle(station, [2]int{j, i}),
					distance(station, [2]int{j, i}),
					float64(j*100 + i),
					float64(0),
				})
			}
		}
	}

	sort.SliceStable(metrics, func(i, j int) bool {
		return metrics[i][1] < metrics[j][1]
	})

	sort.SliceStable(metrics, func(i, j int) bool {
		return metrics[i][0] < metrics[j][0]
	})

	var lastAngle float64 = -1
	vaporCount := 0

	for {
		for i:=0; i < len(metrics); i++ {
			if math.Abs(metrics[i][0] - lastAngle) > 0.00001 && metrics[i][3] == 0 {
				lastAngle = metrics[i][0]
				vaporCount++

				metrics[i][3] = 1

				if vaporCount == 200 {
					fmt.Printf("Output: %v\n", metrics[i][2])
					break
				}
			}
		}
		if vaporCount > 200 {
			break
		}
	}
}

func distance(a [2]int, b [2]int) float64 {
	return math.Sqrt(math.Pow(float64(a[0]-b[0]), 2) + math.Pow(float64(a[1]-b[1]), 2))
}

func angle(a [2]int, b[2]int) float64 {
	deltaX := b[0] - a[0]
	deltaY := b[1] - a[1]

	var angle float64 = 0
	if deltaX == 0 {
		if deltaY > 0 {
			return math.Pi
		} else {
			return 0
		}
	}
	if deltaX > 0 {
		if deltaY >= 0 {
			return math.Atan(float64(deltaY)/float64(deltaX)) + math.Pi/2
		} else {  
			return math.Atan(float64(deltaX)/float64(-deltaY))
		}
	} else {
		if deltaY >= 0 {
			return math.Atan(float64(-deltaX)/float64(deltaY)) + math.Pi
		} else {
			return math.Atan(float64(-deltaY)/float64(-deltaX)) + 3*math.Pi/2
		}
		
	}
	return angle

}