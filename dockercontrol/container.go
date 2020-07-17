package dockercontrol

import (
	. "applinh/elephant/models"
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

func createContainers(cli *client.Client, services map[string]Service, networks map[string]string, elephantName string) (map[string]RunningContainer, error) {

	containers := map[string]RunningContainer{}

	for name, service := range services {
		portBinding, _ := bindPortsBuilder(service.Ports)
		endpointsConfig := make(map[string]*network.EndpointSettings)

		for n, id := range networks {
			endpointsConfig[n] = &network.EndpointSettings{NetworkID: id}
		}
		//fmt.Println("ServiceImage: " + service.Image)
		if _, err := cli.ImagePull(context.Background(), "docker.io/"+service.Image, types.ImagePullOptions{}); err != nil {
			//fmt.Println("ImagePull: " + err.Error())
			return nil, err
		}

		var containerConfig = &container.Config{Image: service.Image}
		//fmt.Println("Speified Command: " + service.Command)

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
			//fmt.Println("ContainerCreate: " + err.Error())
			return nil, err
		} else {
			containers[cont.ID] = RunningContainer{ID: cont.ID, Name: elephantName + "_" + name, Elephant: elephantName}
			//fmt.Println(cont.ID)
		}

	}

	for _, container := range containers {
		if err := startContainer(container.ID, cli); err == nil {
			fmt.Println("Sucessfully started " + container.Name + " from elephant " + container.Elephant)
		} else {
			fmt.Println(err.Error())
		}

		//logsChan := make(chan io.ReadCloser)
	}
	fmt.Println(strings.Repeat("_", 25))
	return containers, nil

}

func startContainer(id string, cli *client.Client) error {
	err := cli.ContainerStart(context.Background(), id, types.ContainerStartOptions{})
	return err

	// ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer cancel()

	//ReadLogs(cli, id)

}
