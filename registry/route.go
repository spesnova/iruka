package registry

import (
	"errors"
	"path"
	"sort"
	"time"

	"code.google.com/p/go-uuid/uuid"

	"github.com/spesnova/iruka/schema"
)

const (
	routePrefix = "routes"
)

type Routes []schema.Route

// imprement the sort interface
func (r Routes) Len() int {
	return len(r)
}

func (r Routes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Routes) Less(i, j int) bool {
	return r[i].ID.String() < r[j].ID.String()
}

func (r *Registry) CreateRoute(appIdentity string, opts schema.RouteCreateOpts) (schema.Route, error) {
	// TODO(dtan4):
	//   etcd does not reflect inserted item immediately...
	// app, err := r.App(appIdentity)

	// if err != nil {
	// 	return schema.Route{}, err
	// }

	if opts.Location == "" {
		return schema.Route{}, errors.New("location parameter is required, but missing")
	}

	if opts.Upstream == "" {
		return schema.Route{}, errors.New("upstream parameter is required, but missing")
	}

	id := uuid.NewUUID()
	currentTime := time.Now()
	route := schema.Route{
		ID:        id,
		Location:  opts.Location,
		Upstream:  opts.Upstream,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	j, err := marshal(route)

	if err != nil {
		return schema.Route{}, err
	}

	// key := path.Join(r.keyPrefix, routePrefix, app.ID.String(), route.ID.String())
	key := path.Join(r.keyPrefix, routePrefix, appIdentity, route.ID.String())

	if _, err = r.etcd.Create(key, string(j), 0); err != nil {
		return schema.Route{}, err
	}

	return route, nil
}

func (r *Registry) DestroyRoute(appIdentity, routeID string) (schema.Route, error) {
	app, err := r.App(appIdentity)

	if err != nil {
		return schema.Route{}, err
	}

	route, err := r.Route(appIdentity, routeID)

	if err != nil {
		return schema.Route{}, err
	}

	key := path.Join(r.keyPrefix, routePrefix, app.ID.String(), routeID)
	_, err = r.etcd.Delete(key, true)

	if err != nil {
		if isKeyNotFound(err) {
			err = nil
		}
		return schema.Route{}, err
	}

	return route, nil
}

func (r *Registry) Route(appIdentity, routeID string) (schema.Route, error) {
	app, err := r.App(appIdentity)

	if err != nil {
		return schema.Route{}, err
	}

	key := path.Join(r.keyPrefix, routePrefix, app.ID.String(), routeID)
	res, err := r.etcd.Get(key, false, true)

	if err != nil {
		if isKeyNotFound(err) {
			err = nil
		}

		return schema.Route{}, err
	}

	var route schema.Route
	err = unmarshal(res.Node.Value, &route)

	if err != nil {
		return schema.Route{}, err
	}

	return route, nil
}

func (r *Registry) Routes(appIdentity string) ([]schema.Route, error) {
	app, err := r.App(appIdentity)

	if err != nil {
		return nil, err
	}

	key := path.Join(r.keyPrefix, routePrefix, app.ID.String())
	res, err := r.etcd.Get(key, false, true)

	if err != nil {
		if isKeyNotFound(err) {
			err = nil
		}
		return nil, err
	}

	if len(res.Node.Nodes) == 0 {
		return nil, nil
	}

	var routes Routes

	for _, node := range res.Node.Nodes {
		var route schema.Route
		err = unmarshal(node.Value, &route)

		if err != nil {
			return nil, err
		}

		routes = append(routes, route)
	}

	sort.Sort(sort.Reverse(routes))

	return routes, nil
}

func (r *Registry) UpdateRoute(appIdentity string, routeID string, opts schema.RouteUpdateOpts) (schema.Route, error) {
	app, err := r.App(appIdentity)

	if err != nil {
		return schema.Route{}, err
	}

	route, err := r.Route(appIdentity, routeID)

	if err != nil {
		return schema.Route{}, err
	}

	if opts.ID.String() == "" {
		return schema.Route{}, errors.New("id parameter is required, but missing")
	}

	if opts.Location != "" {
		route.Location = opts.Location
	}

	if opts.Upstream != "" {
		route.Upstream = opts.Upstream
	}

	j, err := marshal(route)

	if err != nil {
		return schema.Route{}, err
	}

	key := path.Join(r.keyPrefix, routePrefix, app.ID.String(), routeID)

	if _, err := r.etcd.Set(key, string(j), 0); err != nil {
		return schema.Route{}, err
	}

	return route, nil
}
