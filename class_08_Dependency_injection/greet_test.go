package class_08_Dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Victor")

	got := buffer.String()
	want := "Hello, Victor"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
