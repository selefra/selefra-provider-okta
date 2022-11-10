package users

import (
	"context"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	"github.com/selefra/selefra-provider-okta/okta_client"
	"github.com/selefra/selefra-provider-okta/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableOktaUsersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOktaUsersGenerator{}

func (x *TableOktaUsersGenerator) GetTableName() string {
	return "okta_users"
}

func (x *TableOktaUsersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOktaUsersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOktaUsersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableOktaUsersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			api := client.(*okta_client.Client)
			users, resp, err := api.Okta.User.ListUsers(ctx, query.NewQueryParams(query.WithLimit(200), query.WithAfter("")))
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			if len(users) == 0 {
				return nil
			}
			resultChannel <- users
			for resp != nil && resp.HasNextPage() {
				var nextUserSet []*okta.User
				resp, err = resp.Next(ctx, &nextUserSet)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- nextUserSet
			}
			return nil
		},
	}
}

func (x *TableOktaUsersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOktaUsersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("profile").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("activated").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_changed").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type_created").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("Type.Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type_description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type.Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credentials_password_hash_value").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Credentials.Password.Hash.Value")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credentials_password_value").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Credentials.Password.Value")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credentials_provider_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Credentials.Provider.Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credentials_recovery_question").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Credentials.RecoveryQuestion.Question")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credentials_password_hash_algorithm").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Credentials.Password.Hash.Algorithm")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credentials_provider_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Credentials.Provider.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type_last_updated").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(column_value_extractor.StructSelector("Type.LastUpdated")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type_last_updated_by").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type.LastUpdatedBy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credentials_password_hash_salt_order").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Credentials.Password.Hash.SaltOrder")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_login").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password_changed").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transitioning_to_status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type_default").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Type.Default")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credentials_password_hook_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Credentials.Password.Hook.Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credentials_recovery_question_answer").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Credentials.RecoveryQuestion.Answer")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type_display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type.DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type.Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credentials_password_hash_salt").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Credentials.Password.Hash.Salt")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credentials_password_hash_work_factor").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("Credentials.Password.Hash.WorkFactor")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type_created_by").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type.CreatedBy")).Build(),
	}
}

func (x *TableOktaUsersGenerator) GetSubTables() []*schema.Table {
	return nil
}
