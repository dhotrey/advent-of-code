package day11

import (
	"21/utils"
	"slices"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	data, file := utils.GetInput(11, mode)
	defer file.Close()
	energyLevels := [][]int{}

	for data.Scan() {
		line := data.Text()
		lineArr := strings.Split(line, "")
		intArr := []int{}
		for _, ele := range lineArr {
			intEle, _ := strconv.Atoi(ele)
			intArr = append(intArr, intEle)
		}
		energyLevels = append(energyLevels, intArr)
	}

	log.Debug("Orignal Arr")
	printArr(energyLevels)
	totalFlashes := getFlashes(100, energyLevels)
	log.Infof("Total flashes after 100 steps is : %d", totalFlashes)
}

func getFlashes(steps int, energyLevels [][]int) int {
	var flashes int
	for i := 0; i < steps; i++ {
		log.Debugf("Step %d", i+1)
		increment(energyLevels)
		printArr(energyLevels)

		if maxEnergyLevelPresent(energyLevels) {

		}

	}
	return flashes
}

func maxEnergyLevelPresent(arr [][]int) bool {
	for _, line := range arr {
		if slices.Contains(line, 9) {
			return true
		}
	}
	return false
}

func increment(arr [][]int) {
	log.Debug("Incrementing by 1")
	for lIdx, line := range arr {
		for eleIdx := range line {
			arr[lIdx][eleIdx]++
		}
	}
}

func printArr(arr [][]int) {
	for _, line := range arr {
		log.Debug(line)
	}
}
