package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {

	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)

	defer cli.Close()

	if err != nil {
		panic(err)
	}

	cli.NegotiateAPIVersion(ctx)

	serviceList, _ := cli.ServiceList(ctx, types.ServiceListOptions{})

	fmt.Println("SID, ServiceName, RAM, CPU, Team, Networks, Constraints, Repo, ImageName, Healthcheck")

	for _, service := range serviceList {

		var sid string
		var name string
		var team string
		var MemoryBytes int64
		var NanoCPUs int64
		var Networks []string
		var Constraints []string
		var Repo string
		var Image string
		var Healthcheck string

		for _, str := range service.Spec.TaskTemplate.ContainerSpec.Env {

			if strings.Contains(str, "team=") {

				team = str

			}

			if strings.Contains(str, "sid=") {

				sid = str

			}

		}

		name = service.Spec.Name

		Constraints = service.Spec.TaskTemplate.Placement.Constraints

		if service.Spec.TaskTemplate.Resources.Limits != nil {

			MemoryBytes = service.Spec.TaskTemplate.Resources.Limits.MemoryBytes
			NanoCPUs = service.Spec.TaskTemplate.Resources.Limits.NanoCPUs

		} else {

			MemoryBytes = 0
			NanoCPUs = 0

		}

		for _, net := range service.Spec.TaskTemplate.Networks {

			networkName, _ := cli.NetworkInspect(ctx, net.Target, types.NetworkInspectOptions{})

			Networks = append(Networks, networkName.Name)

		}

		Repo = (strings.SplitN(service.Spec.TaskTemplate.ContainerSpec.Image, "/", 2))[0]
		Image = service.Spec.TaskTemplate.ContainerSpec.Image

		Healthcheck = "False"

		if service.Spec.TaskTemplate.ContainerSpec.Healthcheck != nil {

			Healthcheck = "True"

		} else {

			Healthcheck = "False"

		}

		fmt.Println(sid, ",", name, ",", MemoryBytes, ",", NanoCPUs, ",", team, ",", Networks, ",", Constraints, ",", Repo, ",", Image, ",", Healthcheck)

	}

}
