package GoPkgSample

type MyType struct {
	Key string
	Val string
}

var MyType_ *MyType

func InitMyType() {
	MyType_ = &MyType{
		Key: "init_key",
		Val: "init_val",
	}
}
