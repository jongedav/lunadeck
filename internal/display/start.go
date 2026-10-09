package display

import (
	"fmt"
	"os"
	"path/filepath"
)

func Start(stateDir string) error {
	path := filepath.Join(stateDir, "active")
	file, err := os.Create(path)

	if err != nil {
		return fmt.Errorf("creating stream state marker %q: %w", path, err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("closing stream state marker %q: %w", path, err)
	}

	return nil
}
