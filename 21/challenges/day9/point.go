package day9

import (
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/charmbracelet/log"
)

type point struct {
	heightMap *[][]string

	xMax int
	yMax int

	xCoord int
	yCoord int

	pointVal  int
	pointType int

	isValid       bool
	leftPossible  bool
	rightPossible bool
	upPossible    bool
	downPossible  bool
}

func (p *point) getSHA() string {
	data := fmt.Sprintf("(%d,%d)", p.xCoord, p.yCoord)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func (p *point) getUp() point {
	if p.upPossible {
		pt := newPoint(p.xCoord, p.yCoord-1, p.heightMap)
		pt.downPossible = false // to prevent infinite loop condition
		log.Debugf("got up : %s", pt)
		return pt
	}
	return point{isValid: false}
}

func (p *point) getDown() point {
	if p.downPossible {
		pt := newPoint(p.xCoord, p.yCoord+1, p.heightMap)
		pt.upPossible = false
		log.Debugf("got down: %s", pt)
		return pt
	}
	return point{isValid: false}
}

func (p *point) getLeft() point {
	if p.leftPossible {
		pt := newPoint(p.xCoord-1, p.yCoord, p.heightMap)
		pt.rightPossible = false
		log.Debugf("got left: %s", pt)
		return pt
	}
	return point{isValid: false}
}

func (p *point) getRight() point {
	if p.rightPossible {
		pt := newPoint(p.xCoord+1, p.yCoord, p.heightMap)
		pt.leftPossible = false
		log.Debugf("got right: %s", pt)
		return pt
	}
	return point{isValid: false}
}

func newPoint(xCoord, yCoord int, heightMap *[][]string) point {
	p := point{}
	p.pointType = 99 // arbitrary value

	p.heightMap = heightMap

	p.isValid = true

	p.xCoord = xCoord
	p.yCoord = yCoord

	pVal, _ := strconv.Atoi((*p.heightMap)[yCoord][xCoord])
	p.pointVal = pVal

	p.xMax = len((*p.heightMap)[0])
	p.yMax = len(*p.heightMap)

	p.leftPossible = true
	p.rightPossible = true
	p.upPossible = true
	p.downPossible = true

	if yCoord == 0 {
		p.upPossible = false
		if xCoord == 0 {
			p.pointType = topLeftCorner
			p.leftPossible = false
		} else if xCoord == p.xMax-1 {
			p.pointType = topRightCorner
			p.rightPossible = false
		} else {
			p.pointType = topEdge
		}
	} else if yCoord == p.yMax-1 {
		p.downPossible = false
		if xCoord == 0 {
			p.pointType = bottomLeftCorner
			p.leftPossible = false
		} else if xCoord == p.xMax-1 {
			p.pointType = bottomRightCorner
			p.rightPossible = false
		} else {
			p.pointType = bottomEdge
		}
	} else if xCoord == 0 {
		if p.pointType != topLeftCorner && p.pointType != bottomLeftCorner {
			p.pointType = leftEdge
			p.leftPossible = false
		}
	} else if xCoord == p.xMax-1 {
		if p.pointType != topRightCorner && p.pointType != bottomRightCorner {
			p.pointType = rightEdge
			p.rightPossible = false
		}
	} else {
		p.pointType = middleElement
	}
	return p
}

func (p point) String() string {

	return fmt.Sprintf("(***)point %d at (%d,%d) | loc : %s", p.pointVal, p.xCoord, p.yCoord, pointTypes[p.pointType])
}
