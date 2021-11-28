package controller

import (
	"maple"
	"maple/example/base/domain"
)

func init() {
	maple.Prepare(func(initiator maple.Initiator) {
		initiator.BindController("/", &Default{})
	})
}

// Default .
type Default struct {
	Sev *domain.Default
}

// Get handles the GET: / route.
func (c *Default) Get() string {
	return c.Sev.RemoteInfo()
}
