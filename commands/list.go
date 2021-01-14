package commands

import (
	"fmt"

	"github.com/applinh/elephant/kvdb"
	"github.com/applinh/elephant/models"

	dctrl "github.com/applinh/elephant/dockercontrol"

	"github.com/boltdb/bolt"
)

// List List all running stacks with their elephants name
func List(db *bolt.DB) {

	list := make(map[string]string)
	list, _ = kvdb.ReadAll(db)
	elephantsList := models.NewElephantArmy(list)
	if len(elephantsList) > 0 {

		for _, item := range elephantsList {
			fmt.Println("🐘 " + item.Name + ":")
			for _, containerID := range item.Containers {
				cont, err := dctrl.InspectContainer(containerID)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(cont.Name + " (" + containerID + ")")
			}

		}
	} else {
		fmt.Println("No big boi found 🙅")
	}

}
