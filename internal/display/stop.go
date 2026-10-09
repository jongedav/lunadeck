package display

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Stop(stateDir string) error {
	path := filepath.Join(stateDir, "active")

	if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("remove stream state marker %q: %w", path, err)
	}

	return nil
}
