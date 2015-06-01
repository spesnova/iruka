package scheduler

import (
	"github.com/spesnova/go-fleet/fleet"
)

type Scheduler struct {
	fleet fleet.Client
}

func NewScheduler(url string) *Scheduler {
	fleetClient := *fleet.NewClient(url)
	return &Scheduler{fleetClient}
}
