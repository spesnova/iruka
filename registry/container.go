package registry

import (
	"errors"
	"path"
	"strings"
	"time"

	"code.google.com/p/go-uuid/uuid"

	"github.com/spesnova/iruka/schema"
)

const (
	containerPrefix = "containers"
)

type Containers []schema.Container

// CreateContainer creates etcd key-value to store container options
func (r *Registry) CreateContainer(appIdOrName string, opts schema.ContainerCreateOpts) (schema.Container, error) {
	app, err := r.App(appIdOrName)
	if err != nil {
		return schema.Container{}, err
	}

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

	// TODO (spesnova): Add a varidation container size
	// Size takes supported value only like "1X", "2X", "3X" ...
	// But I haven't define the values yet.

	// TODO (spesnova): Add a varidation for container type

	container := schema.Container{
		ID:          id,
		AppID:       app.ID,
		Name:        strings.Join([]string{app.Name, opts.Type, id.String()}, "."),
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

func (r *Registry) DestroyContainer(idOrName string) (schema.Container, error) {
	container, err := r.Container(idOrName)
	if err != nil {
		return schema.Container{}, err
	}

	key := path.Join(r.keyPrefix, containerPrefix, container.ID.String())

	_, err = r.etcd.Delete(key, true)
	if err != nil {
		return schema.Container{}, errors.New("Failed to delete container: " + container.ID.String())
	}

	return container, err
}

// Container returns a container
func (r *Registry) Container(idOrName string) (schema.Container, error) {
	var container schema.Container

	cs, err := r.Containers()
	if err != nil {
		return schema.Container{}, err
	}

	if uuid.Parse(idOrName) == nil {
		for _, c := range cs {
			if c.Name == idOrName {
				return c, nil
			}
		}
	} else {
		for _, c := range cs {
			if uuid.Equal(c.ID, uuid.Parse(idOrName)) {
				return c, nil
			}
		}
	}

	return container, errors.New("No such container: " + idOrName)
}

// Container returns a container of an app
func (r *Registry) ContainerFilteredByApp(appIdOrName, idOrName string) (schema.Container, error) {
	container, err := r.Container(idOrName)
	if err != nil {
		return schema.Container{}, err
	}

	app, err := r.App(appIdOrName)
	if err != nil {
		return schema.Container{}, err
	}

	if uuid.Equal(container.AppID, app.ID) {
		return container, nil
	}

	return container, errors.New("No such container: " + idOrName)
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

// ContainersFilteredByApp returns containers of an app
func (r *Registry) ContainersFilteredByApp(appIdOrName string) ([]schema.Container, error) {
	var containers []schema.Container

	app, err := r.App(appIdOrName)
	if err != nil {
		return nil, err
	}

	cs, err := r.Containers()
	if err != nil {
		return nil, err
	}

	for _, c := range cs {
		if uuid.Equal(c.AppID, app.ID) {
			containers = append(containers, c)
		}
	}

	return containers, nil
}

// UpdateContainer updates an options of a contaienr
func (r *Registry) UpdateContainer(idOrName string, opts schema.ContainerUpdateOpts) (schema.Container, error) {
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

	container, err := r.Container(idOrName)
	if err != nil {
		return schema.Container{}, err
	}
	container.Image = opts.Image
	container.Size = opts.Size
	container.Command = opts.Command
	container.Type = opts.Type
	container.UpdatedAt = time.Now()

	j, err := marshal(container)
	if err != nil {
		return schema.Container{}, err
	}

	key := path.Join(r.keyPrefix, containerPrefix, container.ID.String())
	res, err := r.etcd.Set(key, string(j), 0)
	if err != nil {
		return schema.Container{}, err
	}

	err = unmarshal(res.Node.Value, &container)
	if err != nil {
		return schema.Container{}, err
	}

	return container, nil
}
