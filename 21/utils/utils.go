package utils

import (
	"bufio"
	"fmt"
	"os"
	"path"

	"github.com/charmbracelet/log"
)

func GetInput(d int, mode string) (*bufio.Scanner, *os.File) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	day := fmt.Sprintf("day%d", d)
	fileName := fmt.Sprintf("%s.txt", mode)
	filePath := path.Join(wd, "challenges", day, fileName)
	log.Infof("Reading %s", filePath)

	readfile, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Trouble reading file", "fileName", filePath, "error", err)
	}
	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner, readfile
}
