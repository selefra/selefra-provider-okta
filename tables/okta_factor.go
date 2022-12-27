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

type TableOktaFactorGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOktaFactorGenerator{}

func (x *TableOktaFactorGenerator) GetTableName() string {
	return "okta_factor"
}

func (x *TableOktaFactorGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOktaFactorGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOktaFactorGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOktaFactorGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := okta_client.Connect(ctx, taskClient.(*okta_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			var userId string
			var userName string
			userData := task.ParentRawResult.(*okta.User)
			userId = userData.Id
			userProfile := *userData.Profile
			userName = userProfile["login"].(string)

			if userId == "" {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}

			factors, resp, err := client.UserFactor.ListFactors(ctx, userId)
			if err != nil {
				if strings.Contains(err.Error(), "Not found") {
					return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, factor := range factors {
				resultChannel <- UserFactorInfo{
					UserId:   userId,
					UserName: userName,
					Factor:   factor,
				}

			}

			for resp.HasNextPage() {
				var nextFactorSet []*okta.Factor
				resp, err = resp.Next(ctx, &nextFactorSet)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, factor := range nextFactorSet {
					resultChannel <- UserFactorInfo{
						UserId:   userId,
						UserName: userName,
						Factor:   *factor,
					}

				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, err)

		},
	}
}

type UserFactorInfo struct {
	UserId   string
	UserName string
	Factor   okta.Factor
}

func userProfile(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	user := result.(*okta.User)
	if user.Profile == nil {
		return nil, nil
	}
	userProfile := *user.Profile

	columnName := column.ColumnName
	if columnName == "title" {
		columnName = "login"
	}

	return userProfile[columnName].(string), nil
}

func (x *TableOktaFactorGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOktaFactorGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique key for Group.").
			Extractor(column_value_extractor.StructSelector("Factor.Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_name").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the user (username).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated").ColumnType(schema.ColumnTypeTimestamp).Description("The timestamp when the factor was last updated.").
			Extractor(okta_client.ExtractorTimestamp("Factor.LastUpdated")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provider").ColumnType(schema.ColumnTypeString).Description("The provider for the factor.").
			Extractor(column_value_extractor.StructSelector("Factor.Provider")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("embedded").ColumnType(schema.ColumnTypeJSON).Description("The Group's Profile properties.").
			Extractor(column_value_extractor.StructSelector("Factor.Embedded")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("The title of the resource.").
			Extractor(column_value_extractor.StructSelector("Factor.Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).Description("Unique key for Group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("factor_type").ColumnType(schema.ColumnTypeString).Description("Description of the Group.").
			Extractor(column_value_extractor.StructSelector("Factor.FactorType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when Group was created.").
			Extractor(okta_client.ExtractorTimestamp("Factor.Created")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The current status of the factor.").
			Extractor(column_value_extractor.StructSelector("Factor.Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("verify").ColumnType(schema.ColumnTypeJSON).Description("List of all users that are a member of this Group.").
			Extractor(column_value_extractor.StructSelector("Factor.Verify")).Build(),
	}
}

func (x *TableOktaFactorGenerator) GetSubTables() []*schema.Table {
	return nil
}
