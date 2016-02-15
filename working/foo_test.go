package foo

import (
	"testing"
)

func Test_Foo_001(t *testing.T) {
	if Foo() != "foo" {
		t.Fatalf("Foo() != foo")
	}
}
