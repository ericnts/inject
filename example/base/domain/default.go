package domain

import (
	"maple/example/base/domain/repository"
	"maple/mate"
)

func init() {
	mate.Prepare(func(initiator mate.Initiator) {
		initiator.BindService(func() *Default {
			return &Default{}
		})
		//initiator.InjectController(func(ctx freedom.Context) (service *Default) {
		//	initiator.FetchService(ctx, &service)
		//	return
		//})
	})
}

// Default .
type Default struct {
	//Worker    freedom.Worker
	DefRepo *repository.Default
}

// RemoteInfo .
func (s *Default) RemoteInfo() string {
	return s.DefRepo.GetIP()
}
