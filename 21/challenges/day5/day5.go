package day5

import (
	"21/utils"
	"strings"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	data, file := utils.GetInput(5, mode)
	defer file.Close()
	lines := []*line{}
	lines2 := []*line{}
	lineId := 0
	for data.Scan() {
		var p1, p2 point
		lineId++
		l1 := data.Text()
		split := strings.Split(l1, "->")
		p1.newPoint(split[0])
		p2.newPoint(split[1])
		log.Debug("", "line", split, "p1", p1, "p2", p2)
		l := line{
			lineId: lineId,
			pointA: p1,
			pointB: p2,
		}
		l2 := line{
			lineId: lineId,
			pointA: p1,
			pointB: p2,
		}
		lines = append(lines, &l)
		lines2 = append(lines2, &l2)
	}
	log.Infof("Read total %d lines", len(lines))
	straightLines := filterLines(lines)
	log.Infof("%d lines are valid", len(straightLines))

	freqMap := getIntersectionPoints(straightLines)
	var pointsWhereAtLeastTwoLinesOverlap int
	for _, v := range freqMap {
		if v > 1 {
			pointsWhereAtLeastTwoLinesOverlap++
		}
	}
	log.Infof("Points where at least two lines intersect %d", pointsWhereAtLeastTwoLinesOverlap)

	log.Infof("Valid lines for updated value : %d", len(lines2))
	horizontalVerticalDiagonalFreqMap := getIntersectionPoints(lines2)
	var overlappingPoints int
	log.Debug("Printing the dictionary")
	for k, v1 := range horizontalVerticalDiagonalFreqMap {
		log.Debugf("%v -> %d", k, v1)
		if v1 > 1 {
			overlappingPoints++
		}
	}
	log.Infof("Updated intersection points : %d", overlappingPoints)
}

func getIntersectionPoints(lines []*line) map[point]int {
	pointFreqMap := map[point]int{}
	for _, l := range lines {
		l.findPointsOnLine()
		for _, p := range l.points {
			log.Debug(p)
			_, ok := pointFreqMap[p] // check if the point is already in the dictionary
			if !ok {                 // add it to the dict and set the freq to one
				pointFreqMap[p] = 1
			} else {
				pointFreqMap[p]++
			}
		}
	}
	// log.Debug(pointFreqMap)
	return pointFreqMap
}

func filterLines(lines []*line) []*line {
	usefulLines := []*line{}
	for _, line := range lines {
		if line.pointA.x == line.pointB.x || line.pointA.y == line.pointB.y {
			log.Debug(line)
			usefulLines = append(usefulLines, line)
		}
	}
	return usefulLines
}
