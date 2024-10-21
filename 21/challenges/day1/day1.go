package day1

import (
	"21/utils"
	"strconv"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	log.Info("Running in", "mode", mode)
	data, file := utils.GetInput(1, mode)
	defer file.Close()
	depths := []int{}
	window := []int{}

	var recordIncrement int
	var prev int
	var current int

	data.Scan()
	prev, _ = strconv.Atoi(data.Text())
	depths = append(depths, prev)

	for data.Scan() {
		current, _ = strconv.Atoi(data.Text())
		depths = append(depths, current)
		if current > prev {
			recordIncrement++
		}
		prev = current
	}
	log.Info("Part 1 solution", "Total Increments", recordIncrement)

	for i := 0; i+2 <= len(depths)-1; i++ {
		w := depths[i] + depths[i+1] + depths[i+2]
		log.Debugf("%d) %d + %d + %d = %d", i, depths[i], depths[i+1], depths[i+2], w)
		window = append(window, w)
	}
	log.Debug(window)
	var depthIncrementCounter int
	prevDepth := window[0]

	for _, currentDepth := range window[1:] {
		if currentDepth > prevDepth {
			depthIncrementCounter++
			log.Debugf("Incrementing : %d > %d - depthCounter : %d", currentDepth, prevDepth, depthIncrementCounter)
		}
		prevDepth = currentDepth
	}
	log.Info("Part 2 Solution", "Total increment in depths", depthIncrementCounter)

}
