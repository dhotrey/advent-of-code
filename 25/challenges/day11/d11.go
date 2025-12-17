package day11

import (
	"25/utils"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
)

func find(curr string, lookupT map[string][]string, count int, path []string) (map[string][]string, int, []string) {
	log.Debug("~>", "current", curr)

	path = append(path, curr)
	children := lookupT[curr]
	if slices.Contains(children, "out") {
		count++
		return lookupT, count, append(path, "out")
	}

	for _, child := range children {
		_, count, path = find(child, lookupT, count, path)
	}

	return lookupT, count, path
}

func Sol(mode string) {
	data, file := utils.GetInput(11, mode)
	defer file.Close()

	lookUpT := map[string][]string{}
	for data.Scan() {
		line := data.Text()
		parent := strings.Split(line, ":")
		children := strings.Split(strings.Trim(parent[1], " "), " ")
		lookUpT[parent[0]] = children
	}

	_, pathCount, pathNodes := find("you", lookUpT, 0, []string{})

	log.Info("~>", "paths", pathCount)
	log.Debug("~>", "nodes", pathNodes)

}
