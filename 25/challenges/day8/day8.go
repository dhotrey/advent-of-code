package day8

import (
	"25/utils"
	"slices"

	"github.com/charmbracelet/log"
	"github.com/ernestosuarez/itertools"
)

func Sol(mode string) {

	data, file := utils.GetInput(8, mode)
	defer file.Close()

	// var iterations int
	// if mode == "test" {
	// 	iterations = 10
	// } else {
	// 	iterations = 1000
	// }

	points := itertools.List{}

	for data.Scan() {
		line := data.Text()
		p := newPoint(line)
		points = append(points, &p)
	}

	circuits := []*circuit{}
	distances := []junctionPair{}

	for combo := range itertools.CombinationsList(points, 2) {
		p1 := combo[0].(*point)
		p2 := combo[1].(*point)
		distances = append(distances, struct {
			distance int
			p1       *point
			p2       *point
		}{distance(p1, p2), p1, p2})
	}

	slices.SortFunc(distances, func(a, b junctionPair) int {
		return a.distance - b.distance
	})

	// for i := range iterations {
	// 	junctionPair := distances[i]
	// 	circuits = addJunctionPair(circuits, junctionPair)
	// }
	//
	// slices.SortFunc(circuits, func(a, b *circuit) int {
	// 	return len(b.boxes) - len(a.boxes)
	// })
	//
	// sol := circuits[:3]
	//
	// part1 := 1
	// for i, s := range sol {
	// 	log.Debugf("Size of %d element is %d", i, len(s.boxes))
	// 	part1 *= len(s.boxes)
	// }
	// log.Debug("", "total circuits created", len(circuits))
	// log.Info("Part 1 ", "solution", part1)

	part2Sol := 0
	for i := 0; ; i++ {
		junctionPair := distances[i]
		circuits = addJunctionPair(circuits, junctionPair)

		if len(circuits) == 1 {
			if len(circuits[0].boxes) == len(points) {
				part2Sol = junctionPair.p1.x * junctionPair.p2.x
				break
			}
		}
	}
	log.Info("Part 2", "solution", part2Sol)
}

func addJunctionPair(circuits []*circuit, jp junctionPair) []*circuit {
	if jp.p1.circuit != nil {
		if jp.p2.circuit != nil && jp.p2.circuit != jp.p1.circuit {
			jp.p1.circuit.extend(jp.p2.circuit)
			circuits = slices.DeleteFunc(circuits, func(target *circuit) bool {
				return target == jp.p2.circuit // remove p2.circuit
			})

			for k, _ := range jp.p2.circuit.boxes {
				k.circuit = jp.p1.circuit
			}

		} else {
			jp.p2.circuit = jp.p1.circuit
			jp.p1.circuit.add(jp.p2)
		}
	} else if jp.p2.circuit != nil {
		jp.p1.circuit = jp.p2.circuit
		jp.p2.circuit.add(jp.p1)
	} else {
		c := newCircuit(jp.p1, jp.p2)
		jp.p1.circuit = c
		jp.p2.circuit = c
		circuits = append(circuits, c)
	}
	return circuits
}
