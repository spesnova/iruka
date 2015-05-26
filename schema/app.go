package schema

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
)

type App struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	WebURL    string    `json:"web_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AppCreateOpts struct {
	Name string `json:"name"`
}

type AppUpdateOpts struct {
	Name string `json:"name"`
}
