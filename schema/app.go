package schema

import (
	"time"
)

type App struct {
	ID        string    `json:"id"`
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
