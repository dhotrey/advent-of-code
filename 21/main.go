package main

import (
	"21/challenges/day1"
	"21/challenges/day2"
	"flag"
	"strconv"

	"github.com/charmbracelet/log"
)

func main() {
	log.SetReportCaller(true)
	log.SetReportTimestamp(false)
	// log.SetTimeFormat("2006-01-02 15:04:02")
	// log.SetPrefix("AOC-2021")
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

	switch chal {
	case 1:
		day1.Sol(mode)
	case 2:
		day2.Sol(mode)
	}

}
