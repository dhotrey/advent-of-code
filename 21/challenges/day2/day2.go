package day2

import (
	"21/utils"
	"bufio"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	data, file := utils.GetInput(2, mode)
	defer file.Close()

	var depth int
	var forward int
	for data.Scan() {
		cmd := strings.Split(data.Text(), " ")
		val, _ := strconv.Atoi(cmd[1])
		switch cmd[0] {
		case "forward":
			forward += val
		case "down":
			depth += val
		case "up":
			depth -= val
		}
	}
	log.Infof("Final horizontal product :%d", depth*forward)

	file.Seek(0, 0)
	data = bufio.NewScanner(file)
	var aim int
	var newDepth int
	var horizontal int

	for data.Scan() {
		log.Debug(data.Text())
		cmd := strings.Split(data.Text(), " ")
		val, _ := strconv.Atoi(cmd[1])
		switch cmd[0] {
		case "forward":
			horizontal += val
			newDepth += aim * val
			log.Debugf("New Horizontal Value: %d", horizontal)
			log.Debugf("New depth Value: %d", newDepth)
		case "down":
			aim += val
			log.Debugf("New aim Value: %d", aim)
		case "up":
			aim -= val
			log.Debugf("New aim Value: %d", aim)
		}
		log.Debug("")
	}
	log.Debug("Calculated newDepth and horizontal value", "newDepth", newDepth, "horizontal", horizontal)

	log.Infof("New Final horizontal product :%d", newDepth*forward)
}
