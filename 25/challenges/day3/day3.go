package day3

import (
	"25/utils"
	"strconv"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	log.Info("Day3!!")
	data, file := utils.GetInput(3, mode)
	defer file.Close()

	sum := 0
	for data.Scan() {
		var tens, uPlace string
		line := data.Text()
		for tIdx, t := range line[:len(line)-1] {
			var units string
			if string(t) > tens {
				tens = string(t)
			} else {
				continue
			}
			for _, u := range line[tIdx+1:] {
				if string(u) > units {
					units = string(u)
				}
			}
			uPlace = units
		}
		joltage, _ := strconv.Atoi(tens + uPlace)
		log.Debugf("%s -> %d", line, joltage)
		sum += joltage
	}
	log.Infof("Total output joltage is %d", sum)
}
