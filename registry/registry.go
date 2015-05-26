package registry

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/coreos/go-etcd/etcd"
)

const DefaultKeyPrefix = "/iruka/"

type Registry struct {
	etcd      etcd.Client
	keyPrefix string
}

func NewRegistry(machines string, keyPrefix string) *Registry {
	m := strings.Split(machines, ",")
	etcdClient := *etcd.NewClient(m)
	return &Registry{etcdClient, keyPrefix}
}

func marshal(obj interface{}) (string, error) {
	encoded, err := json.Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("unable to JSON-serialize object: %s", err)
	}
	return string(encoded), nil
}

func unmarshal(val string, obj interface{}) error {
	err := json.Unmarshal([]byte(val), &obj)
	if err != nil {
		return fmt.Errorf("unable to JSON-deserialize object: %s", err)
	}
	return nil
}

func isKeyNotFound(err error) bool {
	e, ok := err.(*etcd.EtcdError)
	return ok && e.ErrorCode == 100
}
