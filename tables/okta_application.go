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

type TableOktaApplicationGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOktaApplicationGenerator{}

func (x *TableOktaApplicationGenerator) GetTableName() string {
	return "okta_application"
}

func (x *TableOktaApplicationGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOktaApplicationGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOktaApplicationGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOktaApplicationGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := okta_client.Connect(ctx, taskClient.(*okta_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			input := query.Params{
				Limit: 200,
			}

			applications, resp, err := client.Application.ListApplications(ctx, &input)
			if err != nil {
				if strings.Contains(err.Error(), "Not found") {
					return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, app := range applications {
				resultChannel <- app
			}

			for resp.HasNextPage() {
				var nextApplicationSet []*okta.Application
				resp, err = resp.Next(ctx, &nextApplicationSet)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, app := range nextApplicationSet {
					resultChannel <- app
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, err)
		},
	}
}

func (x *TableOktaApplicationGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOktaApplicationGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Unique key for app definition.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label").ColumnType(schema.ColumnTypeString).Description("User-defined display name for app.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("Current status of app. Valid values are ACTIVE or INACTIVE.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("visibility").ColumnType(schema.ColumnTypeJSON).Description("Visibility settings for app.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("The title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique key for app.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when user was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when app was last updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sign_on_mode").ColumnType(schema.ColumnTypeString).Description("Authentication mode of app. Can be one of AUTO_LOGIN, BASIC_AUTH, BOOKMARK, BROWSER_PLUGIN, Custom, OPENID_CONNECT, SAML_1_1, SAML_2_0, SECURE_PASSWORD_STORE and WS_FEDERATION.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("settings").ColumnType(schema.ColumnTypeJSON).Description("Settings for app.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credentials").ColumnType(schema.ColumnTypeJSON).Description("Credentials for the specified signOnMode.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("accessibility").ColumnType(schema.ColumnTypeJSON).Description("Access settings for app.").Build(),
	}
}

func (x *TableOktaApplicationGenerator) GetSubTables() []*schema.Table {
	return nil
}
