package day2

import (
	"24/utils"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	data, file := utils.GetInput(2, mode)
	defer file.Close()
	var safe = 0
	var dampnedSafety = 0
	for data.Scan() {
		line := data.Text()
		reports := strings.Split(line, " ")
		reportInt := []int{}
		for _, r := range reports {
			rInt, _ := strconv.Atoi(r)
			reportInt = append(reportInt, rInt)
		}
		if isSafe(reportInt) {
			safe++
		}

		if isSafe2(reportInt) {
			dampnedSafety++
		}
	}
	log.Infof("%d entries are safe", safe)
	log.Infof("%d entries safe after problem dampner", dampnedSafety)
}

func isSafe2(report []int) bool {
	if sortedAscending(report) {
		log.Infof("%v conventionally sorted", report)
		return true
	}
	if sortedDescending(report) {
		log.Infof("%v conventionally sorted", report)
		return true
	}

	log.Warnf("checking dampning for %v", report)
	if isAscending(report) {
		return problemDampner(report, true)
	} else {
		return problemDampner(report, false)
	}
}

func problemDampner(report []int, asc bool) bool { // returns true if report is safe after removing one value
	log.Debugf("%v sorted asc = %v", report, asc)
	if asc {
		for idx, val := range report {
			cpy := make([]int, len(report))
			copy(cpy, report)
			newReport := append(cpy[:idx], cpy[idx+1:]...)
			log.Debugf("removing %d from %v. | new arr %v", val, report, newReport)
			if sortedAscending(newReport) {
				return true
			}
		}
		log.Warnf("%v cannot be dampned", report)
	} else {
		for idx, val := range report {
			cpy := make([]int, len(report))
			copy(cpy, report)
			newReport := append(cpy[:idx], cpy[idx+1:]...)
			log.Debugf("removing %d from %v. | new arr %v", val, report, newReport)
			if sortedDescending(newReport) {
				return true
			}
		}
		log.Warnf("%v cannot be dampned", report)
	}
	return false
}

func isSafe(report []int) bool {
	if sortedAscending(report) {
		log.Debugf("%v is sorted ascending", report)
		return true
	}
	if sortedDescending(report) {
		log.Debugf("%v is sorted descending", report)
		return true
	}
	return false
}

func isAscending(report []int) bool {
	sum := 0
	for i := 1; i < len(report); i++ {
		val := report[i] - report[i-1]
		sum += val
	}
	if sum > 0 {
		return true
	}
    return false
}

func sortedDescending(report []int) bool { //descending
	for i := 1; i < len(report); i++ {
		if report[i] > report[i-1] {
			return false
		}
		diff := report[i-1] - report[i]
		if diff > 3 || diff == 0 {
			return false
		}
	}
	return true
}

func sortedAscending(report []int) bool { // ascending
	log.Debug("checking report", "report", report)
	for i := 1; i < len(report); i++ {
		if report[i] < report[i-1] {
			return false
		}
		diff := report[i] - report[i-1]
		if diff > 3 || diff == 0 {
			return false
		}
	}
	return true
}
