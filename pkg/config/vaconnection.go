package config

type VAConnection struct {
	Server string `yaml:"server"`
	User   string `yaml:"user"`
	PWord  string `yaml:"pword"`
	Token  string `yaml:"token"`
}
