package GoPkgSample

import (
	"testing"

	. "github.com/mix3/go-pkg-sample"
)

func TestReverse(t *testing.T) {
	i := "ABCD"
	o := Reverse(i)
	e := "DCBA"
	if o != e {
		t.Errorf("%q expected %q got %q", i, e, o)
	}
}
