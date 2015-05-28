package iruka

import (
	"github.com/spesnova/iruka/schema"
)

func (c *Client) AppCreate(opts schema.AppCreateOpts) (schema.App, error) {
	var appRes schema.App
	return appRes, c.Post(&appRes, "/apps", opts)
}

func (c *Client) AppDelete(appIdentity string) error {
	return c.Delete("/apps/" + appIdentity)
}

func (c *Client) AppList() ([]schema.App, error) {
	var appsRes []schema.App
	return appsRes, c.Get(&appsRes, "/apps")
}

func (c *Client) AppUpdate(appIdentity string, opts schema.AppUpdateOpts) (schema.App, error) {
	var appRes schema.App
	return appRes, c.Update(&appRes, "/apps/"+appIdentity, opts)
}
