package agent

import (
	"fmt"
	"os"
	"strings"
	"time"

	"code.google.com/p/go-uuid/uuid"
	"github.com/fsouza/go-dockerclient"

	"github.com/spesnova/iruka/registry"
)

const (
	DefaultHost   = "unix:///var/run/docker.sock"
	pulseInterval = 5
)

type Agent interface {
	Pulse()
}

type IrukaAgent struct {
	docker *docker.Client
	reg    *registry.Registry
}

func NewAgent(host string, reg *registry.Registry) Agent {
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
		client,
		reg,
	}
}

// Pulse retrieves container state from local docker API and registers
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
				fmt.Println(name)
				continue
			}

			if container.Status == "" {
				fmt.Println(name)
				continue
			}

			_, err := a.reg.UpdateContainerState(name, container.Status)
			if err != nil {
				fmt.Println(err.Error())
			}
			time.Sleep(pulseInterval * time.Second)
		}
	}
}
