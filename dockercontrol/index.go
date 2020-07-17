package dockercontrol

import (
	"context"

	u "applinh/elephant/utils"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Vol struct {
	Labels map[string]string
}

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
	} else {
		return nil, &u.ErrorString{S: "NO_STACK"}
	}
}

func StopContainer(containerID string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	err = cli.ContainerStop(context.Background(), containerID, nil)
	if err != nil {
		panic(err)
	}
	return err
}
