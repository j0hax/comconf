package comconf

import "github.com/BurntSushi/toml"

// Load attempts to load the existing configuration file.
//
// If the file is empty, the defaults are used.
func (c *ConfManager[T]) Load(defaults T) error {
	file, err := c.OpenFile(MainFile)
	if err != nil {
		return err
	}

	info, err := file.Stat()
	if err != nil {
		return err
	}

	// Check if file is empty
	if info.Size() == 0 {
		c.Configuration = defaults
	}

	// Otherwise, decode
	dec := toml.NewDecoder(file)
	_, err = dec.Decode(&c.Configuration)
	if err != nil {
		return err
	}

	return nil
}
