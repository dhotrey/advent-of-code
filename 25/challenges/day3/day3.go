package day3

import (
	"25/utils"
	"slices"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	log.Info("Day3!!")
	data, file := utils.GetInput(3, mode)
	defer file.Close()

	sum1 := 0
	sum2 := 0
	for data.Scan() {
		line := data.Text()
		sum1 += getJoltage(line, 2)
		sum2 += getJoltage(line, 12)
	}
	log.Infof("Total output joltage is %d", sum1)
	log.Infof("Got joltage for p2  %d", sum2)
}

func getJoltage(line string, finalSize int) int {
	var joltageBldr string
	var offset int
	bb := strings.Split(line, "")
	log.Debug(bb)
	for i := finalSize - 1; i >= 0; i-- {
		slc := bb[offset : len(line)-i]
		max := slices.Max(slc)
		offset = slices.Index(slc, max) + offset + 1
		log.Debug("~", "i", i, "slc", slc, "max", max, "offset", offset)
		joltageBldr += max
	}

	joltage, _ := strconv.Atoi(joltageBldr)
	log.Debug(joltage)
	return joltage
}
