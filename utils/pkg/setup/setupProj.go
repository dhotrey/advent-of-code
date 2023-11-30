package setup

import (
	"log"
)

func SetupProject(args string) {
	err := envSetup()
	if err != nil {
		log.Fatal(err)
	}

}
