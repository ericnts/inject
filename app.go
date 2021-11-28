package maple

import "maple/mate"

var (
	app *mate.Application
)

func initApp() {
	app = mate.NewApplication()
}

type Application interface {
	Run(addr ...string) (err error)
}

func NewApplication() Application {
	return app
}