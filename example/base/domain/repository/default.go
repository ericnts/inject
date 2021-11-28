package repository

import (
	"maple"
)

func init() {
	maple.Prepare(func(initiator maple.Initiator) {
		initiator.BindRepository(func() *Default {
			return &Default{}
		})
	})
}

type Default struct {
	maple.Repository
}

func (repo *Default) GetIP() string {
	return "IP"
}
