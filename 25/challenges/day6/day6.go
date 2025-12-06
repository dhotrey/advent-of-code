package day6

import (
	"25/utils"
	"errors"
	"slices"
	"strconv"
	"strings"

	"go/ast"
	"go/parser"
	"go/token"

	"github.com/charmbracelet/log"
)

func transformInput(input [][]string) [][]string {
	transformed := []string{}
	for j := range len(input[0][0]) {
		num := []byte{}
		for i := range len(input) - 1 {
			r := input[i][0][j]
			num = append(num, r)
		}
		transformed = append(transformed, string(num))
	}
	tempNum := []string{}
	nums := [][]string{}
	for i := range len(transformed) - 1 {
		ele := strings.Trim(transformed[i], " ")
		tempNum = append(tempNum, ele)
		if transformed[i+1] == "    " { // TODO : make this 4 inputs for final and 3 for test
			nums = append(nums, tempNum)
			tempNum = []string{}
		}
	}
	nums = append(nums, tempNum)
	nums[len(nums)-1] = append(nums[len(nums)-1], strings.Trim(transformed[len(transformed)-1], " "))
	return nums
}

func solve(operators string, inputs [][]string) int {
	var sol int
	op := strings.Fields(operators)
	for oIdx, o := range op {
		nums := inputs[oIdx]
		log.Debug(nums)
		var exp string
		for _, num := range nums {
			exp += strings.Trim(num, " ") + o
		}
		exp = strings.Trim(exp, o)
		astExp, _ := parser.ParseExpr(exp)
		res := Eval(astExp)
		log.Debugf("%s = %d", exp, res)
		sol += res
	}
	return sol
}

func Sol(mode string) {
	data, file := utils.GetInput(6, mode)
	defer file.Close()
	part1Problem := [][]string{}
	part2Problem := [][]string{}
	for data.Scan() {
		line := data.Text()
		nums := strings.Fields(line)
		nums2 := strings.Split(line, "\n")
		part1Problem = append(part1Problem, nums)
		part2Problem = append(part2Problem, nums2)
	}

	grandTotal := 0
	for i, op := range part1Problem[len(part1Problem)-1] {
		var exp string
		for _, numbers := range part1Problem[:len(part1Problem)-1] {
			exp += numbers[i] + " " + op + " "
		}
		exp = strings.TrimRight(exp, " "+op+" ")
		astExpr, err := parser.ParseExpr(exp)
		if err != nil {
			log.Fatalf("Could not parse expression %s", exp)
		}
		res := Eval(astExpr)
		grandTotal += res
		log.Debugf("%s = %d", exp, res)
	}
	log.Infof("Solution to part 1 is %d", grandTotal)
	transformedInput := transformInput(part2Problem)
	log.Debug(transformedInput)
	sol := solve(part2Problem[len(part2Problem)-1][0], transformedInput)
	log.Infof("Solution to p2 is %d", sol)
}

func Eval(exp ast.Expr) int {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(exp)
	case *ast.BasicLit:
		switch exp.Kind {
		case token.INT:
			i, _ := strconv.Atoi(exp.Value)
			return i
		}
	}
	return 0
}

func EvalBinaryExpr(exp *ast.BinaryExpr) int {
	left := Eval(exp.X)
	right := Eval(exp.Y)
	switch exp.Op {
	case token.ADD:
		return left + right
	case token.SUB:
		return left - right
	case token.MUL:
		return left * right
	case token.QUO:
		return left / right
	}
	return 0
}
