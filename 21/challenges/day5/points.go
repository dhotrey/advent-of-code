package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

type line struct {
	lineId int
	pointA point
	pointB point
	points []point
}

func (l *line) findSlope() int {
	if l.pointB.x == l.pointA.x {
		log.Debugf("line %v is vertical", l)
		return 2
	}
	slope := (l.pointB.y - l.pointA.y) / (l.pointB.x - l.pointA.x)
	log.Debugf("Slope of line %v is %d", l, slope)
	return slope
}

func (l *line) findPointsOnLine() {
	log.Debug("Calculating Points on line")
	log.Debug(l)
	l.points = append(l.points, l.pointA)
	l.points = append(l.points, l.pointB)

	slope := l.findSlope()
	switch slope {
	case 2:
		var start, end int
		if l.pointA.y > l.pointB.y {
			start = l.pointB.y
			end = l.pointA.y
		} else if l.pointA.y < l.pointB.y {
			start = l.pointA.y
			end = l.pointB.y
		} else {
			log.Fatal("Unhandled case")
		}

		xCoord := l.pointA.x
		for yCoord := start + 1; yCoord < end; yCoord++ {
			p := point{
				x: xCoord,
				y: yCoord,
			}
			l.points = append(l.points, p)
		}
	case 0:
		var start, end int
		if l.pointA.x > l.pointB.x {
			start = l.pointB.x
			end = l.pointA.x
		} else if l.pointA.x < l.pointB.x {
			start = l.pointA.x
			end = l.pointB.x
		} else {
			log.Fatal("Unhandeled case")
		}

		yCoord := l.pointA.y
		for xCoord := start + 1; xCoord < end; xCoord++ {
			p := point{
				x: xCoord,
				y: yCoord,
			}
			l.points = append(l.points, p)
		}
	default:
		var startX, end, startY int
		if l.pointA.x < l.pointB.x {
			startX = l.pointA.x
			startY = l.pointA.y
			end = l.pointB.x
		} else {
			startX = l.pointB.x
			startY = l.pointB.y
			end = l.pointA.x
		}
		yCoord := startY
		for xCoord := startX + 1; xCoord < end; xCoord++ {
			yCoord += slope
			p := point{
				x: xCoord,
				y: yCoord,
			}
			l.points = append(l.points, p)
		}
	}
}

func (line1 *line) overlaps(line2 *line) bool {
	overlap := (line1.pointA.x == line1.pointB.x && line2.pointA.x == line2.pointB.x && line1.pointA.x == line2.pointB.x) || (line1.pointA.y == line1.pointB.y && line2.pointA.y == line2.pointB.y && line1.pointA.y == line2.pointB.y)
	if overlap {
		log.Debugf("Lines %v and %v are overlapping", line1, line2)
	}
	return overlap
}

func (l *line) String() string {
	return fmt.Sprintf("line %d from (%d,%d) to (%d,%d)", l.lineId, l.pointA.x, l.pointA.y, l.pointB.x, l.pointB.y)
}

type point struct {
	x int
	y int
}

func (p *point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func (p *point) newPoint(xy string) {
	coordinates := strings.Split(xy, ",")
	xCoord, _ := strconv.Atoi(strings.Trim(coordinates[0], " "))
	yCoord, _ := strconv.Atoi(strings.Trim(coordinates[1], " "))

	p.x = xCoord
	p.y = yCoord
}
