package day11

import (
	"25/utils"
	"slices"
	"strings"

	"github.com/charmbracelet/log"
)

func find(curr string, lookupT map[string][]string, paths int) (map[string][]string, int) {
	children := lookupT[curr]
	if slices.Contains(children, "out") {
		paths++
		return lookupT, paths
	}

	for _, child := range children {
		_, paths = find(child, lookupT, paths)
	}

	return lookupT, paths
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

	_, paths := find("you", lookUpT, 0)

	log.Debug("~>", "paths", paths)

}
