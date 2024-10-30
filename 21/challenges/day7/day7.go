package day7

import (
	"21/utils"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	data, file := utils.GetInput(7, mode)
	defer file.Close()
	data.Scan()
	input := data.Text()
	splitInput := strings.Split(input, ",")
	horizontalPosition := []int{}

	for _, i := range splitInput {
		pos, _ := strconv.Atoi(i)
		horizontalPosition = append(horizontalPosition, pos)
	}

	log.Debug("Got initial Horizontal Positions", "pos", horizontalPosition)
	log.Debugf("Total horizontal positions : %d", len(horizontalPosition))

	fuelConsumptionArr := getFuelConsumption(horizontalPosition)
	log.Debug(fuelConsumptionArr)

	minFuelConsumption := slices.Min(fuelConsumptionArr)
	log.Infof("Minimum fuel consumption is %d", minFuelConsumption)

	newFuelCost := getNewFuelConsumption(horizontalPosition)
	newMinFuelConsumption := slices.Min(newFuelCost)
	log.Debug(newFuelCost)
	log.Infof("updated fuel consumption is %d", newMinFuelConsumption)
}

func getNewFuelConsumption(positions []int) []int {
	maxDist := slices.Max(positions)
	cost := make([]int, maxDist)

	for i := 0; i < maxDist; i++ {
		for _, pos := range positions {
			n := math.Abs(float64(pos - i))
			fuelConsumed := (n * (n + 1)) / 2
			cost[i] += int(fuelConsumed)
		}
	}
	return cost
}

func getFuelConsumption(positions []int) []int {
	fuelConsumption := make([]int, len(positions))
	for i := 0; i < len(positions); i++ {
		for _, pos := range positions {
			fuelConsumed := math.Abs(float64(i - pos))
			fuelConsumption[i] += int(math.Abs(float64(fuelConsumed)))
		}
	}
	return fuelConsumption
}
