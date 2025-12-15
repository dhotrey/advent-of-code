package main

import (
	"25/challenges/day1"
	"25/challenges/day11"
	"25/challenges/day2"
	"25/challenges/day3"
	"25/challenges/day4"
	"25/challenges/day5"
	"25/challenges/day6"
	"25/challenges/day7"
	"25/challenges/day8"
	"25/challenges/day9"
	"flag"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
)

func main() {
	log.SetReportCaller(false)
	log.SetReportTimestamp(false)
	log.SetFormatter(log.TextFormatter)

	var mode string
	verbosePtr := flag.Bool("v", false, "Show debug logs")
	flag.Parse()
	if *verbosePtr {
		log.SetLevel(log.DebugLevel)
	}

	if len(flag.Args()) < 1 {
		log.Fatalf("Need to specify challenge number, got args %v", flag.Args())
	}
	chal, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		log.Fatal("Challenge date must be an int", "invalid challenge", chal, "error", err)
	}

	if len(flag.Args()) != 2 {
		mode = "test"
	} else {
		mode = "final"
	}

	startTime := time.Now()
	switch chal {
	case 1:
		day1.Sol(mode)
	case 2:
		day2.Sol(mode)
	case 3:
		day3.Sol(mode)
	case 4:
		day4.Sol(mode)
	case 5:
		day5.Sol(mode)
	case 6:
		day6.Sol(mode)
	case 7:
		day7.Sol(mode)
	case 8:
		day8.Sol(mode)
	case 9:
		day9.Sol(mode)
	case 11:
		day11.Sol(mode)
	}
	executionTime := time.Since(startTime)
	log.Infof("Solved challenge in %v seconds", executionTime.Seconds())
}
