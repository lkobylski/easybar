package easybar

import (
	"testing"
)

func TestNewEasyBar(t *testing.T) {
	bar := NewEasyBar(100, "Test Task")

	if bar == nil {
		t.Fatal("Expected NewEasyBar to return a non-nil pointer")
	}

	if bar.GetMax() != 100 {
		t.Errorf("Expected max to be 100, got %d", bar.GetMax())
	}

	if bar.GetCurrent() != 0 {
		t.Errorf("Expected current progress to be 0, got %d", bar.GetCurrent())
	}

	if bar.name != "Test Task" {
		t.Errorf("Expected name to be 'Test Task', got '%s'", bar.name)
	}
}
