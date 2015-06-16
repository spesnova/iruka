package iruka

import (
	"github.com/spesnova/iruka/schema"
)

func (c *Client) ConfigVarsInfo(appIdentity string) (schema.ConfigVars, error) {
	var configVarsRes schema.ConfigVars
	return configVarsRes, c.Get(&configVarsRes, "/apps/"+appIdentity+"/config-vars")
}

func (c *Client) ConfigVarsUpdate(appIdentity string, opts map[string]string) (schema.ConfigVars, error) {
	var configVarsRes schema.ConfigVars
	return configVarsRes, c.Update(&configVarsRes, "/apps/"+appIdentity+"/config-vars", opts)
}
