package dockercontrol

import (
	. "applinh/elephant/dockermng"
	. "applinh/elephant/models"
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

func createContainers(cli *client.Client, services map[string]Service, networks map[string]string, elephantName string) {

	for name, service := range services {

		portBinding, _ := bindPortsBuilder(service.Ports)
		endpointsConfig := make(map[string]*network.EndpointSettings)

		for n, id := range networks {
			endpointsConfig[n] = &network.EndpointSettings{NetworkID: id}
		}

		var containerConfig = &container.Config{Image: service.Image}

		if service.Command != "" {
			containerConfig = &container.Config{
				Image: service.Image,
				Cmd:   strings.Split(service.Command, " "),
			}
		}
		cont, err := cli.ContainerCreate(
			context.Background(), containerConfig,
			&container.HostConfig{
				PortBindings: portBinding,
			}, &network.NetworkingConfig{
				EndpointsConfig: endpointsConfig,
			}, elephantName+"_"+name)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(cont.ID)
		}
		cli.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
		//logsChan := make(chan io.ReadCloser)
		ReadLogs(cli, cont.ID)

	}

}
