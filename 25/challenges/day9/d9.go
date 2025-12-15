package day9

import (
	"25/utils"
	"math"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
	"gonum.org/v1/gonum/stat/combin"
)

func area(p1, p2 []int) int {
	l := p1[0] - p2[0] + 1
	w := p1[1] - p2[1] + 1
	area := math.Abs(float64(l * w))
	return int(area)
}

func Sol(mode string) {
	data, file := utils.GetInput(9, mode)
	defer file.Close()

	coordinates := [][]int{}
	for data.Scan() {
		line := data.Text()
		splitInput := strings.Split(line, ",")
		y, _ := strconv.Atoi(splitInput[1])
		x, _ := strconv.Atoi(splitInput[0])
		coordinates = append(coordinates, []int{y, x})
	}

	combin := combin.Combinations(len(coordinates), 2)
	log.Debug(coordinates)

	rectArea := -1
	for _, c := range combin {
		point1 := coordinates[c[0]]
		point2 := coordinates[c[1]]
		a := area(point1, point2)
		log.Debug("~>", "point1", point1, "point2", point2, "area", a)
		if a > rectArea {
			rectArea = a
		}
	}
	log.Info("Solution is ", "Sol", rectArea)
}
