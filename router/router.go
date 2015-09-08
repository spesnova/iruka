package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/coreos/go-etcd/etcd"

	"github.com/spesnova/iruka/schema"
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

func (r *Router) AddBackend(name string) error {
	backend := schema.VulcandBackend{
		Type: "http",
	}
	j, err := marshal(backend)

	if err != nil {
		return err
	}

	key := path.Join(r.keyPrefix, "backends", name, "backend")
	_, err = r.etcd.Create(key, string(j), 0)

	if err != nil {
		return err
	}

	return nil
}

func (r *Router) RemoveBackend(name string) error {
	key := path.Join(r.keyPrefix, "backends", name, "backend")
	_, err := r.etcd.Delete(key, true)

	if err != nil {
		return err
	}

	return nil
}

func (r *Router) AddServer(appID, containerName, url string) error {
	server := schema.VulcandServer{
		URL: url,
	}
	j, err := marshal(server)

	if err != nil {
		return err
	}

	key := path.Join(r.keyPrefix, "backends", appID, "servers", containerName)
	_, err = r.etcd.Create(key, string(j), 0)

	if err != nil {
		return err
	}

	return nil
}

func (r *Router) RemoveServer(appID, containerName string) error {
	key := path.Join(r.keyPrefix, "backends", appID, "servers", containerName)
	_, err := r.etcd.Delete(key, true)

	if err != nil {
		return err
	}

	return nil
}

func (r *Router) IsServerExists(appID, containerName string) bool {
	key := path.Join(r.keyPrefix, "backends", appID, "servers", containerName)
	_, err := r.etcd.Get(key, false, false)

	return err == nil
}

func (r *Router) AddRoute(name, host, location string) error {
	frontend := schema.VulcandFrontend{
		Type:      "http",
		BackendID: name,
		Route:     routeString(host, location),
	}
	j, err := marshal(frontend)

	if err != nil {
		return err
	}

	key := path.Join(r.keyPrefix, "frontends", fmt.Sprintf("%s-%s", name, host), "frontend")
	_, err = r.etcd.Create(key, j, 0)

	if err != nil {
		return err
	}

	return nil
}

func (r *Router) RemoveRoute(name, host string) error {
	key := path.Join(r.keyPrefix, "frontends", fmt.Sprintf("%s-%s", name, host), "frontend")
	_, err := r.etcd.Delete(key, true)

	if err != nil {
		return err
	}

	return nil
}

func (r *Router) UpdateRoute(name, host, location string) error {
	frontend := schema.VulcandFrontend{
		Type:      "http",
		BackendID: name,
		Route:     routeString(host, location),
	}
	j, err := marshal(frontend)

	if err != nil {
		return err
	}

	key := path.Join(r.keyPrefix, "frontends", name, "frontend")
	_, err = r.etcd.Set(key, j, 0)

	if err != nil {
		return err
	}

	return nil
}

func routeString(hostname, path string) string {
	return fmt.Sprintf("Host(`%s`) && PathRegexp(`%s`)", hostname, path)
}

func marshal(obj interface{}) (string, error) {
	encoded, err := json.Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("unable to JSON-serialize object: %s", err)
	}

	// To print '&'
	return string(bytes.Replace(encoded, []byte("\\u0026"), []byte("&"), -1)), nil
}

func unmarshal(val string, obj interface{}) error {
	err := json.Unmarshal([]byte(val), &obj)
	if err != nil {
		return fmt.Errorf("unable to JSON-deserialize object: %s", err)
	}
	return nil
}
