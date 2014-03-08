package module2

import pkg "../"

func ModuleMethod2(from string) {
	pkg.ModuleMethod1(from)
}

func GetModuleType() *pkg.ModuleType {
	return pkg.ModuleType_
}
