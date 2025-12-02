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

		switch direction {
		case "L":
			dist := distance % 100
			zerosCrossed += (distance - dist) / 100
			for _ = range dist {
				currPos = ((currPos-1)%100 + 100) % 100
				if currPos == 0 {
					zerosCrossed++
				}
			}

		case "R":
			dist := distance % 100
			zerosCrossed += (distance - dist) / 100
			for _ = range dist {
				currPos = (currPos + 1) % 100
				if currPos == 0 {
					zerosCrossed++
				}
			}
		}

		log.Debugf("op (%s) -> pos (%d) . zeros  crossed (%d)", line, currPos, zerosCrossed)
	}
	log.Infof("Password to open door is %d", zeroCount)
	log.Infof("Total zeros crossed is %d", zerosCrossed)
}
