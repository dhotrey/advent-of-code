package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type funnySymbolLoc struct {
	locType  string
	position string
}

func main() {
	data, file := GetData("day3-input.txt")
	defer file.Close()
	lines := convertTo2dArr(data)
	var sumOfAllParts int
	var sumOfAllGearRatios int 
	positionLookUpT := map[string]func(lines [][]string, charLoc []int, loc funnySymbolLoc) int{
		"middle": getSumForAnyElement,
		"corner": getSumForAnyElement,
		"edge":   getSumForAnyElement,
	}
	for lineNumber, line := range lines {
		for charNum, char := range line {
			charLoc := []int{lineNumber, charNum}
			isAFunnySymbol := funnySymbolChecker(char)
			if isAFunnySymbol {
				position := determinePosition(charLoc, len(lines), len(line))
				sumCalculator := positionLookUpT[position.locType]
				sumOfNumbersAroundFunnySymbol := sumCalculator(lines, charLoc, position)
				sumOfAllParts += sumOfNumbersAroundFunnySymbol
				gearRatio := getGearRatio(lines,charLoc)
				sumOfAllGearRatios += gearRatio
			}
		}
	}
	fmt.Printf("sumOfAllParts: %v\n", sumOfAllParts)
	fmt.Printf("sumOfAllGearRatios: %v\n", sumOfAllGearRatios)
}

func getGearRatio(lines [][]string, charLoc []int) int {
	left := getLeftElement(lines, charLoc)
	right := getRightElement(lines, charLoc)
	top, topLeftExists, topRightExists := getTopElement(lines, charLoc)
	bottom, bottomLeftExists, bottomRightExists := getBottomElement(lines, charLoc)
	topLeft := getTopLeftElement(topLeftExists, lines, charLoc)
	topRight := getTopRightElement(topRightExists, lines, charLoc)
	bottomLeft := getBottomLeftElement(bottomLeftExists, lines, charLoc)
	bottomRight := getBottomRightElement(bottomRightExists, lines, charLoc)

	partNum := checkHowManyExist([]int{left,right,top,bottom,topLeft,topRight,bottomLeft,bottomRight})
	return partNum[0]*partNum[1]
}


func checkHowManyExist(partLocations[]int) []int {
	var partsAround []int  
	for _ , partNum := range partLocations {
		if partNum != 0 {
			partsAround = append(partsAround, partNum)
		}
	}
	if len(partsAround) == 2 {
		return partsAround
	} else {
		return []int{0,0}
	}	
}

func determinePosition(charLoc []int, totalLines int, totalChar int) funnySymbolLoc {
	charNum := charLoc[1]
	lineNum := charLoc[0]
	var loc funnySymbolLoc
	loc.locType = "middle"

	if lineNum == 0 && charNum == 0 {
		loc.locType = "corner"
		loc.position = "topLeft"
		return loc
	} else if lineNum == totalLines-1 && charNum == 0 {
		loc.locType = "corner"
		loc.position = "bottomLeft"
	} else if lineNum == 0 && charNum == totalChar-1 {
		loc.locType = "corner"
		loc.position = "topRight"
	} else if lineNum == totalLines-1 && charNum == totalChar-1 {
		loc.locType = "corner"
		loc.position = "bottomRight"
	} else if charNum == 0 {
		loc.locType = "edge"
		loc.position = "leftEdge"
	} else if lineNum == 0 {
		loc.locType = "edge"
		loc.position = "topEdge"
	} else if charNum == totalChar-1 {
		loc.locType = "edge"
		loc.position = "rightEdge"
	} else if lineNum == totalLines-1 {
		loc.locType = "edge"
		loc.position = "bottomEdge"
	} else {
		loc.locType = "middle"
	}

	return loc
}
func getSumForAnyElement(lines [][]string, charLoc []int, loc funnySymbolLoc) int {
	left := getLeftElement(lines, charLoc)
	right := getRightElement(lines, charLoc)
	top, topLeftExists, topRightExists := getTopElement(lines, charLoc)
	bottom, bottomLeftExists, bottomRightExists := getBottomElement(lines, charLoc)
	topLeft := getTopLeftElement(topLeftExists, lines, charLoc)
	topRight := getTopRightElement(topRightExists, lines, charLoc)
	bottomLeft := getBottomLeftElement(bottomLeftExists, lines, charLoc)
	bottomRight := getBottomRightElement(bottomRightExists, lines, charLoc)

	sum := left + right + top + bottom + topLeft + topRight + bottomLeft + bottomRight
	return sum
}

