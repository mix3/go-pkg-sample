package GoPkgSample

import (
	"testing"

	. "github.com/mix3/go-pkg-sample"
	. "github.com/mix3/go-pkg-sample/nest-pkg"
)

func TestMyType(t *testing.T) {
	i := MyType_
	if i != nil {
		t.Error("MyType_ is nil")
	}
	InitMyType2()
	i = MyType_
	if i.Key != "init_key" {
		t.Error("MyType_.key is init_key")
	}
	if i.Val != "init2_val" {
		t.Error("MyType_.val is init2_val")
	}
}
