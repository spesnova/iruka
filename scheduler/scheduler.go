package scheduler

import (
	"os"

	"github.com/spesnova/go-fleet/fleet"
)

const (
	DefaultAPIURL = "http://localhost:4002"
)

type Scheduler struct {
	fleet fleet.Client
}

func NewScheduler(url string) *Scheduler {
	if os.Getenv("IRUKA_FLEET_API_URL") != "" {
		url = os.Getenv("IRUKA_FLEET_API_URL")
	}

	if url == "" {
		url = DefaultAPIURL
	}

	fleetClient := *fleet.NewClient(url)
	return &Scheduler{fleetClient}
}
