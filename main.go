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

	if err != nil {
		panic(err)
	}

	cli.NegotiateAPIVersion(ctx)

	serviceList, _ := cli.ServiceList(ctx, types.ServiceListOptions{})

	fmt.Println("sid, serviceName, CPULimit&RAMLimit, Team")

	for _, service := range serviceList {

		team := "team=null"
		sid := "null"

		for _, str := range service.Spec.TaskTemplate.ContainerSpec.Env {

			if strings.Contains(str, "team=") {

				team = str

			}

			if strings.Contains(str, "sid=") {

				sid = str

			}

		}

		fmt.Println(sid, ",", service.Spec.Name, ",", service.Spec.TaskTemplate.Resources.Limits, ",", team)

	}

}
