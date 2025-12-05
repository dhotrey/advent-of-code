package day5

import (
	"25/utils"
	"sort"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func Sol(mode string) {

	// if true {
	// 	log.Fatal(mergeRanges([][2]uint64{{1, 2}, {2, 3}, {5, 6}}))
	// }
	data, file := utils.GetInput(5, mode)
	defer file.Close()

	ranges := [][2]uint64{}
	fresh := map[uint64]bool{}

	for data.Scan() {
		line := data.Text()
		if strings.Trim(line, "\n") == "" {
			continue
		}
		if strings.Contains(line, "-") {
			r := strings.Split(line, "-")
			low, _ := strconv.ParseUint(r[0], 10, 64)
			high, _ := strconv.ParseUint(r[1], 10, 64)
			ranges = append(ranges, [2]uint64{low, high})
		} else {
			num, _ := strconv.ParseUint(line, 10, 64)
			for _, r := range ranges {
				low := r[0]
				high := r[1]
				if num >= low && num <= high {
					fresh[num] = true
				}
			}
		}
	}

	log.Infof("%v ingredients are fresh", len(fresh))
	// p2
	sort.SliceStable(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	merged := mergeRanges(ranges)
	log.Debug("Merged slice", "m", merged)

	var p2 uint64
	for _, r := range merged {
		p2 += r[1] - r[0] + 1

	}
	log.Infof("Solution to p2 is %d", p2)
}

func mergeRanges(r [][2]uint64) [][2]uint64 {
	merged := [][2]uint64{}
	merges := 0

	for i := 0; i < len(r); i++ {
		if i+1 != len(r) {
			a := r[i]
			b := r[i+1]
			if b[0] <= a[1] { // merge condition
				log.Debug("Merging", "A", a, "B", b)
				lowerBound := a[0]
				upperBound := max(b[1], a[1])
				merged = append(merged, [2]uint64{lowerBound, upperBound})
				merges++
				i++
			} else {
				log.Debugf("Appending -> %v", a)
				merged = append(merged, a)
			}
		} else {
			// i.e. when i is the last index
			a := merged[len(merged)-1] // get the last element added to the merged list
			b := r[i]
			if !(b[0] >= a[0] && b[1] <= a[1]) { // when b values don't exist inside the a range
				merged = append(merged, b)
			}
		}
	}

	if merges == 0 {
		log.Debug("No merges!")
		return r
	} else {
		log.Debugf("merged %d ranges %v", merges, merged)
		return mergeRanges(merged)
	}
}
