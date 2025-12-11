package day8

import (
	"strconv"
	"strings"
)

type junctionPair struct {
	distance int
	p1       *point
	p2       *point
}

type circuit struct {
	boxes map[*point]bool
}

func (c1 *circuit) extend(c2 *circuit) {
	for k, v := range c2.boxes {
		c1.boxes[k] = v
	}
}

func (c *circuit) add(p *point) {
	c.boxes[p] = true
}

func newCircuit(p1, p2 *point) *circuit {
	var c circuit
	c.boxes = map[*point]bool{}
	c.add(p1)
	c.add(p2)
	return &c
}

type point struct {
	x       int
	y       int
	z       int
	circuit *circuit
}

func newPoint(l string) point {
	var p point
	xyz := strings.Split(l, ",")
	p.x, _ = strconv.Atoi(xyz[0])
	p.y, _ = strconv.Atoi(xyz[1])
	p.z, _ = strconv.Atoi(xyz[2])
	return p
}

func distance(p1, p2 *point) int {
	dx := p1.x - p2.x
	dy := p1.y - p2.y
	dz := p1.z - p2.z

	return (dx*dx + dy*dy + dz*dz)
}
