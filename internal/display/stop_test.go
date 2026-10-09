package display

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestStopRemovesMarker(t *testing.T) {
	stateDir := t.TempDir()
	marker := filepath.Join(stateDir, "active")

	if err := os.WriteFile(marker, nil, 0600); err != nil {
		t.Fatalf("create marker: %v", err)
	}

	if err := Stop(stateDir); err != nil {
		t.Fatalf("Stop returned an error: %v", err)
	}

	if _, err := os.Stat(marker); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expected marker to be absent, got: %v", err)
	}

	// Stopping again should also succeed.
	if err := Stop(stateDir); err != nil {
		t.Fatalf("second Stop returned an error: %v", err)
	}
}

func TestStopWithoutMarker(t *testing.T) {
	if err := Stop(t.TempDir()); err != nil {
		t.Fatalf("Stop without marker returned an error: %v", err)
	}
}
