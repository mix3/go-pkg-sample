package GoPkgSample

import (
	"testing"

	. "github.com/mix3/go-pkg-sample"
)

func TestMyType(t *testing.T) {
	i := MyType_
	if i != nil {
		t.Error("MyType_ is nil")
	}
	InitMyType()
	i = MyType_
	if i.Key != "init_key" {
		t.Error("MyType_.key is init_key")
	}
	if i.Val != "init_val" {
		t.Error("MyType_.val is init_val")
	}
}
