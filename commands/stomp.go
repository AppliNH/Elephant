package commands

import (
	"fmt"
	"strings"

	"github.com/applinh/elephant/models"

	dctrl "github.com/applinh/elephant/dockercontrol"
	"github.com/applinh/elephant/kvdb"

	"github.com/boltdb/bolt"
)

// Stomp stops a stack by providing an elephant name, and update in db
func Stomp(db *bolt.DB, elephantName string) {
	containersString, err := kvdb.ReadData(db, elephantName)
	if err != nil {
		fmt.Println(err)
		return
	}

	elephant := models.Elephant{Name: elephantName, Containers: strings.Split(containersString, ",")}

	if len(elephant.Containers) > 0 {
		fmt.Println("This elephant carries: ")

		for _, containerID := range elephant.Containers {
			cont, err := dctrl.InspectContainer(containerID)
			if err != nil {
				fmt.Println("Couldn't get info for container " + containerID)
				break
			}
			fmt.Println("- " + cont.Name + "(" + containerID + ")")
		}

		for _, containerID := range elephant.Containers {

			if err := dctrl.StopContainer(containerID); err == nil {
				fmt.Println("Successfully stopped container " + containerID)
				elephant.RemContainerByID(containerID)
			} else {
				fmt.Println("Couldn't stop container " + containerID)
			}
		}

		if len(elephant.Containers) == 0 {
			elephant.DelFromDB(db)
		} else {
			elephant.WriteToDB(db)
		}

	} else {
		fmt.Println("The elephant named " + elephantName + " doesn't seem to carry any container")
	}

}
