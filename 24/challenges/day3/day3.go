package day3

import (
	"24/utils"
	"strconv"

	"regexp"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	data, file := utils.GetInput(3, mode)
	defer file.Close()

	var totalSum = 0
	var newSum = 0
	for data.Scan() {
		line := data.Text()
		log.Debug(line)
		matches := extractMul(line)
		totalSum += parseMul(matches)
		newMatches := extractMulDoDont(line)
		newSumVal := parseMulDoDont(newMatches)
		newSum += newSumVal
		log.Warn(newSumVal)

	}
	log.Infof("Adding results of all multiplications %d", totalSum)
	log.Infof("Adding results of modified instructionset  %d", newSum)

}

func parseMulDoDont(mulInstruction [][]string) int {
	parseInstruction := true
	sum := 0
	for _, instruction := range mulInstruction {
		switch instruction[0] {
		case "do()":
			log.Debug("got do", "ins", instruction)
			parseInstruction = true
		case "don't()":
			log.Debug("got dont", "ins", instruction)
			parseInstruction = false
		default:
			if parseInstruction {
				log.Debug("got mul", "ins", instruction)
				val := mul(instruction)
				sum += val
			} else {
				// log.Warn("skipping ", "ins", instruction)
				log.Debug("skipping ", "ins", instruction)
			}
		}
	}
	return sum
}

func mul(instruction []string) int {
	valA, _ := strconv.Atoi(instruction[1])
	valB, _ := strconv.Atoi(instruction[2])
	return valA * valB
}

func parseMul(mulInstruction [][]string) int {
	sum := 0
	for _, val := range mulInstruction {
		valA, _ := strconv.Atoi(val[1])
		valB, _ := strconv.Atoi(val[2])
		sum += valA * valB
	}
	return sum
}

func extractMulDoDont(input string) [][]string {
	pattern := `(?:` +
		`(?P<mul>mul\((?P<mul_a>\d+),(?P<mul_b>\d+)\))|` +
		`(?P<do>do\(\))|` +
		`(?P<dont>don't\(\))` +
		`)`

	// Compile the regex
	regex := regexp.MustCompile(pattern)

	// Find all matches
	matches := regex.FindAllStringSubmatch(input, -1)

	// Subgroup names to track
	subgroups := []string{"mul", "do", "dont"}

	// Process matches
	var results [][]string
	for _, match := range matches {
		// Find which group was matched
		for _, groupName := range subgroups {
			// Get the full match for each group
			fullMatch := match[regex.SubexpIndex(groupName)]
			if fullMatch != "" {
				// For mul, include the two numeric arguments
				if groupName == "mul" {
					a := match[regex.SubexpIndex("mul_a")]
					b := match[regex.SubexpIndex("mul_b")]
					results = append(results, []string{fullMatch, a, b})
				} else {
					// For do and don't, just include the full match
					results = append(results, []string{fullMatch})
				}
				break
			}
		}
	}

	return results
}

func extractMul(input string) [][]string {
	pattern := `mul\((\d+),(\d+)\)`
	// Compile the regex
	regex := regexp.MustCompile(pattern)
	// Find all matches
	matches := regex.FindAllStringSubmatch(input, -1)
	return matches
}
