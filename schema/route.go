package schema

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
)

type Route struct {
	ID        uuid.UUID `json:"id"`
	Location  string    `json:"location"`
	Upstream  string    `json:"upstream"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RouteCreateOpts struct {
	Location string `json:"location"`
	Upstream string `json:"upstream"`
}

type RouteUpdateOpts struct {
	ID       uuid.UUID `json:"id"`
	Location string    `json:"location"`
	Upstream string    `json:"upstream"`
}
