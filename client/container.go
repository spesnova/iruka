package iruka

import (
	"github.com/spesnova/iruka/schema"
)

func (c *Client) ContainerCreate(appIdentity string, opts schema.ContainerCreateOpts) error {
	var cRes schema.Container
	return c.Post(&cRes, "/apps/"+appIdentity+"/containers", opts)
}

func (c *Client) ContainerDelete(appIdentity, containerIdentity string) error {
	return c.Delete("/apps/" + appIdentity + "/containers/" + containerIdentity)
}

func (c *Client) ContainerList(appIdentity string) ([]schema.Container, error) {
	var cRes []schema.Container
	return cRes, c.Get(&cRes, "/apps/"+appIdentity+"/containers")
}
