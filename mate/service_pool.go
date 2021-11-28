package mate

import (
	"fmt"
	"reflect"
	"sync"
)

func newSerPool() *serPool {
	return &serPool{
		pool: make(map[reflect.Type]*sync.Pool),
	}
}

type serPool struct {
	pool map[reflect.Type]*sync.Pool
}
type serviceElement struct {
	workers       []reflect.Value
	serviceObject interface{}
}
func (pool *serPool) bind(t reflect.Type, f interface{}) {
	pool.pool[t] = &sync.Pool{
		New: func() interface{}{
			values := reflect.ValueOf(f).Call([]reflect.Value{})
			if len(values) == 0 {
				panic(fmt.Sprintf("[Freedom] BindService: func return to empty, %v", reflect.TypeOf(f)))
			}

			newService := values[0].Interface()
			result := serviceElement{serviceObject: newService, workers: []reflect.Value{}}
			allFields(newService, func(fieldValue reflect.Value) {
				kind := fieldValue.Kind()
				if kind == reflect.Interface &&
					//workerType.AssignableTo(fieldValue.Type()) &&
					fieldValue.CanSet() {
					//如果是运行时对象
					result.workers = append(result.workers, fieldValue)
					return
				}
				//globalApp.rpool.diRepoFromValue(fieldValue, &result)
				//globalApp.comPool.diInfraFromValue(fieldValue)
				//globalApp.factoryPool.diFactoryFromValue(fieldValue, &result)
				//
				//if fieldValue.IsNil() {
				//	return
				//}
				//if br, ok := fieldValue.Interface().(BeginRequest); ok {
				//	result.calls = append(result.calls, br)
				//}
			})

			//if br, ok := newService.(BeginRequest); ok {
			//	result.calls = append(result.calls, br)
			//}
			return result
		},
	}
}