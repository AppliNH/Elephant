package dockercontrol

import (
	"fmt"

	"github.com/applinh/elephant/models"
	"github.com/applinh/elephant/utils"

	"github.com/boltdb/bolt"
	"github.com/docker/docker/client"
)

// StartStack starts the stack
func StartStack(db *bolt.DB, DC models.DockerCompose, elephantName string) (map[string]models.RunningContainer, *client.Client, error) {
	fmt.Println("Starting stack..")
	cli, _ := client.NewEnvClient()

	networks, _ := createNetworks(cli, DC.Networks)

	fmt.Println("Networks created")
	go utils.Loading()
	fmt.Println("Starting containers..")

	containers, err := createContainers(cli, DC.Services, networks, elephantName)

	return containers, cli, err
}
