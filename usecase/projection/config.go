package projection

type TomlConfig struct {
	Projection map[string]interface{}
}

var GlobalConfig TomlConfig
