package maple

import "maple/mate"

func init() {
	initApp()
}

type (
	Initiator = mate.Initiator
	Repository = mate.Repository
)

func Prepare(f func(initiator Initiator)) {
	mate.Prepare(f)
}
