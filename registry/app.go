package registry

import (
	"errors"
	"path"
	"time"

	"code.google.com/p/go-uuid/uuid"

	"github.com/spesnova/iruka/schema"
)

const (
	appPrefix = "apps"
	domain    = "irukaapp.com"
)

// CreateApp creates etcd key-value to store meta data for an app
func (r *Registry) CreateApp(opts schema.AppCreateOpts) (schema.App, error) {
	// Validate opts
	if opts.Name == "" {
		return schema.App{}, errors.New("name parameter is required, but missing")
	}

	id := uuid.NewUUID()
	webURL := id.String() + "." + domain

	app := schema.App{
		ID:        id,
		Name:      opts.Name,
		WebURL:    webURL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	j, err := marshal(app)
	if err != nil {
		return schema.App{}, err
	}

	// Key: /iruka/apps/<UUID>
	key := path.Join(r.keyPrefix, appPrefix, app.ID.String())

	// TODO (spesnova): validate app name is unique before create a new key
	_, err = r.etcd.Create(key, string(j), 0)
	if err != nil {
		return schema.App{}, err
	}

	return app, nil
}

// DestroyApp removes an etcd key-value
func (r *Registry) DestroyApp(idOrName string) (schema.App, error) {
	app, err := r.App(idOrName)
	if err != nil {
		return schema.App{}, err
	}

	key := path.Join(r.keyPrefix, appPrefix, app.ID.String())

	_, err = r.etcd.Delete(key, true)
	if err != nil {
		return schema.App{}, errors.New("Failed to delete app: " + app.ID.String())
	}

	return app, err
}

// App returns an app, you can find it by id or name
func (r *Registry) App(idOrName string) (schema.App, error) {
	var app schema.App

	apps, err := r.Apps()
	if err != nil {
		return app, err
	}

	if uuid.Parse(idOrName) == nil {
		for _, app := range apps {
			if app.Name == idOrName {
				return app, nil
			}
		}
	} else {
		for _, app := range apps {
			if uuid.Equal(app.ID, uuid.Parse(idOrName)) {
				return app, nil
			}
		}
	}

	return app, errors.New("No such app: " + idOrName)
}

func (r *Registry) Apps() ([]schema.App, error) {
	key := path.Join(r.keyPrefix, appPrefix)
	res, err := r.etcd.Get(key, false, true)
	if err != nil {
		return nil, err
	}

	apps := make([]schema.App, len(res.Node.Nodes))
	for i, node := range res.Node.Nodes {
		var app schema.App
		err = unmarshal(node.Value, &app)
		if err != nil {
			return nil, err
		}

		apps[i] = app
	}

	return apps, nil
}

// UpdateApp updates an app
func (r *Registry) UpdateApp(idOrName string, opts schema.AppUpdateOpts) (schema.App, error) {
	// Validate opts
	if opts.Name == "" {
		return schema.App{}, errors.New("name parameter is required, but missing")
	}

	app, err := r.App(idOrName)
	if err != nil {
		return schema.App{}, err
	}
	app.UpdatedAt = time.Now()
	app.Name = opts.Name

	j, err := marshal(app)
	if err != nil {
		return schema.App{}, err
	}

	key := path.Join(r.keyPrefix, appPrefix, app.ID.String())
	res, err := r.etcd.Set(key, string(j), 0)
	if err != nil {
		return schema.App{}, err
	}

	err = unmarshal(res.Node.Value, &app)
	if err != nil {
		return schema.App{}, err
	}

	return app, nil
}
