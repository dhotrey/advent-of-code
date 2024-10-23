package main

import (
	"21/challenges/day1"
	"21/challenges/day2"
	"21/challenges/day3"
	"21/challenges/day4"
	"flag"
	"strconv"
	"time"

	"github.com/charmbracelet/log"
)

func main() {
	log.SetReportCaller(true)
	log.SetReportTimestamp(false)
	log.SetFormatter(log.TextFormatter)

	var mode string
	verbosePtr := flag.Bool("v", false, "Show debug logs")
	flag.Parse()
	if *verbosePtr {
		log.SetLevel(log.DebugLevel)
	}

	if len(flag.Args()) < 1 {
		log.Fatalf("Need to specify challenge date, got args %v", flag.Args())
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
	}
	executionTime := time.Since(startTime)
	log.Infof("Finished executing in %s", executionTime)
}
