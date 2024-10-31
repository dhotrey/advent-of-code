// -1. identify 1,4,7,8
// 0. isolate C/F from 1
// 1. find A using 1 and 7
// 2. find B/D uisng 1 and 4
// 3. find e/g using 4,7,8
// 4. 5 segment display with c and f is 3
// 5. 6 segment display without both (c and f) is 6
// 6. 8 - 6 gives us C and thus f
// 7. 5 segment display that has c is 2
// 8. 5 segment display that does not have c
// 9. 2 - 3 gives us E
// 10. 6 segment display without E is 9
// 11. last one is 0
package day8

import (
	"21/utils"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	data, file := utils.GetInput(8, mode)
	defer file.Close()

	var one478Counter int
	var totalOutputVal int
	var counter int
	for data.Scan() {
		line := data.Text()
		splitLine := strings.Split(line, "|")
		digitsStr := splitLine[1]
		cipherStr := splitLine[0]
		cipher := strings.Split(strings.Trim(cipherStr, " "), " ")
		digits := strings.Split(strings.Trim(digitsStr, " "), " ")
		log.Debug("Solving", "cipher", cipher, "digits", digits)
		characterOccurance := count(digits) // part 1
		one478Counter += characterOccurance
		decipherKey := decipher(cipher)
		counter++
		displayVal := getDisplayValue(digits, decipherKey)
		totalOutputVal += displayVal
	}
	log.Infof("Got %d occurances of 1,4,7,8", one478Counter)
	log.Infof("Total of all output values %d", totalOutputVal)
}

func getDisplayValue(digits []string, key map[int]string) int {
	log.Debug("Getting display value")
	log.Debug("")
	var value int
	num1 := strings.Split(digits[0], "")
	slices.Sort(num1)
	log.Debug("", "num1", num1)
	num2 := strings.Split(digits[1], "")
	slices.Sort(num2)
	log.Debug("", "num2", num2)
	num3 := strings.Split(digits[2], "")
	slices.Sort(num3)
	log.Debug("", "num3", num3)
	num4 := strings.Split(digits[3], "")
	slices.Sort(num4)
	log.Debug("", "num4", num4)

	ones := -1
	tens := -1
	hundreds := -1
	thousands := -1

loop:
	for k, val := range key {
		valArr := strings.Split(val, "")
		slices.Sort(valArr)
		log.Debug("k-v pair", "valArr", valArr, "key", k)

		if slices.Equal(valArr, num1) && thousands == -1 {
			thousands = k
			log.Debug("Set thousand position val")
			goto loop
		} else if slices.Equal(valArr, num2) && hundreds == -1 {
			hundreds = k
			log.Debug("Set hundred position val")
			goto loop
		} else if slices.Equal(valArr, num3) && tens == -1 {
			tens = k
			log.Debug("Set ten position val")
			goto loop
		} else if slices.Equal(valArr, num4) && ones == -1 {
			ones = k
			log.Debug("Set one position val")
			goto loop
		}
	}
	value = thousands*1000 + hundreds*100 + tens*10 + ones
	return value
}

