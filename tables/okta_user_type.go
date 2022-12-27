package tables

import (
	"context"
	"github.com/selefra/selefra-provider-okta/okta_client"
	"strings"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/selefra/selefra-provider-okta/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableOktaUserTypeGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOktaUserTypeGenerator{}

func (x *TableOktaUserTypeGenerator) GetTableName() string {
	return "okta_user_type"
}

func (x *TableOktaUserTypeGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOktaUserTypeGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOktaUserTypeGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOktaUserTypeGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := okta_client.Connect(ctx, taskClient.(*okta_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			userTypes, resp, err := client.UserType.ListUserTypes(ctx)
			if err != nil {
				if strings.Contains(err.Error(), "Not found") {
					return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, userType := range userTypes {
				resultChannel <- userType

			}

			for resp.HasNextPage() {
				var nextUserTypeSet []*okta.UserType
				resp, err = resp.Next(ctx, &nextUserTypeSet)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, user := range nextUserTypeSet {
					resultChannel <- user

				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, err)
		},
	}
}

func (x *TableOktaUserTypeGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOktaUserTypeGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated_by").ColumnType(schema.ColumnTypeString).Description("The user ID of the last user to edit this type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The display name for the type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default").ColumnType(schema.ColumnTypeBool).Description("Boolean to indicate if this type is the default.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when the User Type was last updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by").ColumnType(schema.ColumnTypeString).Description("The user ID of the creator of this type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("A human-readable description of the type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("The title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name for the type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique key for the User Type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when the User Type was created.").Build(),
	}
}

func (x *TableOktaUserTypeGenerator) GetSubTables() []*schema.Table {
	return nil
}
