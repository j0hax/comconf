package comconf

const MainFile = "config.toml"

// The ConfManager handles an existing configuration data type.
type ConfManager[T any] struct {
	Configuration T
	namespace     string
}

// New attempts to load a configuration file based on its namespace.
//
// If the configuration file is empty, the defaults are used.
func New[T any](namespace string, defaults T) (*ConfManager[T], error) {
	cfg := &ConfManager[T]{
		namespace: namespace,
	}

	err := cfg.Load(defaults)
	return cfg, err
}
