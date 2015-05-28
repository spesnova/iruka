package iruka

import (
	"github.com/spesnova/iruka/schema"
)

func (c *Client) AppList() ([]schema.App, error) {
	var appsRes []schema.App
	return appsRes, c.Get(&appsRes, "/apps")
}
