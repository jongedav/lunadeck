package display

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStartCreatesMarker(t *testing.T) {
	stateDir := t.TempDir()

	if err := Start(stateDir); err != nil {
		t.Fatalf("Start returned an error: %v", err)
	}

	marker := filepath.Join(stateDir, "active")
	info, err := os.Stat(marker)
	if err != nil {
		t.Fatalf("check marker: %v", err)
	}

	if !info.Mode().IsRegular() {
		t.Fatal("expected marker to be a regular file")
	}

	if info.Size() != 0 {
		t.Fatalf("expected empty marker, got %d bytes", info.Size())
	}
}

func TestStartWithMissingDirectory(t *testing.T) {
	stateDir := filepath.Join(t.TempDir(), "missing")

	if err := Start(stateDir); err == nil {
		t.Fatal("expected an error for a missing state directory")
	}
}
