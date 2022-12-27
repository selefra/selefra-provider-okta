package tables

import (
	"context"
	"github.com/selefra/selefra-provider-okta/okta_client"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	"github.com/selefra/selefra-provider-okta/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableOktaTrustedOriginGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOktaTrustedOriginGenerator{}

func (x *TableOktaTrustedOriginGenerator) GetTableName() string {
	return "okta_trusted_origin"
}

func (x *TableOktaTrustedOriginGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOktaTrustedOriginGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOktaTrustedOriginGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOktaTrustedOriginGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := okta_client.Connect(ctx, taskClient.(*okta_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			input := query.Params{
				Limit: 200,
			}

			origins, resp, err := client.TrustedOrigin.ListOrigins(ctx, &input)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, origin := range origins {
				resultChannel <- origin
			}

			for resp.HasNextPage() {
				var nextOriginSet []*okta.TrustedOrigin
				resp, err = resp.Next(ctx, &nextOriginSet)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, origin := range nextOriginSet {
					resultChannel <- origin
				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, err)
		},
	}
}

func (x *TableOktaTrustedOriginGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOktaTrustedOriginGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("The title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the trusted origin.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by").ColumnType(schema.ColumnTypeString).Description("The ID of the user who created the trusted origin.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("origin").ColumnType(schema.ColumnTypeString).Description("The origin of the trusted origin.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("Current status of the trusted origin. Valid values are 'ACTIVE' or 'INACTIVE'.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scopes").ColumnType(schema.ColumnTypeJSON).Description("The scopes for the trusted origin. Valid values are 'CORS' or 'REDIRECT'.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("A unique key for the trusted origin.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("The timestamp when the trusted origin was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated").ColumnType(schema.ColumnTypeTimestamp).Description("The timestamp when the trusted origin was last updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_by").ColumnType(schema.ColumnTypeString).Description("The ID of the user who last updated the trusted origin.").Build(),
	}
}

func (x *TableOktaTrustedOriginGenerator) GetSubTables() []*schema.Table {
	return nil
}
