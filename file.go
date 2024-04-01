package comconf

import (
	"os"
	"path/filepath"
)

// Directory returns the path to the configurable's configuration directory.
//
// The configuration directory is determined by os.UserConfigDir() and the configurable's name.
func (c *ConfManager[T]) Directory() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, c.namespace), nil
}

// OpenFile returns a file handler to the file specified by name, located inside the configurable's directory.
//
// If the directory or its parents do not exist, they will be created. If a file does not exist, it will be opened for writing.
func (c *ConfManager[T]) OpenFile(name string) (*os.File, error) {
	folder, err := c.Directory()
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(folder, 0755)
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(folder, name)

	return os.OpenFile(configPath, os.O_RDWR|os.O_CREATE, 0755)
}
