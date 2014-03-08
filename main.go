package main

import (
	"log"
	. "./pkg"
	. "./pkg/pkg"
)

func main() {
	ModuleMethod1("main")
	ModuleMethod2("main")
	log.Printf("%v", ModuleType_)
	log.Printf("%v", GetModuleType())
	InitModuleType()
	log.Printf("%v", ModuleType_)
	log.Printf("%v", GetModuleType())
}
