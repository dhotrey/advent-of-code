package day9

import (
	"21/utils"
	"slices"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

const (
	topEdge = iota
	bottomEdge
	rightEdge
	leftEdge
	topLeftCorner
	topRightCorner
	bottomLeftCorner
	bottomRightCorner
	middleElement
)

var pointTypes = map[int]string{
	topEdge:           "top Edge",
	bottomEdge:        "bottom Edge",
	rightEdge:         "right Edge",
	leftEdge:          "left edge",
	topLeftCorner:     "top left corner",
	topRightCorner:    "top right corner",
	bottomLeftCorner:  "bottom left corner",
	bottomRightCorner: "bottom right corner",
	middleElement:     "middle element",
}

func Sol(mode string) {
	data, file := utils.GetInput(9, mode)
	defer file.Close()
	heightMap := [][]string{}

	for data.Scan() {
		log.Debug(data.Text())
		heightLine := strings.Split(data.Text(), "")
		heightMap = append(heightMap, heightLine)
	}
	log.Debug(heightMap)

	riskLevel, lowPts := getTotalRiskLevel(heightMap)
	log.Debug(lowPts)
	log.Info("Got risk level", "riskLevel", riskLevel)

	sinkSizeArr := []int{}
	for _, pt := range lowPts {
		size := getBasinSize(pt)
		log.Info(pt, "sinkSize", size+1)
		log.Info(" ")
		sinkSizeArr = append(sinkSizeArr, size+1)
	}
	slices.Sort(sinkSizeArr)
	log.Info(sinkSizeArr)
	sinkSizeSize := len(sinkSizeArr)
	largest := sinkSizeArr[sinkSizeSize-1]
	secondLargest := sinkSizeArr[sinkSizeSize-2]
	thirdLargest := sinkSizeArr[sinkSizeSize-3]

	log.Infof("Result : %d", largest*secondLargest*thirdLargest)

}

func getBasinSize(p point) int {
	var size int
	sinkPoints := []*point{&p}
	checkedEle := make(map[string]int)
	// cycles := 0

	for len(sinkPoints) != 0 {
		lowPt := sinkPoints[0]
		log.Debugf("Getting sink for : %s", lowPt)
		sinkPoints = sinkPoints[1:] // remove the 0th element

		surroundingPts := []point{lowPt.getLeft(), lowPt.getRight(), lowPt.getUp(), lowPt.getDown()}

		for _, pt := range surroundingPts {
			if pt.isValid {
				if pt.pointVal != 9 && pt.pointVal > p.pointVal {
					sha := pt.getSHA()
					_, ok := checkedEle[sha]
					if !ok {
						sinkPoints = append(sinkPoints, &pt)
						size++
						checkedEle[sha] = 1
					} else {
						log.Warnf("ignoring %s", pt)
					}
				}
			}
		}
		log.Debug(sinkPoints)
		sinkPoints = removeDuplicates(sinkPoints)
		log.Debug(" ")

		// cycles++
		// if cycles == 5 {
		// 	log.Fatal(" ")
		// }
	}

	return size
}

func removeDuplicates(sinks []*point) []*point {
	sinkSet := []*point{}
	sinkOccuranceMap := make(map[string]int)

	for _, p := range sinks {
		sha := p.getSHA()
		_, ok := sinkOccuranceMap[sha]
		if !ok {
			sinkOccuranceMap[sha] = 1
			sinkSet = append(sinkSet, p)
		}
	}

	if len(sinkSet) != len(sinks) {
		log.Warnf("removed %d points from list", len(sinks)-len(sinkSet))
		log.Debugf("new list %s \n", sinkSet)
	}
	return sinkSet
}

func getTotalRiskLevel(heightMap [][]string) (int, []point) { // TODO : refactor to use new abstractions
	lowPoints := []int{}
	positionIdx := []point{}
	for lIdx, line := range heightMap { // y coordinate
		for hIdx, height := range line { // x coordinate
			curr, _ := strconv.Atoi(height)

			p := newPoint(hIdx, lIdx, &heightMap)

			if lIdx == 0 { // top line
				below, _ := strconv.Atoi(heightMap[lIdx+1][hIdx])
				if hIdx == 0 { // top left corner
					right, _ := strconv.Atoi(line[hIdx+1])
					if curr < right && curr < below {
						lowPoints = append(lowPoints, curr)
						positionIdx = append(positionIdx, p)
					}
				} else if hIdx == len(line)-1 { // top right corner
					left, _ := strconv.Atoi(line[hIdx-1])
					if curr < left && curr < below {
						lowPoints = append(lowPoints, curr)
						positionIdx = append(positionIdx, p)
					}
				} else { // top edge
					left, _ := strconv.Atoi(line[hIdx-1])
					right, _ := strconv.Atoi(line[hIdx+1])
					if curr < left && curr < right && curr < below {
						lowPoints = append(lowPoints, curr)
						positionIdx = append(positionIdx, p)
					}

				}
			} else if lIdx == len(heightMap)-1 { // bottom line
				above, _ := strconv.Atoi(heightMap[lIdx-1][hIdx])
				if hIdx == 0 { // bottom left corner
					right, _ := strconv.Atoi(line[hIdx+1])
					if curr < right && curr < above {
						lowPoints = append(lowPoints, curr)
						positionIdx = append(positionIdx, p)
					}
				} else if hIdx == len(line)-1 { // bottom right corner
					left, _ := strconv.Atoi(line[hIdx-1])
					if curr < left && curr < above {
						lowPoints = append(lowPoints, curr)
						positionIdx = append(positionIdx, p)
					}
				} else { // bottm edge
					left, _ := strconv.Atoi(line[hIdx-1])
					right, _ := strconv.Atoi(line[hIdx+1])
					if curr < left && curr < right && curr < above {
						lowPoints = append(lowPoints, curr)
						positionIdx = append(positionIdx, p)
					}

				}
			} else if lIdx != 0 && hIdx == 0 || lIdx != len(heightMap)-1 && hIdx == 0 { // left edge without corners
				above, _ := strconv.Atoi(heightMap[lIdx-1][hIdx])
				below, _ := strconv.Atoi(heightMap[lIdx+1][hIdx])
				right, _ := strconv.Atoi(line[hIdx+1])
				if curr < above && curr < right && curr < below {
					lowPoints = append(lowPoints, curr)
					positionIdx = append(positionIdx, p)
				}

			} else if lIdx != 0 && hIdx == len(line)-1 || lIdx != len(heightMap)-1 && hIdx == len(line)-1 { // right edge without corners
				above, _ := strconv.Atoi(heightMap[lIdx-1][hIdx])
				below, _ := strconv.Atoi(heightMap[lIdx+1][hIdx])
				left, _ := strconv.Atoi(line[hIdx-1])
				if curr < above && curr < left && curr < below {
					lowPoints = append(lowPoints, curr)
					positionIdx = append(positionIdx, p)
				}

			} else { // point lies in the middle
				above, _ := strconv.Atoi(heightMap[lIdx-1][hIdx])
				below, _ := strconv.Atoi(heightMap[lIdx+1][hIdx])
				left, _ := strconv.Atoi(line[hIdx-1])
				right, _ := strconv.Atoi(line[hIdx+1])
				if curr < above && curr < left && curr < below && curr < right {
					lowPoints = append(lowPoints, curr)
					positionIdx = append(positionIdx, p)
				}
			}
		}
	}

	var riskLevel int
	for _, l := range lowPoints {
		riskLevel += l
	}
	return riskLevel + len(lowPoints), positionIdx
}
