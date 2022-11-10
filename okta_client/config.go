package okta_client

// OktaProviderConfigs defines Provider Configuration
type OktaProviderConfigs struct {
	Providers []OktaProviderConfig `yaml:"providers"  mapstructure:"providers"`
}

// OktaProviderConfig Domain="https://<YOUR_OKTA_DOMAIN>.okta.com/"
type OktaProviderConfig struct {
	Domain string `yaml:"domain"  mapstructure:"domain"`
	Token  string `yaml:"token,omitempty"  mapstructure:"token"`
}
