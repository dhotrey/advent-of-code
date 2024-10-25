package day4

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type Board struct {
	winningBoard          bool
	stopCountingWins      bool
	id                    int
	numbersAnnouncedSoFar int
	numbersMarkedSoFar    map[string]int
	bitMap                uint32
	lookupT               map[string]int
	latestNum             string
}

func (b *Board) getBoardScore() int {
	sumOfUnmarkedNumbers := 0
	for k := range b.lookupT {
		_, marked := b.numbersMarkedSoFar[k]
		if !marked {
			val, _ := strconv.Atoi(k)
			sumOfUnmarkedNumbers += val
		}
	}
	latest, _ := strconv.Atoi(b.latestNum)
	return sumOfUnmarkedNumbers * latest
}

func (b *Board) winsAt(num string) bool {
	b.mark(num)
	return b.checkWin()
}

func (b *Board) mark(num string) {
	b.latestNum = num
	b.numbersAnnouncedSoFar++
	boardPosition, ok := b.lookupT[num]
	if ok {
		b.numbersMarkedSoFar[num] = boardPosition
		b.bitMap |= (1 << boardPosition)
		log.Debugf("Marked %d num in board %d | val %s at %d | : %025b", len(b.numbersMarkedSoFar), b.id, num, boardPosition, b.bitMap)
	}
}

func (b *Board) checkWin() bool {
	if b.numbersAnnouncedSoFar < 5 {
		return false // need at least 5 numbers in a row or col to win
	}
	if b.stopCountingWins {
		return false
	}

	for i := 0; i < 5; i++ {
		if b.isWinningColumn(i) || b.isWinningRow(i) {
			log.Infof("Board State %025b", b.bitMap)
			b.stopCountingWins = true
			return true
		}
	}
	return false
}

func (b *Board) isWinningRow(rowIdx int) bool {
	isWinning := ((b.bitMap>>uint((4-rowIdx)*5))&0b11111 == 0b11111)
	if isWinning {
		log.Infof("Board %d has row %d marked", b.id, 4-rowIdx)
	}
	return isWinning
}

func (b *Board) isWinningColumn(columnNumber int) bool {
	colElements := []uint32{}
	for i := columnNumber; i < 25; i += 5 { // can be further optimized to run without a for loop
		colElements = append(colElements, b.getBit(i))
	}
	isWinning := (colElements[0] & colElements[1] & colElements[2] & colElements[3] & colElements[4]) == uint32(1)

	if isWinning {
		log.Infof("Board %d has col %d marked", b.id, columnNumber)
	}
	return isWinning
}

func (b *Board) getBit(bitLoc int) uint32 {
	var mask uint32 = 1
	if bitLoc == 0 {
		return b.bitMap & mask
	}
	return (b.bitMap >> uint32(bitLoc)) & mask
}

func (b *Board) initBoard(initState string, bId int) {
	b.id = bId
	b.lookupT = make(map[string]int)
	b.numbersMarkedSoFar = make(map[string]int)
	values := strings.Split(initState, " ")
	log.Debug("got values", "val", values)
	loc := 0
	for _, val := range values {
		if val != "" {
			log.Debugf("got %s at %d", val, loc)
			b.lookupT[val] = loc
			loc++
		}
	}
}
