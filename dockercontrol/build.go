package dockercontrol

import (
	. "applinh/elephant/dockermng"
	. "applinh/elephant/kvdb"
	. "applinh/elephant/models"
	"fmt"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/docker/docker/client"
)

func StartStack(db *bolt.DB, DC DockerCompose, elephantName string) {
	fmt.Println("Starting stack..")
	cli, _ := client.NewEnvClient()

	networks, _ := createNetworks(cli, DC.Networks)
	fmt.Println("Networks created")
	fmt.Println("Starting containers..")
	containers, err := createContainers(cli, DC.Services, networks, elephantName)
	if err == nil {
		contIDs := []string{}

		for k, _ := range containers {
			contIDs = append(contIDs, k)
		}
		fmt.Println("All good ! Containers have been packed on your big boi 🐘")
		WriteData(db, elephantName, strings.Join(contIDs, ","))
		//fmt.Println(ReadData(db, elephantName))
		//db.Close()
		fmt.Println(strings.Repeat("_", 25))
		fmt.Println("Logs")
		fmt.Println(strings.Repeat("_", 25))
		fmt.Println()
		ReadLogs(cli, containers)
	} else {
		fmt.Println("An error has occured")
		fmt.Println(err.Error())
		return
	}

}
