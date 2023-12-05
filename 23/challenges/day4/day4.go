package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, file := GetData("day4-input.txt")
	defer file.Close()
	var totalPoints int
	cardOccurancemap := make(map[int]int)
	for data.Scan() {
		cNum, winningNums, myNums := cleanData(data.Text())
		points := getPoints(cNum, winningNums, myNums)
		totalPoints += points
		//p2
		cardNum, _ := strconv.Atoi(cNum)
		cardOccurancemap[cardNum] += 1
		currentCardCount := cardOccurancemap[cardNum]
		cardsAhead := getNumberOfNextCards(cNum, winningNums, myNums)
		for i := 0; i < cardsAhead; i++ {
			cardAhead := cardNum + i + 1
			cardOccurancemap[cardAhead] += currentCardCount * 1
		}
	}
	var totalCards int
	for _, v := range cardOccurancemap {
		totalCards += v
	}
	fmt.Printf("totalCards: %v\n", totalCards)
	// fmt.Printf("cardOccurancemap: %v\n", cardOccurancemap)
	// for k,v := range cardOccurancemap {
	// 	fmt.Printf("%d : %d\n",k,v)
	// }
	fmt.Printf("totalPoints: %v\n", totalPoints)

}

func getNumberOfNextCards(cardnum string, winningNums []string, myNums []string) int {
	var points int
	for _, num := range myNums {
		if eleInArr(winningNums, num) {
			points++
		}
	}
	return points
}

func getPoints(cardnum string, winningNums []string, myNums []string) int {
	var points int

	for _, num := range myNums {
		if eleInArr(winningNums, num) {
			points++
		}
	}

	return int(math.Pow(2, float64(points-1)))
}

func eleInArr(arr []string, element string) bool {
	for _, e := range arr {
		if e == element {
			return true
		}
	}
	return false
}

func cleanData(numbers string) (string, []string, []string) {
	splitInput := strings.Split(numbers, "|")
	myNumbers := convertStringToCharArr(strings.TrimSpace(splitInput[1]))
	notQuiteWinningNumbers := strings.Split(strings.TrimSpace(splitInput[0]), ":")
	winningNumbers := convertStringToCharArr(strings.TrimSpace(notQuiteWinningNumbers[1]))
	notQuiteCardNumber := strings.Split(strings.TrimSpace(notQuiteWinningNumbers[0]), " ")
	cardNumber := strings.TrimSpace(notQuiteCardNumber[len(notQuiteCardNumber)-1])
	return cardNumber, winningNumbers, myNumbers
}

// "12 34 66  7" becomes [12,34,66,7]
func convertStringToCharArr(input string) []string {
	var result []string
	splitStr := strings.Split(input, " ")
	for _, char := range splitStr {
		char = strings.TrimSpace(char)
		if char != "" {
			result = append(result, string(char))
		}
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
