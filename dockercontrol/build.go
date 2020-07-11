package dockercontrol

import (
	. "applinh/elephant/models"

	"github.com/docker/docker/client"
)

func StartStack(DC DockerCompose) {

	cli, _ := client.NewEnvClient()

	networks, _ := createNetworks(cli, DC.Networks)

	createContainers(cli, DC.Services, networks)

}
