package mytest

import (
	"fmt"
	"testing"
)

func TestType(t *testing.T) {
	if tp := fmt.Sprintf("%T", constBad); tp != "int" {
		t.Errorf("Wrong type - got %s, want 'int'", tp)
	}
}
