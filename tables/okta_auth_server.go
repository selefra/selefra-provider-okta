package tables

import (
	"context"
	"github.com/selefra/selefra-provider-okta/okta_client"
	"strings"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	"github.com/selefra/selefra-provider-okta/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableOktaAuthServerGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOktaAuthServerGenerator{}

func (x *TableOktaAuthServerGenerator) GetTableName() string {
	return "okta_auth_server"
}

func (x *TableOktaAuthServerGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOktaAuthServerGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOktaAuthServerGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOktaAuthServerGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := okta_client.Connect(ctx, taskClient.(*okta_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			input := query.Params{
				Limit: 200,
			}

			servers, resp, err := client.AuthorizationServer.ListAuthorizationServers(ctx, &input)
			if err != nil {
				if strings.Contains(err.Error(), "Not found") {
					return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, server := range servers {
				resultChannel <- server
			}

			for resp.HasNextPage() {
				var nextAuthorizationServerSet []*okta.AuthorizationServer
				resp, err = resp.Next(ctx, &nextAuthorizationServerSet)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, server := range nextAuthorizationServerSet {
					resultChannel <- server
				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, err)
		},
	}
}

func (x *TableOktaAuthServerGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOktaAuthServerGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique key for the authorization server.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("issuer").ColumnType(schema.ColumnTypeString).Description("The issuer URI of the authorization server.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("issuer_mode").ColumnType(schema.ColumnTypeString).Description("The issuer mode of the authorization server.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The status of the authorization server.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("The title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name for the authorization server.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when the authorization server was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("A human-readable description of the authorization server.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when the authorization server was last updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("audiences").ColumnType(schema.ColumnTypeJSON).Description("The audiences of the authorization server.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credentials").ColumnType(schema.ColumnTypeJSON).Description("The authorization server credentials.").Build(),
	}
}

func (x *TableOktaAuthServerGenerator) GetSubTables() []*schema.Table {
	return nil
}