func decipher(cipher []string) map[int]string {
	uniqueSegmentSizes := map[int]int{2: 1, 4: 4, 3: 7, 7: 8} // map letter occurance to number
	decipherKey := map[int]string{1: "", 4: "", 7: "", 8: ""}
	nonUniqueSegmentSize := map[int][]string{5: {}, 6: {}}
	for _, digit := range cipher {
		size := len(digit)
		number, ok := uniqueSegmentSizes[size]
		if !ok {
			eleArr := nonUniqueSegmentSize[size]
			eleArr = append(eleArr, digit)
			nonUniqueSegmentSize[size] = eleArr
		} else {
			decipherKey[number] = digit
		}
	}
	log.Debug("Digits with size 5/6", "eleArr", nonUniqueSegmentSize)
	log.Debug("Got 1,4,7,8", "dict", decipherKey)

	cfVal := decipherKey[1]
	log.Debug("Step 0", "c/f val", cfVal)
	aVal := subtract(decipherKey[7], decipherKey[1])
	log.Debug("Step 1", "a Val", aVal)
	bdVal := subtract(decipherKey[1], decipherKey[4])
	log.Debug("Step 2", "b/d Val", bdVal)
	intermediate := subtract(decipherKey[8], decipherKey[4])
	egVal := subtract(intermediate, aVal)
	log.Debug("Step 3", "e/g Val", egVal)

	for _, val := range nonUniqueSegmentSize[5] {
		log.Debug("Checking for 3", "val", val)
		cProbable := strings.ContainsRune(val, rune(cfVal[0]))
		fProbable := strings.ContainsRune(val, rune(cfVal[1]))
		if cProbable && fProbable {
			log.Debug("Step 4", "Got value of 3", val)
			decipherKey[3] = val
			log.Debug("Values found so far", "key", decipherKey)
			break
		}
	}

	for _, val := range nonUniqueSegmentSize[6] {
		log.Debug("Checking for 6", "val", val)
		cProbable := strings.ContainsRune(val, rune(cfVal[0]))
		fProbable := strings.ContainsRune(val, rune(cfVal[1]))
		if !(cProbable && fProbable) {
			log.Debug("", "cProbable", cProbable, "fProbable", fProbable)
			log.Debug("Step 5", "Got value of 6", val)
			decipherKey[6] = val
			log.Debug("Values found so far", "key", decipherKey)
			break
		}
	}

	cValue := subtract(decipherKey[8], decipherKey[6])
	fValue := subtract(cfVal, cValue)
	log.Debug("Step 6", "c value", cValue, "f value", fValue)

	for _, val := range nonUniqueSegmentSize[5] {
		if val == decipherKey[3] {
			continue
		}
		if strings.ContainsRune(val, rune(cValue[0])) {
			log.Debug("Step 7", "Got value of 2", val)
			decipherKey[2] = val
		}
		if strings.ContainsRune(val, rune(fValue[0])) {
			log.Debug("Step 8", "Got value of 5", val)
			decipherKey[5] = val
		}
	}
	log.Debug("Values found so far", "key", decipherKey)

	modified2Val := decipherKey[2] + fValue
	log.Debug("", "modified 2 value", modified2Val)
	eValue := subtract(modified2Val, decipherKey[3])
	log.Debug("Step 9", "Got value of e", eValue)

	for _, val := range nonUniqueSegmentSize[6] {
		if val == decipherKey[6] {
			continue
		}
		if strings.Contains(val, eValue) {
			log.Debug("Step 10", "Got value of 0", val)
			decipherKey[0] = val
		} else {
			log.Debug("Step 11", "Got value of 9", val)
			decipherKey[9] = val
		}
	}
	log.Debug("Values found so far", "key", decipherKey)
	return decipherKey
}

// does digit1 - digit2 and returns the result
func subtract(digit1, digit2 string) string {
	var larger, smaller string
	d1Size := len(digit1)
	d2Size := len(digit2)

	if d1Size > d2Size {
		larger = digit1
		smaller = digit2
	} else {
		larger = digit2
		smaller = digit1
	}
	largeArr := strings.Split(larger, "")
	smallArr := strings.Split(smaller, "")
	resultArr := []string{}

	for _, ele := range largeArr {
		if !slices.Contains(smallArr, ele) {
			resultArr = append(resultArr, ele)
		}

	}
	return strings.Join(resultArr, "")
}

func count(digits []string) int {
	var count int
	sevenSegDisplay := map[int]int{2: 1, 4: 4, 3: 7, 7: 8} // map letter occurance to number
	for _, digit := range digits {
		size := len(digit)
		log.Debug("", "digit", digit, "size", size)
		_, ok := sevenSegDisplay[size]
		if !ok {
			continue
		} else {
			count++
		}
	}
	return count
}
