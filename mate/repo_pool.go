package mate

import "reflect"

func newRepositoryPool() *repositoryPool {
	return &repositoryPool{
		creator: make(map[reflect.Type]interface{}),
	}
}

type repositoryPool struct {
	creator map[reflect.Type]interface{}
}

func (r *repositoryPool) bind(outType reflect.Type, f interface{})  {
	r.creator[outType] = f
}