package provider

import (
	"context"

	"github.com/selefra/selefra-provider-okta/okta_client"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
)

const Version = "v0.0.1"

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      "okta",
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var oktaConfig okta_client.Configs

				err := config.Unmarshal(&oktaConfig.Providers)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				if len(oktaConfig.Providers) == 0 {
					oktaConfig.Providers = append(oktaConfig.Providers, okta_client.Config{})
				}

				clients, err := okta_client.NewClients(oktaConfig)

				if err != nil {
					clientMeta.ErrorF("new clients err: %s", err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg("account information not found")
				}

				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					res = append(res, clients[i])
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# domain: "https://<YOUR_OKTA_DOMAIN>.okta.com"
# token: "<YOUR_OKTA_TOKEN>"`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var client_config okta_client.Configs
				err := config.Unmarshal(&client_config.Providers)

				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				if len(client_config.Providers) == 0 {
					return schema.NewDiagnostics().AddErrorMsg("analysis config err: no configuration")
				}

				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				"",
				"N/A",
				"not_supported",
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