func getBottomRightElement(calculable bool, lines [][]string, charLoc []int) int {
	if !calculable {
		return 0
	}
	lineNum := charLoc[0]
	charNum := charLoc[1]
	lineBelow := lineNum + 1
	bottomRightChar := lines[lineBelow][charNum+1]

	rightChar := "#"
	doubleRightChar := "#"

	if rightElementAccessible(lines, []int{lineBelow, charNum + 1}) {
		rightChar = lines[lineBelow][charNum+2]
	}
	if doubleRightElementAccessible(lines, []int{lineBelow, charNum + 1}) {
		doubleRightChar = lines[lineBelow][charNum+3]
	}

	rightExists := charIsNum(rightChar)
	doubleRightExists := charIsNum(doubleRightChar)
	bottomRightExists := charIsNum(bottomRightChar)

	if !bottomRightExists {
		return 0
	} else if bottomRightExists && !rightExists {
		bottomRight, _ := strconv.Atoi(bottomRightChar)
		return bottomRight
	} else if bottomRightExists && rightExists && !doubleRightExists {
		num := bottomRightChar + rightChar
		bottomRight, _ := strconv.Atoi(num)
		return bottomRight
	} else {
		num := bottomRightChar + rightChar + doubleRightChar
		bottomRight, _ := strconv.Atoi(num)
		return bottomRight
	}
}

func getBottomLeftElement(calculable bool, lines [][]string, charLoc []int) int {
	if !calculable {
		return 0
	}
	lineNum := charLoc[0]
	charNum := charLoc[1]
	lineBelow := lineNum + 1
	bottomLeftChar := lines[lineBelow][charNum-1]

	leftChar := "#"
	doubleLeftChar := "#"

	if leftElementAccessible(lines, []int{lineBelow, charNum - 1}) {
		leftChar = lines[lineBelow][charNum-2]
	}
	if doubleLeftElementAccessible(lines, []int{lineBelow, charNum - 1}) {
		doubleLeftChar = lines[lineBelow][charNum-3]
	}

	leftExists := charIsNum(leftChar)
	doubleLeftExists := charIsNum(doubleLeftChar)
	bottomLeftExists := charIsNum(bottomLeftChar)

	if !bottomLeftExists {
		return 0
	} else if bottomLeftExists && !leftExists {
		bottomLeft, _ := strconv.Atoi(bottomLeftChar)
		return bottomLeft
	} else if bottomLeftExists && leftExists && !doubleLeftExists {
		num := leftChar + bottomLeftChar
		bottomLeft, _ := strconv.Atoi(num)
		return bottomLeft
	} else {
		num := doubleLeftChar + leftChar + bottomLeftChar
		bottomLeft, _ := strconv.Atoi(num)
		return bottomLeft
	}

}

func getTopRightElement(calculable bool, lines [][]string, charLoc []int) int {
	if !calculable {
		return 0
	}
	lineNum := charLoc[0]
	charNum := charLoc[1]
	lineAbove := lineNum - 1
	topRightChar := lines[lineAbove][charNum+1]

	rightChar := "#"
	doubleRightChar := "#"

	if rightElementAccessible(lines, []int{lineAbove, charNum + 1}) {
		rightChar = lines[lineAbove][charNum+2]
	}
	if doubleRightElementAccessible(lines, []int{lineAbove, charNum + 1}) {
		doubleRightChar = lines[lineAbove][charNum+3]
	}

	rightExists := charIsNum(rightChar)
	doubleRightExists := charIsNum(doubleRightChar)
	topRightExists := charIsNum(topRightChar)

	if !topRightExists {
		return 0
	} else if topRightExists && !rightExists {
		topRight, _ := strconv.Atoi(topRightChar)
		return topRight
	} else if topRightExists && rightExists && !doubleRightExists {
		num := topRightChar + rightChar
		topRight, _ := strconv.Atoi(num)
		return topRight
	} else {
		num := topRightChar + rightChar + doubleRightChar
		topRight, _ := strconv.Atoi(num)
		return topRight
	}

}

func getTopLeftElement(calculable bool, lines [][]string, charLoc []int) int {
	if !calculable {
		return 0
	}
	lineNum := charLoc[0]
	charNum := charLoc[1]
	lineAbove := lineNum - 1
	topLeftChar := lines[lineAbove][charNum-1]

	leftChar := "#"
	doubleLeftChar := "#"

	if leftElementAccessible(lines, []int{lineAbove, charNum - 1}) {
		leftChar = lines[lineAbove][charNum-2]
	}
	if doubleLeftElementAccessible(lines, []int{lineAbove, charNum - 1}) {
		doubleLeftChar = lines[lineAbove][charNum-3]
	}

	leftExists := charIsNum(leftChar)
	doubleLeftExists := charIsNum(doubleLeftChar)
	topLeftExists := charIsNum(topLeftChar)

	if !topLeftExists {
		return 0
	} else if topLeftExists && !leftExists {
		topLeft, _ := strconv.Atoi(topLeftChar)
		return topLeft
	} else if topLeftExists && leftExists && !doubleLeftExists {
		num := leftChar + topLeftChar
		topLeft, _ := strconv.Atoi(num)
		return topLeft
	} else {
		num := doubleLeftChar + leftChar + topLeftChar
		topLeft, _ := strconv.Atoi(num)
		return topLeft
	}
}

