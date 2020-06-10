package sampleapp_test

import (
	"testing"

	. "github.com/pellared/sampleapp"
)

func TestFoo(t *testing.T) {
	want := "Bar"
	if got := Foo(); got != want {
		t.Errorf("Foo() = %v, want %v", got, want)
	}
}
