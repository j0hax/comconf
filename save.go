// comconf is my very opinionated configuration management
package comconf

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Save writes the configuration to disk as a TOML file.
func (c *ConfManager[T]) Save() error {
	file, err := c.OpenFile(MainFile)
	if err != nil {
		return err
	}

	enc := toml.NewEncoder(file)
	return enc.Encode(c.Configuration)
}

// SaveExit saves the configuration file.
//
// If an error occurs, log.Fatal is called, otherwise os.Exit(0)
func (c *ConfManager[T]) SaveExit() {
	err := c.Save()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
