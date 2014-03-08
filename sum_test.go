package GoPkgSample

import (
	"testing"

	mytools "github.com/mix3/go-pkg-sample"
)

func TestSum(t *testing.T) {
	got := mytools.Sum(10, 20)
	expected := 30
	if got != expected {
		t.Errorf("got %v warn %v", got, expected)
	}
}
