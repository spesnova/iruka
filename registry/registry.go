package registry

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/coreos/go-etcd/etcd"
)

const (
	DefaultMachines  = "http://localhost:4001"
	DefaultKeyPrefix = "/iruka/"
)

type Registry struct {
	etcd      etcd.Client
	keyPrefix string
}

func NewRegistry(machines string, keyPrefix string) *Registry {
	if os.Getenv("IRUKA_ETCD_MACHINES") != "" {
		machines = os.Getenv("IRUKA_ETCD_MACHINES")
	}

	if machines == "" {
		machines = DefaultMachines
	}

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
