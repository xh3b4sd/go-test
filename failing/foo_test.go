package foo_test

import (
	"testing"

	"github.com/xh3b4sd/go-test/failing"
)

func Test_Foo_001(t *testing.T) {
	if foo.Foo() != "foo" {
		t.Fatalf("Foo() != foo")
	}
}
