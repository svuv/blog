package config

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	env  string `yaml:"env"`
}
