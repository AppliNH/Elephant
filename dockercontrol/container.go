package dockercontrol

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/applinh/elephant/models"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

func createContainers(cli *client.Client, services map[string]models.Service, networks map[string]string, elephantName string) (map[string]models.RunningContainer, error) {

	containers := map[string]models.RunningContainer{}

	for name, service := range services {
		portBinding, _ := bindPortsBuilder(service.Ports)
		endpointsConfig := make(map[string]*network.EndpointSettings)

		for n, id := range networks {
			endpointsConfig[n] = &network.EndpointSettings{NetworkID: id}
		}
		//fmt.Println("ServiceImage: " + service.Image)
		if _, err := cli.ImagePull(context.Background(), "docker.io/"+service.Image, types.ImagePullOptions{}); err != nil {
			return nil, err
		}

		var containerConfig = &container.Config{Image: service.Image}

		if service.Command != "" {
			containerConfig = &container.Config{
				Image: service.Image,
				Cmd:   strings.Split(service.Command, " "),
			}
		}
		platform := &v1.Platform{OS: runtime.GOOS}

		cont, err := cli.ContainerCreate(
			context.Background(), containerConfig,
			&container.HostConfig{
				PortBindings: portBinding,
			}, &network.NetworkingConfig{
				EndpointsConfig: endpointsConfig,
			}, platform, elephantName+"_"+name)

		if err != nil {
			//fmt.Println("ContainerCreate: " + err.Error())
			return nil, err
		}

		containers[cont.ID] = models.RunningContainer{ID: cont.ID, Name: elephantName + "_" + name, Elephant: elephantName}

	}

	for _, container := range containers {

		if err := startContainer(container.ID, cli); err == nil {
			fmt.Println("Sucessfully started " + container.Name + " from elephant " + container.Elephant)
		} else {
			fmt.Println(err.Error())
		}
	}
	fmt.Println(strings.Repeat("_", 25))
	return containers, nil

}

func startContainer(id string, cli *client.Client) error {
	err := cli.ContainerStart(context.Background(), id, types.ContainerStartOptions{})
	return err
}
