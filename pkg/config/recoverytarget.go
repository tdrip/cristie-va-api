package config

type RecoveryTarget struct {
	Hostname   string `yaml:"hostname"`
	MacAddress string `yaml:"macaddress"`
	BiosUuid   string `yaml:"biosuuid"`
	OS         string `yaml:"os"`
}
