package mate

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	globalApp     *Application
	globalAppOnce sync.Once
	prepares      []func(Initiator)
)

type Initiator interface {
	BindController(relativePath string, controller interface{}, handlers ...gin.HandlerFunc)
	BindService(f interface{})
	BindRepository(f interface{})
}

func Prepare(f func(Initiator)) {
	prepares = append(prepares, f)
}
