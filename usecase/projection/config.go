package projection

type TomlConfig struct {
	Projection map[string]interface{}
}

var GlobalConfig TomlConfig

type Config struct {
	Projection map[string]interface{}
}

var GlobalConfig2 Config
