package day1

import (
	"24/utils"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	log.Info("Aoc 2024 - d1")
	data, file := utils.GetInput(1, mode)
	defer file.Close()
	leftLoc := []int{}
	rightLoc := []int{}

	for data.Scan() {
		line := data.Text()
		locationId := strings.Split(line, " ")
		left, _ := strconv.Atoi(locationId[0])
		right, _ := strconv.Atoi(locationId[3])
		leftLoc = append(leftLoc, left)
		rightLoc = append(rightLoc, right)
	}
	log.Debugf("Locations on the left %v", leftLoc)
	log.Debugf("Locations on the right %v", rightLoc)
	sort.Ints(leftLoc)
	sort.Ints(rightLoc)
	log.Debugf("left after sorting: %v", leftLoc)
	log.Debugf("right after sorting: %v", rightLoc)

	distances := []int{}
	for i := 0; i < len(leftLoc); i++ {
		d := math.Abs(float64(leftLoc[i]) - float64(rightLoc[i]))
		distances = append(distances, int(d))
	}
	log.Debugf("Distances : %v", distances)
	sum := 0
	for _, i := range distances {
		sum += i
	}
	log.Infof("Total distance between lists: %d", sum)

	similarity := getSimilarityScore(leftLoc, rightLoc)
	log.Infof("Got similarity score: %d", similarity)

}

func getSimilarityScore(leftArr, rightArr []int) int {
	sScore := 0
	numFreqMap := make(map[int]int)
	leftOccuranceArr := []int{}

	for _, left := range leftArr {
		if slices.Contains(leftOccuranceArr, left) {
			continue
		} else {
			leftOccuranceArr = append(leftOccuranceArr, left)
		}
		for _, right := range rightArr {
			if left == right {
				if numFreqMap[left] == 0 {
					numFreqMap[left] = 1
				} else {
					numFreqMap[left]++
				}
			}
		}
	}
	log.Debug(numFreqMap)
	log.Info("Calculating similarity score . . .")
	for _, num := range leftArr {
		s := num * numFreqMap[num]
		log.Debugf("similarity score for %d * %d = %d", num, numFreqMap[num], s)
		sScore += s
	}
	return sScore
}
