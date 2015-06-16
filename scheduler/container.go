package scheduler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spesnova/go-fleet/fleet"

	"github.com/spesnova/iruka/schema"
)

// CreateContaier creates and starts a new container as systemd unit via fleet
func (s *Scheduler) CreateContainer(c schema.Container, cv schema.ConfigVars) error {
	opts, err := s.containerToUnitOptions(c, cv)
	if err != nil {
		return err
	}
	return s.fleet.Start(c.Name+".service", opts)
}

func (s *Scheduler) DestroyContainer(c schema.Container) error {
	return s.fleet.Destroy(c.Name + ".service")
}

func (s *Scheduler) Containers() ([]schema.Container, error) {
	var containers []schema.Container

	units, err := s.fleet.Units()
	if err != nil {
		return nil, err
	}

	for _, unit := range units {
		container := schema.Container{
			Name:  strings.Replace(unit.Name, ".service", "", -1),
			State: unit.CurrentState,
		}
		containers = append(containers, container)
	}

	return containers, nil
}

func (s *Scheduler) containerToUnitOptions(c schema.Container, cv schema.ConfigVars) ([]*fleet.UnitOption, error) {
	return []*fleet.UnitOption{
		// Unit section example:
		//
		//   [Unit]
		//   Description=hello.web.a888f12a-0806-11e5-b898-5cf93896cc38
		//   Requires=docker.service
		//   After=docker.service
		//
		&fleet.UnitOption{
			Section: "Unit",
			Name:    "Description",
			Value:   c.Name,
		},
		&fleet.UnitOption{
			Section: "Unit",
			Name:    "Requires",
			Value:   "docker.service",
		},
		&fleet.UnitOption{
			Section: "Unit",
			Name:    "After",
			Value:   "docker.service",
		},
		// Service section example:
		//
		//   [Service]
		//   ExecStartPre=/usr/bin/docker pull quay.io/spesnova/example:latest
		//   ExecStartPre=-/usr/bin/docker kill hello.web.a888f12a-0806-11e5-b898-5cf93896cc38
		//   ExecStartPre=-/usr/bin/docker rm hello.web.a888f12a-0806-11e5-b898-5cf93896cc38
		//   ExecStart=/usr/bin/docker run \
		//     --name hello.web.a888f12a-0806-11e5-b898-5cf93896cc38 \
		//     -p 80 \
		//     -e HELLO=world \
		//     quay.io/spesnova/example:latest \
		//     /bin/hello
		//   ExecStop=/usr/bin/docker rm -f hello.web.a888f12a-0806-11e5-b898-5cf93896cc38
		//   Restart=on-failure
		//
		&fleet.UnitOption{
			Section: "Service",
			Name:    "ExecStartPre",
			Value:   strings.Join([]string{"/usr/bin/docker", "pull", c.Image}, " "),
		},
		&fleet.UnitOption{
			Section: "Service",
			Name:    "ExecStartPre",
			Value:   strings.Join([]string{"-/usr/bin/docker", "kill", c.Name}, " "),
		},
		&fleet.UnitOption{
			Section: "Service",
			Name:    "ExecStartPre",
			Value:   strings.Join([]string{"-/usr/bin/docker", "rm", c.Name}, " "),
		},
		&fleet.UnitOption{
			Section: "Service",
			Name:    "ExecStart",
			Value: strings.Join([]string{
				"/usr/bin/docker",
				"run",
				"--name",
				c.Name,
				configVarsOptions(cv),
				portOption(c),
				sizeOption(c),
				c.Image,
				c.Command,
			}, " "),
		},
		&fleet.UnitOption{
			Section: "Service",
			Name:    "ExecStop",
			Value:   strings.Join([]string{"/usr/bin/docker", "rm -f", c.Name}, " "),
		},
		&fleet.UnitOption{
			Section: "Service",
			Name:    "Restart",
			Value:   "on-failure",
		},
	}, nil
}

func configVarsOptions(cv schema.ConfigVars) string {
	var opts []string

	for k, v := range cv {
		opts = append(opts, strings.Join([]string{"-e", k, "=", v}, ""))
	}
	return strings.Join(opts, " ")
}

func portOption(c schema.Container) string {
	//return strings.Join([]string{"-p", strconv.FormatInt(c.Port, 10)}, " ")
	return fmt.Sprintf("-p %s", strconv.FormatInt(c.Port, 10))
}

func sizeOption(c schema.Container) string {
	return ""
}
