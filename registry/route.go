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
	app, err := r.App(appIdentity)

	if err != nil {
		return schema.Route{}, err
	}

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
		AppID:     app.ID,
		Location:  opts.Location,
		Upstream:  opts.Upstream,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	j, err := marshal(route)

	if err != nil {
		return schema.Route{}, err
	}

	key := path.Join(r.keyPrefix, routePrefix, route.ID.String())

	if _, err = r.etcd.Create(key, string(j), 0); err != nil {
		return schema.Route{}, err
	}

	return route, nil
}

func (r *Registry) DestroyRoute(identity string) (schema.Route, error) {
	route, err := r.Route(identity)

	if err != nil {
		return schema.Route{}, err
	}

	key := path.Join(r.keyPrefix, routePrefix, route.ID.String())
	_, err = r.etcd.Delete(key, true)

	if err != nil {
		return schema.Route{}, errors.New("Failed to delete route: " + route.ID.String())
	}

	return route, nil
}

func (r *Registry) Route(identity string) (schema.Route, error) {
	var route schema.Route

	routes, err := r.Routes()

	if err != nil {
		return route, err
	}

	if uuid.Parse(identity) == nil {
		return route, errors.New("No such route: " + identity)
	} else {
		for _, route := range routes {
			if uuid.Equal(route.ID, uuid.Parse(identity)) {
				return route, nil
			}
		}
	}

	return route, errors.New("No such route: " + identity)
}

func (r *Registry) RouteFilteredByApp(appIdentity, identity string) (schema.Route, error) {
	route, err := r.Route(identity)

	if err != nil {
		return schema.Route{}, err
	}

	app, err := r.App(appIdentity)

	if err != nil {
		return schema.Route{}, err
	}

	if uuid.Equal(route.AppID, app.ID) {
		return route, nil
	}

	return route, errors.New("No such route: " + identity)
}

func (r *Registry) Routes() ([]schema.Route, error) {
	key := path.Join(r.keyPrefix, routePrefix)
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

func (r *Registry) RoutesFilteredByApp(appIdentity string) ([]schema.Route, error) {
	var routes []schema.Route

	app, err := r.App(appIdentity)

	if err != nil {
		return nil, err
	}

	rs, err := r.Routes()

	if err != nil {
		return nil, err
	}

	if rs == nil {
		return nil, nil
	}

	for _, r := range rs {
		if uuid.Equal(r.AppID, app.ID) {
			routes = append(routes, r)
		}
	}

	return routes, nil
}

func (r *Registry) UpdateRoute(appIdentity, identity string, opts schema.RouteUpdateOpts) (schema.Route, error) {
	route, err := r.RouteFilteredByApp(appIdentity, identity)

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

	key := path.Join(r.keyPrefix, routePrefix, identity)

	if _, err := r.etcd.Set(key, string(j), 0); err != nil {
		return schema.Route{}, err
	}

	return route, nil
}
