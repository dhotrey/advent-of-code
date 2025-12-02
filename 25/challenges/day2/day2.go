package day2

import (
	"25/utils"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {

	log.Info("Day 1!!")

	data, file := utils.GetInput(2, mode)
	defer file.Close()
	data.Scan()
	line := data.Text()
	ranges := strings.Split(line, ",")
	sumOfIds := 0
	partTwoSum := 0

	for _, idRange := range ranges {
		ids := strings.Split(idRange, "-")
		start, _ := strconv.Atoi(ids[0])
		end, _ := strconv.Atoi(ids[1])
		log.Debugf(" %d -> %d ", start, end)

		for i := start; i <= end; i++ {
			if idIsInvalid(i) {
				sumOfIds += i
			}
			if idIsInvalidPartTwo(i) {
				partTwoSum += i
			}
		}
	}

	log.Info("Got sum of ids", "sum", sumOfIds)
	log.Info("Part 2 solution", "sol", partTwoSum)
}

func checkRepeating(length int, strId string) bool {
	if len(strId)%length == 0 {
		delim := len(strId) / length
		num := strId[0:delim]
		for i := range length {
			next := strId[i*delim : (i+1)*delim]
			if next != num {
				return false
			}
		}
		return true
	}
	return false
}

func idIsInvalidPartTwo(id int) bool {
	strId := strconv.Itoa(id)
	possibleCombinations := []int{5, 3, 2, 7}

	for _, combination := range possibleCombinations {
		if checkRepeating(combination, strId) {
			return true
		}
	}
	return false
}

func idIsInvalid(id int) bool {
	strId := strconv.Itoa(id)
	delim := int(len(strId) / 2)
	firstHalf := strId[0:delim]
	secondHalf := strId[delim:]

	if strings.Compare(firstHalf, secondHalf) == 0 {
		return true
	}
	return false
}
