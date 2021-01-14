package dockercontrol

import (
	"context"
	"fmt"

	"github.com/applinh/elephant/utils"
	u "github.com/applinh/elephant/utils"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// Vol struct defines a docker-container volume
type Vol struct {
	Labels map[string]string
}

// ListContainers list running containers
func ListContainers() ([]types.Container, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	if len(containers) > 0 {
		return containers, nil
	}

	return nil, &u.ErrorString{S: "NO_STACK"}

}

// StopContainer stop a container
func StopContainer(containerID string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	fmt.Println("Stopping " + containerID + " ...")
	go utils.Loading()
	return cli.ContainerStop(context.Background(), containerID, nil)
}

// InspectContainer retrieves container informations
func InspectContainer(containerID string) (types.ContainerJSON, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}
	return cli.ContainerInspect(context.Background(), containerID)
}
