package mate

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

var _ Initiator = (*Application)(nil)

func NewApplication() *Application {
	globalAppOnce.Do(func() {
		globalApp = &Application{
			gin:     gin.New(),
			serPool: newSerPool(),
			repPool: newRepositoryPool(),
			facPool: newFactoryPool(),
		}
	})
	return globalApp
}

type Application struct {
	gin     *gin.Engine
	serPool *serPool
	repPool *repositoryPool
	facPool *factoryPool
}

func (a *Application) BindController(relativePath string, controller interface{}, handlers ...gin.HandlerFunc) {
	typ := reflect.TypeOf(controller)
	for i := 0; i < typ.NumMethod(); i++ {
		httpMethod, httpPath, err := parseMethod(typ.Method(i), func(s string) bool {
			return false
		})
		if err != nil {
			continue
		}
		index:=i
		t :=typ.Elem()
		a.gin.Handle(httpMethod, httpPath, func(context *gin.Context) {
			reflect.New(t).Method(index).Call([]reflect.Value{})
		})
	}
}

func (a *Application) BindService(f interface{}) {
	outType, err := parsePoolFunc(f)
	if err != nil {
		panic(err)
	}
	a.serPool.bind(outType, f)
}

func (a *Application) BindRepository(f interface{}) {
	outType, err := parsePoolFunc(f)
	if err != nil {
		panic(err)
	}
	a.repPool.bind(outType, f)
}

func (a *Application) BindFactory(f interface{}) {
	outType, err := parsePoolFunc(f)
	if err != nil {
		panic(err)
	}
	a.facPool.bind(outType, f)
}

func (a *Application) Run(addr ...string) (err error) {
	for i := 0; i < len(prepares); i++ {
		prepares[i](a)
	}
	return a.gin.Run(addr...)
}
