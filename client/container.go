package iruka

import (
	"github.com/spesnova/iruka/schema"
)

func (c *Client) ContainerList(appIdentity string) ([]schema.Container, error) {
	var cRes []schema.Container
	return cRes, c.Get(&cRes, "/apps/"+appIdentity+"/containers")
}
