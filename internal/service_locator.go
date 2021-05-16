package internal

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/kataras/iris/v12/context"
)

func newServiceLocator() *ServiceLocatorImpl {
	return &ServiceLocatorImpl{}
}

// ServiceLocatorImpl The implementation of service positioning.
type ServiceLocatorImpl struct {
	beginCallBack []func(Worker)
	endCallBack   []func(Worker)
}

// InstallBeginCallBack Install the callback function that started.
// Triggered before the callback function.
func (locator *ServiceLocatorImpl) InstallBeginCallBack(f func(Worker)) {
	locator.beginCallBack = append(locator.beginCallBack, f)
}

// InstallEndCallBack The callback function at the end of the installation.
// Triggered after the callback function.
func (locator *ServiceLocatorImpl) InstallEndCallBack(f func(Worker)) {
	locator.endCallBack = append(locator.endCallBack, f)
}

// Call Called with a service locator.
func (locator *ServiceLocatorImpl) Call(fun interface{}) {
	ctx := context.NewContext(globalApp.IrisApp)
	ctx.BeginRequest(nil, new(http.Request))
	worker := newWorker(ctx)
	ctx.Values().Set(WorkerKey, worker)
	worker.bus = newBus(make(http.Header))

	serviceObj, err := parseCallServiceFunc(fun)
	if err != nil {
		panic(fmt.Sprintf("[Freedom] ServiceLocatorImpl.Call, %v : %s", fun, err.Error()))
	}
	newService := globalApp.pool.create(worker, serviceObj)
	for _, beginCallBack := range locator.beginCallBack {
		beginCallBack(worker)
	}
	reflect.ValueOf(fun).Call([]reflect.Value{reflect.ValueOf(newService.(serviceElement).serviceObject)})
	for _, endCallBack := range locator.endCallBack {
		endCallBack(worker)
	}

	if worker.IsDeferRecycle() {
		return
	}
	globalApp.pool.free(newService)
}
