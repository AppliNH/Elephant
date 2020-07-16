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

	containers := map[string]RunningContainer{}
	fmt.Println(services)

	for name, service := range services {
		portBinding, _ := bindPortsBuilder(service.Ports)
		endpointsConfig := make(map[string]*network.EndpointSettings)

		for n, id := range networks {
			endpointsConfig[n] = &network.EndpointSettings{NetworkID: id}
		}
		fmt.Println("ServiceImage: " + service.Image)
		if _, err := cli.ImagePull(context.Background(), service.Image, types.ImagePullOptions{}); err != nil {
			fmt.Println("ImagePull: " + err.Error())
		}

		var containerConfig = &container.Config{Image: service.Image}
		fmt.Println("Speified Command: " + service.Command)

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
			fmt.Println("ContainerCreate: " + err.Error())
		} else {
			containers[cont.ID] = RunningContainer{ID: cont.ID, Name: elephantName + "_" + name, Elephant: elephantName}
			fmt.Println(cont.ID)
		}

	}
	fmt.Println(strings.Repeat("_", 25))
	for _, container := range containers {
		startContainer(container.ID, cli)
		//logsChan := make(chan io.ReadCloser)
	}

	ReadLogs(cli, containers)

}

func startContainer(id string, cli *client.Client) {
	if err := cli.ContainerStart(context.Background(), id, types.ContainerStartOptions{}); err != nil {
		fmt.Println(err)
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer cancel()

	//ReadLogs(cli, id)

}
