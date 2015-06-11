package registry

import (
	"fmt"
	"path"

	"github.com/spesnova/iruka/schema"
)

const (
	configVarsPrefix = "config-vars"
)

func (r *Registry) CreateConfigVars(appIdentity string, opts schema.ConfigVars) (schema.ConfigVars, error) {
	j, err := marshal(opts)
	if err != nil {
		fmt.Println(err.Error())
		return schema.ConfigVars{}, err
	}

	key := path.Join(r.keyPrefix, configVarsPrefix, appIdentity)
	res, err := r.etcd.Create(key, string(j), 0)
	if err != nil {
		fmt.Println(err.Error())
		return schema.ConfigVars{}, err
	}

	var configVars map[string]string
	err = unmarshal(res.Node.Value, &configVars)
	if err != nil {
		fmt.Println(err.Error())
		return schema.ConfigVars{}, err
	}

	return configVars, nil
}
