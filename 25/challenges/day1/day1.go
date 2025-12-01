package day1

import (
	"25/utils"
	"strconv"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	log.Info("Day 1!!")

	data, file := utils.GetInput(1, mode)
	defer file.Close()

	currPos := 50 // we start at 50
	log.Infof("Current position %d", currPos)
	zeroCount := 0
	zerosCrossed := 0
	for data.Scan() {
		line := data.Text()
		direction := string(line[0])
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		crossed, newPos, finalZeros := getZerosCrossed(currPos, distance, direction)
		zerosCrossed += crossed
		zeroCount += finalZeros
		currPos = newPos
		log.Debugf("op (%s) -> pos (%d) . zeros  crossed (%d)", line, currPos, zerosCrossed)
	}
	log.Infof("Password to open door is %d", zeroCount)
	log.Infof("Total zeros crossed is %d", zerosCrossed)
}

func getZerosCrossed(currPos int, distance int, direction string) (zerosCrossed int, newPos int, finalZeros int) {
	switch direction {
	case "L":
		newPos = currPos - distance

		distRemainder := distance % 100
		zerosCrossed = int(distance / 100)
		if newPos < 0 {
			newPos = newPos%100 + 100
		}

		if distRemainder >= currPos && currPos != 0 {
			zerosCrossed++
		}
	case "R":
		newPos = currPos + distance
		zerosCrossed = int(newPos / 100)

		newPos = newPos % 100
	}

	if newPos%100 == 0 { // part 1
		finalZeros++
	}

	return zerosCrossed, newPos, finalZeros
}
