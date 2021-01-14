package dockercontrol

import (
	"context"

	"github.com/applinh/elephant/models"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func createNetworks(cli *client.Client, networkNames map[string]models.Network) (map[string]string, map[string]string) {
	networks := make(map[string]string)
	errors := make(map[string]string)

	for name := range networkNames {
		if r, e := cli.NetworkCreate(context.Background(), name, types.NetworkCreate{}); e != nil {
			errors[name] = e.Error()
		} else {
			networks[name] = r.ID
		}

	}
	return networks, errors
}
