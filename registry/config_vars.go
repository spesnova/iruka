package registry

import (
	"fmt"
	"path"

	"github.com/spesnova/iruka/schema"
)

const (
	configVarsPrefix = "config-vars"
)

func (r *Registry) ConfigVars(appIdentity string) (schema.ConfigVars, error) {
	key := path.Join(r.keyPrefix, configVarsPrefix, appIdentity)
	res, err := r.etcd.Get(key, false, false)
	if err != nil {
		fmt.Println(err.Error())
		if isKeyNotFound(err) {
			err = nil
		}

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

func (r *Registry) UpdateConfigVars(appIdentity string, opts schema.ConfigVars) (schema.ConfigVars, error) {
	configVars, err := r.ConfigVars(appIdentity)
	if err != nil {
		fmt.Println(err.Error())
		return schema.ConfigVars{}, err
	}

	// merge key-value to update with current config-vars
	for k, v := range opts {
		if v == "" {
			delete(configVars, k)
		} else {
			configVars[k] = v
		}
	}

	j, err := marshal(configVars)
	if err != nil {
		fmt.Println(err.Error())
		return schema.ConfigVars{}, err
	}

	key := path.Join(r.keyPrefix, configVarsPrefix, appIdentity)
	res, err := r.etcd.Set(key, string(j), 0)
	if err != nil {
		fmt.Println(err.Error())
		return schema.ConfigVars{}, err
	}

	err = unmarshal(res.Node.Value, &configVars)
	if err != nil {
		fmt.Println(err.Error())
		return schema.ConfigVars{}, err
	}

	return configVars, nil
}
