package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var sum int
	var newSum int
	data, file := GetData("input.txt")
	defer file.Close()
	var calibrationValues []int
	var calibrationValues2 []int
	for data.Scan() {
		calval := getCalibrationValue(data.Text())
		calval2 := getCalibrationValueOtherWay(data.Text())
		calibrationValues = append(calibrationValues, calval)
		calibrationValues2 = append(calibrationValues2, calval2)
	}
	// fmt.Println(calibrationValues)
	for _, val := range calibrationValues {
		sum += val
	}
	for _, val := range calibrationValues2 {
		newSum += val
	}
	fmt.Println("The sum of all integer calibration values is :", sum)
	fmt.Println("The sum of all calibration values with different calculation is :", newSum)
}

// new way. I don't like this way.
func getCalibrationValueOtherWay(data string) int {
	calval := 0
	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	hasDigit, digits := checkSubstringForDigits(data, digits)
	if !hasDigit {
		return getCalibrationValue(data)
	}
	fdigit := getDigit(data, digits, true)
	sdigit := getDigit(data, digits, false)
	calval = fdigit*10 + sdigit
	return calval
}

func getDigit(data string, digitsPresent []string, firstDigit bool) (digit int) {
	if firstDigit {
		for i := 0; i < len(data); i++ {
			char := string(data[i])
			t, err := strconv.Atoi(char)
			if err == nil {
				substr := data[:i]
				hasDigit, digits := checkSubstringForDigits(substr, digitsPresent)
				if !hasDigit {
					digit = t
					return digit
				} else {
					return getAsciiDigit(substr, digits, firstDigit, true)
				}
			}
		}
		return getAsciiDigit(data, digitsPresent, firstDigit, true) // this condition arises when the string is of type xyzoneabtwo with no ints
	} else {
		for j := len(data) - 1; j >= 0; j-- {
			char := string(data[j])
			t, err := strconv.Atoi(char)
			if err == nil {
				substr := data[j+1:]
				hasDigit, digits := checkSubstringForDigits(substr, digitsPresent)
				if !hasDigit {
					digit = t
					return digit
				} else {
					return getAsciiDigit(substr, digits, false, true)
				}
			}
		}
		return getAsciiDigit(data, digitsPresent, firstDigit, true) // again condition arises when string is of type xyzonetwo with no ints
	}
}

// searches the given substr for digitsPresent to find their order of occurance then returns first or last occuring digit acc to firstDigit var.
func getAsciiDigit(substr string, digitsPresent []string, firstDigit bool, edge bool) int { // "twoneeighttow" ["one","two","eight","two","three"]
	asciiToInt := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	workingDigitscpy := make([]string, len(digitsPresent))
	workingSubStrcpy := substr[:]
	copy(workingDigitscpy, digitsPresent)

	loc := make(map[int]string) // map of key value pairs of elements and their index of occurance eg indices[0] = "two"
	var repeated []string       // repeated elements
	var indexed []string        // indexed elements
	var indices []int
	firstPass := true
	totalDigits := len(digitsPresent)
	for {
		for _, digit := range workingDigitscpy {

			idx := strings.Index(substr, digit)
			loc[idx] = digit
			if !eleInArr(indexed, digit) {
				indexed = append(indexed, digit)
				indices = append(indices, idx)
			} else if firstPass {
				repeated = append(repeated, digit)
			}
		}

		if len(workingDigitscpy) == 0 && len(repeated) != 0 {
			for _, digit := range repeated {
				idx := strings.Index(substr, digit)
				loc[idx] = digit

				indexed = append(indexed, digit)
				indices = append(indices, idx)
			}
		}
		firstPass = false

		sort.Ints(indices)

		if firstDigit {
			return asciiToInt[loc[indices[0]]]
		}

		if len(indexed) == totalDigits { // return condition for second digit
			return asciiToInt[loc[indices[totalDigits-1]]]
		}

		substr = strings.Replace(substr, repeated[0], strings.Repeat("x", len(repeated[0])), 1) // remove the first recurring element from str
		workingDigitscpy = popFromArray(workingDigitscpy, indexed)                              // remove the indexed elements from digits array
		if len(repeated) > 0 {
			indexed = popFromArray(indexed, repeated[:1])                    // remove the repeated element from indexed array
			indexed = append(indexed, strings.Repeat("x", len(repeated[0]))) // append a dummy element for exit condition
			if len(workingDigitscpy) != 0 {
				repeated = repeated[1:]
			}
			// remove the first repeated element from the repeated array
			// append a dummy element for exit condition
		}
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(workingSubStrcpy)
				fmt.Println(substr)
				fmt.Println("panic", r)
			}
		}()
	}

}

func popFromArray(source []string, itemsToRemove []string) []string {
	for _, item := range itemsToRemove {
		for i, value := range source {
			if value == item {
				source = append(source[:i], source[i+1:]...)
				break
			}
		}
	}
	return source
}

func eleInArr(arr []string, element string) bool {
	for _, e := range arr {
		if e == element {
			return true
		}
	}
	return false
}

// checks if a given substring provided has the digits from digits array , returns true and which digits are present
func checkSubstringForDigits(substr string, digits []string) (hasDigit bool, digitsPresent []string) {
	// [one,two,eight,two]

	for _, digit := range digits {
		occurance := strings.Count(substr, digit)
		if occurance > 0 {
			hasDigit = true
			for i := 0; i < occurance; i++ {
				digitsPresent = append(digitsPresent, digit)
			}
		}
	}
	return hasDigit, digitsPresent
}

func getCalibrationValue(data string) int {
	calval := 0
	for i := 0; i < len(data); i++ { // to get the digit in 10's place
		char := string(data[i])
		t, err := strconv.Atoi(char)
		if err == nil {
			calval = calval + t*10
			break
		}
	}

	for j := len(data) - 1; j >= 0; j-- { // traversing the string in reverse to get the last digit
		char := string(data[j])
		u, err := strconv.Atoi(char)
		if err == nil {
			calval = calval + u
			break
		}
	}

	return calval
}

func GetData(filename string) (*bufio.Scanner, *os.File) {
	readfile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner, readfile
}
