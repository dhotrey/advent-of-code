package setup

import (
	"fmt"
	"log"
	"strings"
)

func SetupProject(chal string) error {
	err := envSetup()
	if err != nil {
		log.Fatal(err)
	}
	prob := strings.Split(chal,"/")
	year , err  := checkYear(prob[0])
	if err != nil {
		return fmt.Errorf("Year %s is not valid",prob[0])
	}
	if len(prob) == 1 {
		setupYear(year)
	} else {
		day , err := checkDay(prob[1])
		if err != nil {
			setupChal(day)
		}
	}

	return nil 
}

func checkYear(year string) (int , error) {
	//TODO
	return 2022 , nil 
}

func checkDay(day string) (int,error){
	//TODO
	return 1 , nil 
} 

