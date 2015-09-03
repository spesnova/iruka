package router

import (
	"os"
	"path"
	"strings"

	"github.com/coreos/go-etcd/etcd"
)

const (
	DefaultKeyPrefix = "/vulcand"
	DefaultMachines  = "http://localhost:4001"
)

type Router struct {
	etcd      etcd.Client
	keyPrefix string
}

func NewRouter(machines, keyPrefix string) *Router {
	if os.Getenv("IRUKA_ETCD_MACHINES") != "" {
		machines = os.Getenv("IRUKA_ETCD_MACHINES")
	}

	if machines == "" {
		machines = DefaultMachines
	}

	m := strings.Split(machines, ",")
	etcdClient := *etcd.NewClient(m)
	return &Router{etcdClient, keyPrefix}
}

func (r *Router) AddRoute(host, location, locPath, upstream string) error {
	pKey := path.Join(r.keyPrefix, "hosts", host, "locations", location, "path")
	_, err := r.etcd.Create(pKey, locPath, 0)

	if err != nil {
		return err
	}

	uKey := path.Join(r.keyPrefix, "hosts", host, "locations", location, "upstream")
	_, err = r.etcd.Create(uKey, upstream, 0)

	if err != nil {
		return err
	}

	return nil
}

func (r *Router) RemoveRoute(host, location string) error {
	key := path.Join(r.keyPrefix, "hosts", host, "locations", location)
	_, err := r.etcd.Delete(key, true)

	if err != nil {
		return err
	}

	return nil
}

func (r *Router) UpdateRoute(host, location, locPath, upstream string) error {
	if locPath != "" {
		pKey := path.Join(r.keyPrefix, "hosts", host, "locations", location, "path")
		_, err := r.etcd.Set(pKey, locPath, 0)
		if err != nil {
			return err
		}
	}

	if upstream != "" {
		uKey := path.Join(r.keyPrefix, "hosts", host, "locations", location, "upstream")
		_, err := r.etcd.Set(uKey, upstream, 0)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Router) RemoveHost(host string) error {
	key := path.Join(r.keyPrefix, "hosts", host)
	_, err := r.etcd.Delete(key, true)

	if err != nil {
		return err
	}

	return nil
}

func (r *Router) AddUpstream(upstream string) error {
	key := path.Join(r.keyPrefix, "upstreams", upstream)
	_, err := r.etcd.SetDir(key, 0)

	if err != nil {
		return err
	}

	return nil
}

func (r *Router) RemoveUpstream(upstream string) error {
	key := path.Join(r.keyPrefix, "upstreams", upstream)
	_, err := r.etcd.Delete(key, true)

	if err != nil {
		return err
	}

	return nil
}
