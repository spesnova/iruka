package schema

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
)

type Domain struct {
	ID        uuid.UUID `json:"id"`
	Hostname  string    `json:"hostname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DomainCreateOpts struct {
	Hostname string `json:"hostname"`
}
