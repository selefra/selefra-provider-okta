package tables

import (
	"context"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/okta/okta-sdk-golang/v2/okta/query"
	"github.com/selefra/selefra-provider-okta/okta_client"
	"github.com/selefra/selefra-provider-okta/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-utils/pkg/reflect_util"
)

type TableOktaGroupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOktaGroupGenerator{}

func (x *TableOktaGroupGenerator) GetTableName() string {
	return "okta_group"
}

func (x *TableOktaGroupGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOktaGroupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOktaGroupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOktaGroupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := okta_client.Connect(ctx, taskClient.(*okta_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			input := query.Params{
				Limit: 10000,
			}

			groups, resp, err := client.Group.ListGroups(ctx, &input)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, group := range groups {
				resultChannel <- group
			}

			for resp.HasNextPage() {
				var nextGroupSet []*okta.Group
				resp, err = resp.Next(ctx, &nextGroupSet)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, group := range nextGroupSet {
					resultChannel <- group

				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, err)

		},
	}
}

var filterTimeFormat = "2006-01-02T15:04:05.000Z"
var operatorsMap = map[string]string{
	"=":  "eq",
	">=": "ge",
	">":  "gt",
	"<=": "le",
	"<":  "lt",
	"<>": "ne",
}

func listGroupMembers(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	groupId := result.(*okta.Group).Id

	client, err := okta_client.Connect(ctx, taskClient.(*okta_client.Client).Config)
	if err != nil {
		return nil, err
	}

	groupMembers, resp, err := client.Group.ListGroupUsers(ctx, groupId, &query.Params{})
	if err != nil {
		return nil, err
	}

	for resp.HasNextPage() {
		var nextgroupMembersSet []*okta.User
		resp, err = resp.Next(ctx, &groupMembers)
		if err != nil {
			return nil, err
		}
		groupMembers = append(groupMembers, nextgroupMembersSet...)
	}

	return groupMembers, nil
}

func transformGroupMembers(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	users := result.([]*okta.User)
	var usersData = []map[string]string{}

	for _, user := range users {
		userProfile := *user.Profile
		usersData = append(usersData, map[string]string{
			"id":    user.Id,
			"email": userProfile["email"].(string),
			"login": userProfile["login"].(string),
		})
	}

	return usersData, nil
}

func (x *TableOktaGroupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOktaGroupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("object_class").ColumnType(schema.ColumnTypeJSON).Description("Determines the Group's profile.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("The title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Name of the Group.").
			Extractor(column_value_extractor.StructSelector("Profile.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique key for Group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when Group was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_membership_updated").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when Group's memberships were last updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("Determines how a Group's Profile and memberships are managed. Can be one of OKTA_GROUP, APP_GROUP or BUILT_IN.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("Description of the Group.").
			Extractor(column_value_extractor.StructSelector("Profile.Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when Group's profile was last updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("profile").ColumnType(schema.ColumnTypeJSON).Description("The Group's Profile properties.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group_members").ColumnType(schema.ColumnTypeJSON).Description("List of all users that are a member of this Group.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 002
				r, err := listGroupMembers(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				if reflect_util.IsNil(r) {
					return nil, nil
				}

				r, err = transformGroupMembers(ctx, clientMeta, taskClient, task, row, column, r)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return r, nil
			})).Build(),
	}
}

func (x *TableOktaGroupGenerator) GetSubTables() []*schema.Table {
	return nil
}
