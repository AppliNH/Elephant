package commands

import (
	. "applinh/elephant/dockercontrol"
	. "applinh/elephant/kvdb"
	. "applinh/elephant/utils"
	"fmt"
	"regexp"
	"strings"

	"github.com/boltdb/bolt"
)

func Stomp(db *bolt.DB, elephantName string) {
	containersString := ReadData(db, elephantName)

	var containNumber = regexp.MustCompile(`\d`)
	if containNumber.MatchString(containersString) {

		fmt.Println("This elephant carries: " + containersString)
		containerIDs := strings.Split(containersString, ",")
		fmt.Println("Our big boi will drop and stomp on these.")
		fmt.Println("Please wait..")
		go Loading()
		for _, v := range containerIDs {
			if err := StopContainer(v); err == nil {
				fmt.Println("Successfully stopped container " + v)
				currentVALS := ReadData(db, elephantName)

				newVALS := strings.ReplaceAll(currentVALS, v, "")
				WriteData(db, elephantName, newVALS)
			} else {
				fmt.Println("Couldn't stop container " + v)
			}

		}
	} else {
		fmt.Println("The elephant named " + elephantName + " doesn't seem to carry any container")
	}
	return

}
