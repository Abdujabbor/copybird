package core

import (
	"sync"
)

var _modules = []Module{}
var _modulesMutex = sync.Mutex{}

func RegisterModule(module Module) {
	_modulesMutex.Lock()
	_modules = append(_modules, module)
	_modulesMutex.Unlock()
}

func GetModule(moduleGroup ModuleGroup, moduleType ModuleType, name string) Module {
	_modulesMutex.Lock()
	defer _modulesMutex.Unlock()
	for _, module := range _modules {
		if module.GetName() == name && module.GetGroup() == moduleGroup && module.GetType() == moduleType {
			return module
		}
	}
	return nil
}