package treenode

import (
	"testing"
)

func TestNew(t *testing.T) {
	expected := 0
	if treeNodeValue := New(expected).Value; treeNodeValue != expected {
		t.Errorf("Expected %d but got %d", expected, treeNodeValue)
	}
}
