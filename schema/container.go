package schema

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
)

type Container struct {
	ID          uuid.UUID `json:"id"`
	AppID       uuid.UUID `json:"app_id"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Size        string    `json:"size"`
	Command     string    `json:"command"`
	Type        string    `json:"type"`
	Ports       []int     `json:"ports"`
	DesireState string    `json:"desire_state"`
	State       string    `json:"state"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ContainerCreateOpts struct {
	Image   string `json:"image"`
	Size    string `json:"size"`
	Command string `json:"command"`
	Type    string `json:"type"`
	Ports   []int  `json:"ports"`
}

type ContainerUpdateOpts struct {
	Image   string `json:"image"`
	Size    string `json:"size"`
	Command string `json:"command"`
	Type    string `json:"type"`
	Ports   []int  `json:"ports"`
}
