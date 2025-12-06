package day6

import (
	"25/utils"
	"strconv"
	"strings"

	"go/ast"
	"go/parser"
	"go/token"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {
	data, file := utils.GetInput(6, mode)
	defer file.Close()

	problem := [][]string{}
	for data.Scan() {
		line := data.Text()
		nums := strings.Fields(line)
		problem = append(problem, nums)
	}

	grandTotal := 0

	for i, op := range problem[len(problem)-1] {
		var exp string
		for _, numbers := range problem[:len(problem)-1] {
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
