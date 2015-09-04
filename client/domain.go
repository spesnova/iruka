package iruka

import (
	"github.com/spesnova/iruka/schema"
)

func (c *Client) DomainCreate(appIdentity string, opts schema.DomainCreateOpts) (schema.Domain, error) {
	var domainRes schema.Domain
	return domainRes, c.Post(&domainRes, "/apps/"+appIdentity+"/domains", opts)
}

func (c *Client) DomainDelete(appIdentity, domainIdentity string) error {
	return c.Delete("/apps/" + appIdentity + "/domains/" + domainIdentity)
}

func (c *Client) DomainList(appIdentity string) ([]schema.Domain, error) {
	var domainsRes []schema.Domain
	return domainsRes, c.Get(&domainsRes, "/apps/"+appIdentity+"/domains")
}
