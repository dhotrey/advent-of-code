package day6

import (
	"21/utils"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	data, file := utils.GetInput(6, mode)
	defer file.Close()

	var initialState []int
	fishCount := make([]int64, 9)

	for data.Scan() {
		input := data.Text()
		splitInput := strings.Split(input, ",")
		for _, i := range splitInput {
			timer, _ := strconv.Atoi(i)
			fishCount[timer]++
			initialState = append(initialState, timer)
		}
	}

	log.Infof("Got initial state %v", initialState)
	model := sim(80, &initialState)
	log.Debug(model)
	log.Infof("Got %d fishes after 80 days", len(model))

	days := 256
	updatedModel := sim2(fishCount, days)
	totalFish := 0

	for _, count := range updatedModel {
		totalFish += int(count)
	}

	log.Infof("Got %d fishes after %d days", totalFish, days)
}

func sim2(fishCount []int64, days int) []int64 {
	days--
	counts := make([]int64, len(fishCount))
	for i, count := range fishCount {
		log.Debugf("%d -> %d", i, count)
		if i == 0 {
			counts[8] += fishCount[0]
			counts[6] += fishCount[0]
		} else {
			counts[i-1] += fishCount[i]
		}
	}
	if days == 0 {
		return counts
	}
	return sim2(counts, days)
}

func sim(days int, initialState *[]int) []int {
	days--
	updatedState := []int{}
	log.Debugf("Days passed %d\n Fish state : %v", days, updatedState)
	for _, fish := range *initialState {
		fish--
		if fish == -1 {
			updatedState = append(updatedState, 6)
			updatedState = append(updatedState, 8)
		} else {
			updatedState = append(updatedState, fish)
		}
	}

	if days == 0 {
		return updatedState
	}
	return sim(days, &updatedState)
}
