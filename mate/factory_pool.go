package mate

import (
	"reflect"
)

func newFactoryPool() *factoryPool {
	return &factoryPool{
		creator: make(map[reflect.Type]interface{}),
	}
}

type factoryPool struct {
	creator map[reflect.Type]interface{}
}

func (fa *factoryPool) bind(outType reflect.Type, f interface{}) {
	fa.creator[outType] = f
}
