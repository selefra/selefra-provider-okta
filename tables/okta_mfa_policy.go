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

type TableOktaMfaPolicyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOktaMfaPolicyGenerator{}

func (x *TableOktaMfaPolicyGenerator) GetTableName() string {
	return "okta_mfa_policy"
}

func (x *TableOktaMfaPolicyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOktaMfaPolicyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOktaMfaPolicyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOktaMfaPolicyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := okta_client.Connect(ctx, taskClient.(*okta_client.Client).Config)

			input := &query.Params{}
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			input.Expand = "rules"
			input.Type = "MFA_ENROLL"

			policies, resp, err := listPoliciesWithSettings(ctx, *client, input)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, policy := range policies {
				resultChannel <- policy
			}

			for resp.HasNextPage() {
				var nextPolicySet []*okta.Policy
				resp, err = resp.Next(ctx, &nextPolicySet)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, policy := range nextPolicySet {
					resultChannel <- policy
				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, err)
		},
	}
}

func (x *TableOktaMfaPolicyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOktaMfaPolicyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the Policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when the Policy was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when the Policy was last modified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("system").ColumnType(schema.ColumnTypeBool).Description("This is set to true on system policies, which cannot be deleted.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("The title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Identifier of the Policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("Description of the Policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("priority").ColumnType(schema.ColumnTypeInt).Description("Priority of the Policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("Status of the Policy: ACTIVE or INACTIVE.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("conditions").ColumnType(schema.ColumnTypeJSON).Description("Conditions for Policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rules").ColumnType(schema.ColumnTypeJSON).Description("Each Policy may contain one or more Rules. Rules, like Policies, contain conditions that must be satisfied for the Rule to be applied.").
			Extractor(column_value_extractor.StructSelector("Embedded.rules")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("settings").ColumnType(schema.ColumnTypeJSON).Description("Settings of the password policy.").Build(),
	}
}

func (x *TableOktaMfaPolicyGenerator) GetSubTables() []*schema.Table {
	return nil
}
