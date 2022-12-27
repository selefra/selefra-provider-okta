package okta_client

import (
	"context"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

func Connect(ctx context.Context, config *Config) (*okta.Client, error) {
	_, client, err := okta.NewClient(ctx, okta.WithOrgUrl(config.Domain), okta.WithToken(config.Token), okta.WithRequestTimeout(30), okta.WithRateLimitMaxRetries(5))
	if err != nil {
		return nil, err
	}
	return client, err
}

func ExtractorTimestamp(path string) schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		v, err := column_value_extractor.StructSelector(path).Extract(ctx, clientMeta, client, task, row, column, result)
		return v, err
	})
}
