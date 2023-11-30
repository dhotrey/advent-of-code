package setup

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

// checks for a .env directory in current directory and the environment variables it holds. creates one if it does not exist.
func envSetup() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fullPath := filepath.Join(dir, ".env")
	exists, err := fileExists(fullPath)
	if err != nil {
		return err
	}
	if exists { // env file exists
		fmt.Println("env file exists")
		err := godotenv.Load()
		if err != nil {
			return err
		}
		aoc := os.Getenv("AOC")
		isAOCdir, err := strconv.ParseBool(aoc)
		if err != nil || !isAOCdir { // env file does not have the aoc environment variable or tis set to false
			fmt.Println("aoc was set to false/aoc envirn var not present")
			err = writeToEnv("AOC=true",fullPath)
			if err != nil {
				return err 
			}

		}
	} else {
		fmt.Println("create an env file ")
		f, err := os.Create(".env")
		if err != nil {
			return err
		}
		defer f.Close()
		if err != nil {
			return err
		}
		aoc := "AOC=true"
		f.WriteString(aoc)
	}
	return nil
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

// appends a line to the env file
func writeToEnv(s string, fullPath string) error { 
	f, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend) 
	if err != nil {
		return err
	}
	f.WriteString(s)
	return nil
}
