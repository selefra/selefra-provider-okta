package okta_client

type Configs struct {
	Providers []Config `yaml:"providers"  mapstructure:"providers"`
}

type Config struct {
	Domain string `yaml:"domain,omitempty" mapstructure:"domain"`
	Token  string `yaml:"token,omitempty" mapstructure:"token"`
}