func getBottomElement(lines [][]string, charLoc []int) (int, bool, bool) {
	lineNum := charLoc[0]
	charNum := charLoc[1]
	lineBelow := lineNum + 1

	charBelow := lines[lineBelow][charNum]

	leftChar := "#"
	rightChar := "#"
	doubleLeftChar := "#"
	doubleRightChar := "#"

	if doubleLeftElementAccessible(lines, []int{lineBelow, charNum}) {
		doubleLeftChar = lines[lineBelow][charNum-2]
	}
	if doubleRightElementAccessible(lines, []int{lineBelow, charNum}) {
		doubleRightChar = lines[lineBelow][charNum+2]
	}
	if rightElementAccessible(lines, []int{lineBelow, charNum}) {
		rightChar = lines[lineBelow][charNum+1]
	}
	if leftElementAccessible(lines, []int{lineBelow, charNum}) {
		leftChar = lines[lineBelow][charNum-1]
	}

	leftExists := charIsNum(leftChar)
	rightExists := charIsNum(rightChar)
	charBelowExists := charIsNum(charBelow)
	doubleLeftExists := charIsNum(doubleLeftChar)
	doubleRightExists := charIsNum(doubleRightChar)

	if !charBelowExists {
		return 0, leftExists, rightExists
	} else if charBelowExists && leftExists && rightExists {
		num := leftChar + charBelow + rightChar
		numBelow, _ := strconv.Atoi(num)
		return numBelow, false, false
	} else if charBelowExists && !leftExists && !rightExists {
		numBelow, _ := strconv.Atoi(charBelow)
		return numBelow, false, false
	} else if charBelowExists && leftExists && !doubleLeftExists {
		num := leftChar + charBelow
		numBelow, _ := strconv.Atoi(num)
		return numBelow, false, false
	} else if charBelowExists && leftExists && doubleLeftExists {
		num := doubleLeftChar + leftChar + charBelow
		numBelow, _ := strconv.Atoi(num)
		return numBelow, false, false
	} else if charBelowExists && rightExists && !doubleRightExists {
		num := charBelow + rightChar
		numBelow, _ := strconv.Atoi(num)
		return numBelow, false, false
	} else if charBelowExists && rightExists && doubleRightExists {
		num := charBelow + rightChar + doubleRightChar
		numBelow, _ := strconv.Atoi(num)
		return numBelow, false, false
	} else {
		return 0, leftExists, rightExists
	}
}

func getTopElement(lines [][]string, charLoc []int) (int, bool, bool) {
	lineNum := charLoc[0]
	charNum := charLoc[1]
	lineAbove := lineNum - 1
	charAbove := lines[lineAbove][charNum]

	leftChar := "#"
	rightChar := "#"
	doubleLeftChar := "#"
	doubleRightChar := "#"

	if doubleLeftElementAccessible(lines, []int{lineAbove, charNum}) {
		doubleLeftChar = lines[lineAbove][charNum-2]
	}
	if doubleRightElementAccessible(lines, []int{lineAbove, charNum}) {
		doubleRightChar = lines[lineAbove][charNum+2]
	}
	if rightElementAccessible(lines, []int{lineAbove, charNum}) {
		rightChar = lines[lineAbove][charNum+1]
	}
	if leftElementAccessible(lines, []int{lineAbove, charNum}) {
		leftChar = lines[lineAbove][charNum-1]
	}

	leftExists := charIsNum(leftChar)
	rightExists := charIsNum(rightChar)
	charAboveExists := charIsNum(charAbove)
	doubleLeftExists := charIsNum(doubleLeftChar)
	doubleRightExists := charIsNum(doubleRightChar)

	if !charAboveExists {
		return 0, leftExists, rightExists
	} else if charAboveExists && leftExists && rightExists {
		num := leftChar + charAbove + rightChar
		numAbove, _ := strconv.Atoi(num)
		return numAbove, false, false
	} else if charAboveExists && !leftExists && !rightExists {
		numAbove, _ := strconv.Atoi(charAbove)
		return numAbove, false, false
	} else if charAboveExists && leftExists && !doubleLeftExists {
		num := leftChar + charAbove
		numAbove, _ := strconv.Atoi(num)
		return numAbove, false, false
	} else if charAboveExists && leftExists && doubleLeftExists {
		num := doubleLeftChar + leftChar + charAbove
		numAbove, _ := strconv.Atoi(num)
		return numAbove, false, false
	} else if charAboveExists && rightExists && !doubleRightExists {
		num := charAbove + rightChar
		numAbove, _ := strconv.Atoi(num)
		return numAbove, false, false
	} else if charAboveExists && rightExists && doubleRightExists {
		num := charAbove + rightChar + doubleRightChar
		numAbove, _ := strconv.Atoi(num)
		return numAbove, false, false
	} else {
		return 0, leftExists, rightExists
	}
}

