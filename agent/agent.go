package agent

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"code.google.com/p/go-uuid/uuid"
	"github.com/fsouza/go-dockerclient"

	"github.com/spesnova/iruka/registry"
	"github.com/spesnova/iruka/router"
	"github.com/spesnova/iruka/schema"
)

const (
	DefaultHost   = "unix:///var/run/docker.sock"
	pulseInterval = 5
)

type Agent interface {
	Pulse()
}

type IrukaAgent struct {
	docker  *docker.Client
	reg     *registry.Registry
	rou     *router.Router
	Machine string
}

func NewAgent(host, machine string, reg *registry.Registry, rou *router.Router) Agent {
	if os.Getenv("IRUKA_DOCKER_HOST") != "" {
		host = os.Getenv("IRUKA_DOCKER_HOST")
	}

	if host == "" {
		host = DefaultHost
	}

	client, err := docker.NewClient(host)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &IrukaAgent{
		docker:  client,
		reg:     reg,
		rou:     rou,
		Machine: machine,
	}
}

// Pulse retrieves container state from local docker API and registers it to etcd
func (a *IrukaAgent) Pulse() {
	for {
		containers, err := a.docker.ListContainers(docker.ListContainersOptions{All: true})
		if err != nil {
			fmt.Println(err.Error())
		}
		for _, container := range containers {
			name := strings.Replace(container.Names[0], "/", "", -1)

			// Regist only containers that managed by iruka
			s := strings.Split(name, ".")
			if uuid.Parse((s[len(s)-1])) == nil {
				fmt.Println("Skipped to register:", name)
				continue
			}

			if container.Status == "" {
				fmt.Println("Skipped to register(status is unknown):", name)
				continue
			}

			opts := schema.ContainerStateUpdateOpts{
				State:         container.Status,
				Machine:       a.Machine,
				PublishedPort: container.Ports[0].PublicPort,
			}

			c, err := a.reg.UpdateContainerState(name, opts)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			url := "http://" + c.Machine + ":" + strconv.FormatInt(c.PublishedPort, 10)

			if !a.rou.IsServerExists(c.AppID.String(), name) {
				err = a.rou.AddServer(c.AppID.String(), name, url)

				if err != nil {
					fmt.Println(err.Error())
					continue
				}
			}
		}

		time.Sleep(pulseInterval * time.Second)
	}
}
