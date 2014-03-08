package module1

import "log"

type ModuleType struct{}

var ModuleType_ *ModuleType

func InitModuleType() {
	ModuleType_ = &ModuleType{}
}

func ModuleMethod1(from string) {
	log.Printf("call ModuleMethod1 from %s", from)
}
