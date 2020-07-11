package dockercontrol

import (
	"fmt"
	"strings"

	"github.com/docker/go-connections/nat"
)

func bindPortsBuilder(ports []string) (map[nat.Port][]nat.PortBinding, []string) {

	bindings := make(map[nat.Port][]nat.PortBinding)
	errors := []string{}
	for _, port := range ports {

		portSlice := strings.Split(port, ":")

		hostBinding := nat.PortBinding{
			//HostIP:   "0.0.0.0",
			HostPort: portSlice[1],
		}

		containerPort, err := nat.NewPort("tcp", portSlice[0])
		if err != nil {
			errors = append(errors, err.Error())
		}

		if v, ok := bindings[containerPort]; !ok {
			bindings[containerPort] = []nat.PortBinding{hostBinding}
		} else {
			var n []nat.PortBinding
			copy(v, n)
			n = append(n, hostBinding)
			bindings[containerPort] = n
		}

	}
	fmt.Println(bindings)
	return bindings, errors
}
