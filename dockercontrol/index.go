package dockercontrol

import (
	"context"
	"fmt"
	"os"
	"strings"

	"time"

	u "applinh/elephant/utils"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type Vol struct {
	Labels map[string]string
}

func CreateNewContainer(path string) error {

	var dir string

	dir, _ = os.Getwd()
	path = strings.Replace(path, "/docker-compose.yml", "", -1)

	cli, err := client.NewEnvClient()

	cli.ImagePull(context.Background(), "docker/compose", types.ImagePullOptions{})

	if err != nil {
		fmt.Println("Unable to create docker client")
		panic(err)
	}

	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: "8000",
	}

	containerPort, err := nat.NewPort("tcp", "80")
	if err != nil {
		panic("Unable to get the port")
	}

	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	cont, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Entrypoint: []string{"/usr/scripts/start.sh"},
			Image:      "docker/compose",
		},
		&container.HostConfig{
			PortBindings: portBinding,
			Binds:        []string{dir + "/scripts:/usr/scripts", "/var/run/docker.sock:/var/run/docker.sock", path + ":/var/tmp/docker/compose"},
		}, nil, "")

	if err != nil {
		return err
	}

	cli.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
	fmt.Printf("Container %s has started \n", cont.ID)
	time.Sleep(5 * time.Second)

	r, e := cli.ContainerExecCreate(context.Background(), cont.ID, types.ExecConfig{
		Cmd:          []string{"/usr/scripts/getinfos.sh"},
		Tty:          true,
		AttachStderr: true,
		AttachStdout: true,
		AttachStdin:  true,
		Detach:       true,
	})

	if e != nil {
		return e
	}

	fmt.Println(r)

	return nil
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
