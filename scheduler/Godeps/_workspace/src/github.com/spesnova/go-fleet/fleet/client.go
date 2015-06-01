package fleet

const (
	basePath = "/fleet/v1"
)

type Client struct {
	URL string
}

// NewClient creates a basic client that is configured to be used
// with the given fleet HTTP API URL.
func NewClient(url string) *Client {
	return &Client{
		URL: url,
	}
}
