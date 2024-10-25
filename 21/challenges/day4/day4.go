package day4

import (
	"21/utils"
	"bufio"
	"strings"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	data, file := utils.GetInput(4, mode)
	defer file.Close()
	nums := getBingoNumbers(data)
	log.Debug("Got input", "input", nums)
	boards := getBoards(data)
	log.Infof("Got %d Game boards", len(boards))
	firstWinner := getFirstWinner(nums, boards)
	score := firstWinner.getBoardScore()
	log.Infof("Board score that wins first is %d", score)

	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	boards2 := getBoards(scanner)
	lastWinner := getLastWinner(nums, boards2)
	score2 := lastWinner.getBoardScore()
	log.Infof("Board wins last is %d with score %d", lastWinner.id, score2)

}

func getLastWinner(bingoNumbers []string, bingoBoards []*Board) *Board {
	winner := Board{}
	boardsYetToWin := len(bingoBoards)
	for _, num := range bingoNumbers {
		log.Debugf("num %s called", num)
		for _, board := range bingoBoards {
			if board.winsAt(num) {
				boardsYetToWin--
				if boardsYetToWin == 0 {
					winner = *board
					break
				}
			}
		}
		if boardsYetToWin == 0 {
			break
		}
	}
	return &winner
}

func getFirstWinner(bingoNumbers []string, bingoBoards []*Board) *Board {
	winningBoard := Board{}
	for _, num := range bingoNumbers {
		log.Debugf("num %s called", num)
		for _, board := range bingoBoards {
			if board.winsAt(num) {
				board.winningBoard = true
				winningBoard = *board
				break
			}
		}
		if winningBoard.winningBoard {
			break
		}
	}
	return &winningBoard
}

func getBingoNumbers(data *bufio.Scanner) []string {
	data.Scan()
	return strings.Split(data.Text(), ",")
}

func getBoards(data *bufio.Scanner) []*Board {
	boards := []*Board{}
	boardId := 0
	for data.Scan() {
		switch data.Text() {
		case "":
			boardInfo := ""
			for i := 0; i < 5; i++ {
				data.Scan()
				row := data.Text()
				boardInfo += row + " "
			}
			var b Board
			b.initBoard(boardInfo, boardId)
			boards = append(boards, &b)
			boardId++
		}
	}
	return boards
}
