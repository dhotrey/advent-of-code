package day7

import (
	"25/utils"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
)

func beamSplitter(grid [][]string, rowIdx, splits int) ([][]string, int, int) {
	if rowIdx == len(grid) {
		return grid, rowIdx, splits
	}
	if rowIdx == 0 { // start condition
		log.Debug(grid[rowIdx])
		Sidx := slices.Index(grid[rowIdx], "S")
		rowIdx++
		grid[rowIdx][Sidx] = "|"
		log.Debug(grid[rowIdx])
		rowIdx++
	} else if slices.Contains(grid[rowIdx], "^") {
		for i, ele := range grid[rowIdx] {
			prev := grid[rowIdx-1][i]
			if ele == "^" && prev == "|" {
				grid[rowIdx][i-1] = "|"
				grid[rowIdx][i+1] = "|"
				splits++
			} else if prev == "|" {
				grid[rowIdx][i] = "|"
			}
		}
		log.Debug(grid[rowIdx])
		rowIdx++
	} else {
		for i := range grid[rowIdx] {
			ele := grid[rowIdx-1][i]
			if ele == "|" {
				grid[rowIdx][i] = "|"
			}
		}
		log.Debug(grid[rowIdx])
		rowIdx++
	}
	return beamSplitter(grid, rowIdx, splits)
}

func Sol(mode string) {
	data, file := utils.GetInput(7, mode)
	defer file.Close()

	grid := [][]string{}
	for data.Scan() {
		line := data.Text()
		row := strings.Split(line, "")
		grid = append(grid, row)
	}

	_, _, splits := beamSplitter(grid, 0, 0)

	log.Infof("Beam was split %d times", splits)

}
