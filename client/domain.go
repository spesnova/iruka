package iruka

import (
	"github.com/spesnova/iruka/schema"
)

func (c *Client) DomainList(appIdentity string) ([]schema.Domain, error) {
	var domainsRes []schema.Domain
	return domainsRes, c.Get(&domainsRes, "/apps/"+appIdentity+"/domains")
}
