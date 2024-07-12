package networks

type DhcpConfig struct {
	Bindings []Binding `json:"bindings,omitempty"`
	Enabled  bool      `json:"enabled,omitempty"`
}
