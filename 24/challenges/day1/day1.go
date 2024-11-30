package day1

import (
	"24/utils"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	log.Info("Aoc 2024 - d1")
	data, file := utils.GetInput(1, mode)
	defer file.Close()

	for data.Scan() {
		line := data.Text()
		log.Info(line)
	}

}
