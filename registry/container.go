package registry

import (
	"errors"
	"path"
	"time"

	"code.google.com/p/go-uuid/uuid"

	"github.com/spesnova/iruka/schema"
)

const (
	containerPrefix = "containers"
)

// CreateContainer creates etcd key-value to store container options
func (r *Registry) CreateContainer(appIdOrName string, opts schema.ContainerCreateOpts) (schema.Container, error) {
	// Validate opts
	if opts.Image == "" {
		return schema.Container{}, errors.New("image parameter is required, but missing")
	}
	if opts.Size == "" {
		return schema.Container{}, errors.New("size parameter is required, but missing")
	}
	if opts.Command == "" {
		return schema.Container{}, errors.New("command parameter is required, but missing")
	}
	if opts.Type == "" {
		return schema.Container{}, errors.New("type parameter is required, but missing")
	}

	id := uuid.NewUUID()
	app, err := r.App(appIdOrName)
	if err != nil {
		return schema.Container{}, err
	}

	// TODO (spesnova): Add a varidation container size
	// Size takes supported value only like "1X", "2X", "3X" ...
	// But I haven't define the values yet.

	// TODO (spesnova): Add a varidation for container type

	container := schema.Container{
		ID:          id,
		AppID:       app.ID,
		Name:        app.Name + "." + opts.Type, // TODO (spesnova): add number like `examle.run.1`
		Image:       opts.Image,
		Size:        opts.Size,
		Command:     opts.Command,
		Type:        opts.Type,
		Ports:       opts.Ports,
		DesireState: "up",
		State:       "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	j, err := marshal(container)
	if err != nil {
		return schema.Container{}, err
	}

	key := path.Join(r.keyPrefix, containerPrefix, container.ID.String())

	_, err = r.etcd.Create(key, string(j), 0)
	if err != nil {
		return schema.Container{}, err
	}

	return container, nil
}

// Containers returns container collection
func (r *Registry) Containers() ([]schema.Container, error) {
	key := path.Join(r.keyPrefix, containerPrefix)
	res, err := r.etcd.Get(key, false, true)
	if err != nil {
		return nil, err
	}

	containers := make([]schema.Container, len(res.Node.Nodes))
	for i, node := range res.Node.Nodes {
		var container schema.Container
		err = unmarshal(node.Value, &container)
		if err != nil {
			return nil, err
		}

		containers[i] = container
	}

	return containers, nil
}
