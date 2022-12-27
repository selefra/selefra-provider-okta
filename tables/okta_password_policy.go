package tables

import (
	"context"
	"github.com/selefra/selefra-provider-okta/okta_client"
	"time"

	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	"github.com/selefra/selefra-provider-okta/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableOktaPasswordPolicyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOktaPasswordPolicyGenerator{}

func (x *TableOktaPasswordPolicyGenerator) GetTableName() string {
	return "okta_password_policy"
}

func (x *TableOktaPasswordPolicyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOktaPasswordPolicyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOktaPasswordPolicyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOktaPasswordPolicyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := okta_client.Connect(ctx, taskClient.(*okta_client.Client).Config)

			input := &query.Params{}
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			input.Expand = "rules"
			input.Type = "PASSWORD"

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

type PolicyStructure struct {
	Embedded    interface{}                `json:"_embedded,omitempty"`
	Links       interface{}                `json:"_links,omitempty"`
	Settings    interface{}                `json:"settings,omitempty"`
	Conditions  *okta.PolicyRuleConditions `json:"conditions,omitempty"`
	Created     *time.Time                 `json:"created,omitempty"`
	Description string                     `json:"description,omitempty"`
	Id          string                     `json:"id,omitempty"`
	LastUpdated *time.Time                 `json:"lastUpdated,omitempty"`
	Name        string                     `json:"name,omitempty"`
	Priority    int64                      `json:"priority,omitempty"`
	Status      string                     `json:"status,omitempty"`
	System      *bool                      `json:"system,omitempty"`
	Type        string                     `json:"type,omitempty"`
}

// Generic policy returned by
func listPoliciesWithSettings(ctx context.Context, client okta.Client, qp *query.Params) ([]*PolicyStructure, *okta.Response, error) {
	url := "/api/v1/policies"
	if qp != nil {
		url = url + qp.String()
	}

	requestExecutor := client.GetRequestExecutor()
	req, err := requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var policies []*PolicyStructure

	resp, err := requestExecutor.Do(ctx, req, &policies)
	if err != nil {
		return nil, resp, err
	}

	return policies, resp, nil
}

func (x *TableOktaPasswordPolicyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOktaPasswordPolicyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when the Policy was last modified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("priority").ColumnType(schema.ColumnTypeInt).Description("Priority of the Policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("Status of the Policy: ACTIVE or INACTIVE.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("settings").ColumnType(schema.ColumnTypeJSON).Description("Settings of the password policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("The title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("Description of the Policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when the Policy was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("system").ColumnType(schema.ColumnTypeBool).Description("This is set to true on system policies, which cannot be deleted.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("conditions").ColumnType(schema.ColumnTypeJSON).Description("Conditions for Policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rules").ColumnType(schema.ColumnTypeJSON).Description("Each Policy may contain one or more Rules. Rules, like Policies, contain conditions that must be satisfied for the Rule to be applied.").
			Extractor(column_value_extractor.StructSelector("Embedded.rules")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the Policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Identifier of the Policy.").Build(),
	}
}

func (x *TableOktaPasswordPolicyGenerator) GetSubTables() []*schema.Table {
	return nil
}
