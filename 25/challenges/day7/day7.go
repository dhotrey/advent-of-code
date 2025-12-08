package day7

import (
	"25/utils"
	"slices"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func beamSplitter(grid [][]string, rowIdx, timelines int) ([][]string, int, int) {
	if rowIdx == len(grid) {
		return grid, rowIdx, timelines
	}
	if rowIdx == 0 { // start condition
		log.Debug(grid[rowIdx])
		Sidx := slices.Index(grid[rowIdx], "S")
		rowIdx++
		grid[rowIdx][Sidx] = "1"
		log.Debug(grid[rowIdx])
		rowIdx++
	} else if slices.Contains(grid[rowIdx], "^") {
		for i, ele := range grid[rowIdx] {
			prev := grid[rowIdx-1][i]
			prevNum, err := strconv.Atoi(prev)
			if ele == "^" && err == nil {
				currL := grid[rowIdx][i-1]
				currR := grid[rowIdx][i+1]

				lNum, lErr := strconv.Atoi(currL)
				rNum, rErr := strconv.Atoi(currR)

				if lErr == nil { // i.e. there is already a num there
					grid[rowIdx][i-1] = strconv.Itoa(lNum + prevNum)
				} else {
					grid[rowIdx][i-1] = prev
				}

				if rErr == nil {
					grid[rowIdx][i+1] = strconv.Itoa(rNum + prevNum)
				} else {
					grid[rowIdx][i+1] = prev
				}
				timelines++
			} else if prev != "." && prev != "^" {
				curr, err := strconv.Atoi(grid[rowIdx][i])
				if err == nil {
					grid[rowIdx][i] = strconv.Itoa(curr + prevNum)
				} else {
					grid[rowIdx][i] = prev
				}
			}
		}
		log.Debug(grid[rowIdx])
		rowIdx++
	} else {
		for i := range grid[rowIdx] {
			prev := grid[rowIdx-1][i]
			if prev != "." && prev != "^" {
				grid[rowIdx][i] = prev
			}
		}
		log.Debug(grid[rowIdx])
		rowIdx++
	}
	return beamSplitter(grid, rowIdx, timelines)
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

	mutatedGrid, _, splits := beamSplitter(grid, 0, 0)
	log.Infof("%d splits created", splits)

	timelines := 0
	for _, ele := range mutatedGrid[len(grid)-1] {
		num, err := strconv.Atoi(ele)
		if err == nil {
			timelines += num
		}
	}
	log.Infof("%d timelines created", timelines)

}
