package config

type VAConnection struct {
	Server        string `yaml:"server"`
	User          string `yaml:"user"`
	PWord         string `yaml:"pword"`
	Token         string `yaml:"token"`
	DumpRequest   bool   `yaml:"dumprequests"`
	DumpResponses bool   `yaml:"dumpresponses"`
	Debug         bool   `yaml:"debug"`
}
