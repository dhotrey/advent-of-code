package day4

import (
	"25/utils"
	"strings"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	data, file := utils.GetInput(4, mode)
	defer file.Close()

	grid := [][]string{}
	for data.Scan() {
		line := data.Text()
		row := strings.Split(line, "")
		grid = append(grid, row)
	}
	prettyPrint(grid)
	accessible := 0
	for i := 1; ; i++ {
		moveableRoles := 0
		for rowIdx, row := range grid {
			for col, ele := range row {
				if ele == "@" {
					neighbours := getNeighbourCount(grid, rowIdx, col)
					if neighbours < 4 {
						log.Debug("~", "ele", ele, "location", []int{rowIdx, col}, "neighbours", neighbours)
						moveableRoles++
						grid[rowIdx][col] = "x"
					}
				}
			}
		}
		log.Debugf("Moved %d roles in iteration - %d", moveableRoles, i)
		prettyPrint(grid)
		cleanUp(grid)
		accessible += moveableRoles
		if moveableRoles == 0 { // no more roles can be extracted
			break
		}
	}
	log.Infof("%d roles of paper are accessible", accessible)
}

func cleanUp(grid [][]string) {
	for rowidx, row := range grid {
		for col, ele := range row {
			if ele == "x" {
				grid[rowidx][col] = "."
			}
		}
	}
}

func prettyPrint(grid [][]string) {
	for _, row := range grid {
		log.Debug(row)
	}
}

func getNeighbourCount(grid [][]string, rowIdx, columnIdx int) int {
	neighbours := 0
	rows := len(grid)
	cols := len(grid[0])
	directions := [][]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	// N , NE , E , SE , S , SW , W , NW

	for _, d := range directions {
		rowOp := d[0]
		colOp := d[1]
		newRow := rowIdx + rowOp
		newCol := columnIdx + colOp
		if (newRow >= 0 && newRow < rows) && (newCol >= 0 && newCol < cols) {
			ele := grid[newRow][newCol]
			if ele == "@" || ele == "x" {
				neighbours++
			}
		}
	}
	return neighbours
}
