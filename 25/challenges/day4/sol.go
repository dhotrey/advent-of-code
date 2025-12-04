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
	log.Debug(grid)

	accessible := 0
	for rowIdx, row := range grid {
		for col, ele := range row {
			if ele == "@" {
				neighbours := getNeighbours(grid, rowIdx, col)
				if len(neighbours) < 4 {
					log.Debug("~", "ele", ele, "location", []int{rowIdx, col}, "neighbours", neighbours)
					accessible++
				}
			}
		}
	}
	log.Infof("%d roles of paper are accessible", accessible)
}

func getNeighbours(grid [][]string, rowIdx, columnIdx int) []string {
	neighbours := []string{}
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
			if ele == "@" {
				neighbours = append(neighbours, ele)
			}
		}
	}
	return neighbours
}
