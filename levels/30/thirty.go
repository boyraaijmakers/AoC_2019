package thirty

import "fmt"

func Thirty() {

	input := []string{
		"  # # # # #   # # # # # # #   # # #   # # #   # # #   # # # # # # # # # # # # #  ",
		"#           #               #       #       #       #                           #",
		"  # #   # # #   # # #   #   #   #   #   #   #   #   # # # # # # #   # # #   #   #",
		"#       #       #       #   #   #       #       #   #           #   #   #   #   #",
		"#   #   #   # # #   # # # # #   # # # # # # # # #   #   # # #   #   #   #   # #  ",
		"#   #   #       #           #   #               #           #   #       #       #",
		"#   # # # # #   # # # # #   #   #   # # # # #   # # # # # # #   #   # # # # #   #",
		"#           #   # ! #       #   #           #       #   #       #   #       #   #",
		"#   #   # # #   #   #   # # #   # # # # #   # # #   #   #   # # #   #   #   #   #",
		"#   #               #       #           #       #       #   #       #   #       #",
		"#   # # # # # # # # # # #   # # # # #   # # #   #   # # #   #   # # #   # # #   #",
		"#   #           #                   #       #   #   #       #       #       #   #",
		"#   #   # # #   #   # # # # # # # # # # #   #   #   #   # # # # # # # # #   #   #",
		"#   #   #   #   #                               #   #           #       #   #   #",
		"#   #   #   #   # # # # # # # # # # # # # # #   # # # # # # #   #   #   #   #   #",
		"#       #   #   #           #               #   #               #   #       #   #",
		"  # # # #   #   #   # # #   #   # # # # #   #   #   # # # # # # #   # # # # #   #",
		"#           #   #       #   #       #       #   #               #           #   #",
		"#   #   # # #   #   #   # # # # #   #   # #   # # # # # # # #   # # #   #   #   #",
		"#   #       #   #   #           #   #     @ #               #       #   #   #   #",
		"#   # # #   #   #   # # #   # # #   # # # # #   # # # # #   # # #   #   #   #   #",
		"#       #       #       #       #   #           #       #           #   #   #   #",
		"  # #   # # # # #   # # # # #   #   #   # # #   #   #   # # # # # # # # #   #   #",
		"#   #   #           #           #   #       #   #   #                   #   #   #",
		"#   #   # # #   # # #   # # # # #   # # #   #   #   # # # # # # #   #   #   #   #",
		"#   #       #   #       #           #       #   #       #   #       #       #   #",
		"#   # # #   # # #   # # #   # # # # # # # # #   # # #   #   #   # # # # # # #   #",
		"#       #           #       #               #   #           #   #       #   #   #",
		"  # #   # # #   # # #   # # # # #   # # #   #   #   # # # # #   #   #   #   #   #",
		"#           #   #       #       #       #   #   #   #       #   #   #       #   #",
		"#   # # #   #   #   # # #   #   # # #   #   #   # # #   #   #   #   # # # # #   #",
		"#   #       #   #           #   #       #       #       #   #   #   #       #   #",
		"#   # # # # #   # # # # # # #   #   # # # # # # #   # # #   #   #   #   #   #   #",
		"#       #       #           #       #               #   #       #   #   #       #",
		"#   #   #   # # # # #   #   # # # # # # #   # # # # #   # # # # #   #   # # # #  ",
		"#   #       #           #   #               #                   #       #       #",
		"#   # # # # #   #   # # #   #   # # # # # # #   # # # # #   #   #   # # # # #   #",
		"#       #       #   #   #   #   #   #       #   #           #   #   #           #",
		"  # #   # # # # #   #   #   #   #   #   #   #   #   # # # # # # #   # # #   #   #",
		"#                   #           #       #       #                           #   #",
		"  # # # # # # # # #   # # # # #   # # #   # # #   # # # # # # # # # # # # #   #  ",

	}

	/*input = []string{
		"  # #      ",
		"#     # #  ",
		"#   #     #",
		"#   !   #  ",
		"# # # #    ",
	}*/

	print("\033[H\033[2J")

	var grid [][]int64
	var start [2]int64

	for i:=0;i<len(input);i++ {

		var line []int64
		for j:=0;j<len(input[0]);j +=2{
			if input[i][j] == '#' {
				line = append(line,-1)
			} else if input[i][j] == '!' {
				start = [2]int64{int64(i), int64(j/2)}
				line = append(line,0)
			} else {
				line = append(line, 0)
			}
		}
		grid = append(grid, line)
	}

	fmt.Printf("%v\n", start)
	walkGrid(grid, start)

	var maxVal int64 = 0
	for _, row := range grid {
		for _, el := range row {
			if el == -1 {
				print("# ")
			} else if el == 0 {
				print(". ")
			} else {
				print("O ")
			}
			if el > maxVal {
				maxVal = el
			}
		}
		println()
	}

	fmt.Printf("Output: %v", maxVal)

}

func walkGrid(grid [][]int64, position [2]int64) {
	fmt.Printf("\033[0;0H")
	for _, row := range grid {
		for _, el := range row {
			if el == -1 {
				print("# ")
			} else if el == 0 {
				print(". ")
			} else {
				print("O ")
			}
		}
		println()
	}
	println()

	currentVal := grid[position[0]][position[1]]

	if grid[position[0] + 1][position[1]] == 0 {
		grid[position[0] + 1][position[1]] = currentVal + 1
		walkGrid(grid, [2]int64{position[0] + 1, position[1]})
	}
	if grid[position[0] - 1][position[1]] == 0 {
		grid[position[0] - 1][position[1]] = currentVal + 1
		walkGrid(grid, [2]int64{position[0] - 1, position[1]})
	}
	if grid[position[0]][position[1] + 1] == 0 {
		grid[position[0]][position[1] + 1] = currentVal + 1
		walkGrid(grid, [2]int64{position[0], position[1] + 1})
	}
	if grid[position[0]][position[1] - 1] == 0 {
		grid[position[0]][position[1] - 1] = currentVal + 1
		walkGrid(grid, [2]int64{position[0], position[1] - 1})
	}
}