// checks if it is possible to index the left element
func leftElementAccessible(lines [][]string, charLoc []int) bool {
	charNum := charLoc[1]
	return (charNum - 1) >= 0
}

// checks if it is possible to index the right elemenet
func rightElementAccessible(lines [][]string, charLoc []int) bool {
	charNum := charLoc[1]
	lineNum := charLoc[0]
	return (charNum + 2) <= len(lines[lineNum])
}

// checks if it is possible to index the element two elements to the left of charLoc
func doubleLeftElementAccessible(lines [][]string, charLoc []int) bool {
	charNum := charLoc[1]
	return (charNum - 2) >= 0
}

// checks if it is possible to index the element two elements to the right of charLoc
func doubleRightElementAccessible(lines [][]string, charLoc []int) bool {
	charNum := charLoc[1]
	lineNum := charLoc[0]

	return (charNum + 2) <= len(lines[lineNum])
}

// returns the numbers to the right of symbol
func getRightElement(lines [][]string, charLoc []int) int {
	lineNum := charLoc[0]
	charNum := charLoc[1]

	rightChar := "x"
	doubleRightChar := "x"
	tripleRightChar := "x"

	if rightElementAccessible(lines, []int{lineNum, charNum}) {
		rightChar = lines[lineNum][charNum+1]
	}
	if doubleRightElementAccessible(lines, []int{lineNum, charNum}) {
		doubleRightChar = lines[lineNum][charNum+2]
	}
	if doubleRightElementAccessible(lines, []int{lineNum, charNum + 1}) {
		tripleRightChar = lines[lineNum][charNum+3]
	}

	rightCharExists := charIsNum(rightChar)
	doubleRightExists := charIsNum(doubleRightChar)
	tripleRightExists := charIsNum(tripleRightChar)

	if !rightCharExists {
		return 0
	} else if rightCharExists && !doubleRightExists {
		right, _ := strconv.Atoi(rightChar)
		return right
	} else if rightCharExists && doubleRightExists && !tripleRightExists {
		num := rightChar + doubleRightChar
		right, _ := strconv.Atoi(num)
		return right
	} else {
		num := rightChar + doubleRightChar + tripleRightChar
		right, _ := strconv.Atoi(num)
		return right
	}
}

// returns the numbers to the left of symbol
func getLeftElement(lines [][]string, charLoc []int) int {
	lineNum := charLoc[0]
	charNum := charLoc[1]

	leftChar := "x"
	doubleLeftChar := "x"
	tripleLeftChar := "x"

	if leftElementAccessible(lines, []int{lineNum, charNum}) {
		leftChar = lines[lineNum][charNum-1]
	}
	if doubleLeftElementAccessible(lines, []int{lineNum, charNum}) {
		doubleLeftChar = lines[lineNum][charNum-2]
	}
	if doubleLeftElementAccessible(lines, []int{lineNum, charNum - 1}) {
		tripleLeftChar = lines[lineNum][charNum-3]
	}

	leftCharExists := charIsNum(leftChar)
	doubleLeftExists := charIsNum(doubleLeftChar)
	tripleLeftExists := charIsNum(tripleLeftChar)

	if !leftCharExists {
		return 0
	} else if leftCharExists && !doubleLeftExists {
		left, _ := strconv.Atoi(leftChar)
		return left
	} else if leftCharExists && doubleLeftExists && !tripleLeftExists {
		num := doubleLeftChar + leftChar
		left, _ := strconv.Atoi(num)
		return left
	} else {
		num := tripleLeftChar + doubleLeftChar + leftChar
		left, _ := strconv.Atoi(num)
		return left
	}

}

// cheks if the given character is funny
func funnySymbolChecker(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return s != "."
	} else {
		return false
	}
}

// returns true if a given character can be converted to an int
func charIsNum(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// convert input to 2d array
func convertTo2dArr(data *bufio.Scanner) [][]string {
	var lines [][]string
	for data.Scan() {
		line := convertStringToCharArr(data.Text())
		lines = append(lines, line)
	}
	return lines
}

// abcd becomes [a,b,c,d]
func convertStringToCharArr(input string) []string {
	var result []string
	for _, char := range input {
		result = append(result, string(char))
	}
	return result
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
