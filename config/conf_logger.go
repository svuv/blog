package config

type Logger struct {
	Level        string `yaml:"Level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"log_in_console"`
}