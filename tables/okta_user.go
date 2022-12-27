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
	"github.com/selefra/selefra-utils/pkg/reflect_util"
)

type TableOktaUserGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOktaUserGenerator{}

func (x *TableOktaUserGenerator) GetTableName() string {
	return "okta_user"
}

func (x *TableOktaUserGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOktaUserGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOktaUserGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOktaUserGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := okta_client.Connect(ctx, taskClient.(*okta_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			input := query.Params{
				Limit: 200,
			}

			users, resp, err := client.User.ListUsers(ctx, &input)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, user := range users {
				resultChannel <- user
			}

			for resp.HasNextPage() {
				var nextUserSet []*okta.User
				resp, err = resp.Next(ctx, &nextUserSet)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, user := range nextUserSet {
					resultChannel <- user
				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, err)
		},
	}
}

func listAssignedRolesForUser(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, error) {

	user := result.(*okta.User)
	client, err := okta_client.Connect(ctx, taskClient.(*okta_client.Client).Config)
	if err != nil {

		return nil, err
	}

	roles, resp, err := client.User.ListAssignedRolesForUser(ctx, user.Id, &query.Params{})
	if err != nil {

		return nil, err
	}

	for resp.HasNextPage() {
		var nextRolesSet []*okta.Role
		resp, err = resp.Next(ctx, &nextRolesSet)
		if err != nil {

			return nil, err
		}
		roles = append(roles, nextRolesSet...)
	}

	return roles, nil
}
func listUserGroups(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {

	user := result.(*okta.User)
	client, err := okta_client.Connect(ctx, taskClient.(*okta_client.Client).Config)
	if err != nil {
		return nil, err
	}

	groups, resp, err := client.User.ListUserGroups(ctx, user.Id)
	if err != nil {
		if strings.Contains(err.Error(), "Not found") {
			return nil, nil
		}
		return nil, err
	}

	for resp.HasNextPage() {
		var nextGroupSet []*okta.Group
		resp, err = resp.Next(ctx, &nextGroupSet)
		if err != nil {
			return nil, err
		}
		groups = append(groups, nextGroupSet...)
	}

	return groups, nil
}
func transformUserGroups(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (interface{}, error) {
	groups := result.([]*okta.Group)
	var groupsData = []map[string]string{}

	for _, group := range groups {
		groupsData = append(groupsData, map[string]string{
			"id":   group.Id,
			"name": group.Profile.Name,
			"type": group.Type,
		})
	}

	return groupsData, nil
}

func (x *TableOktaUserGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOktaUserGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).Description("A self-referential link to this user.").
			Extractor(column_value_extractor.StructSelector("Links.self.href")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_changed").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when status last changed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("profile").ColumnType(schema.ColumnTypeJSON).Description("User profile properties.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeJSON).Description("User type that determines the schema for the user's profile.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assigned_roles").ColumnType(schema.ColumnTypeJSON).Description("List of roles assigned to user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				// 003
				result, err := listAssignedRolesForUser(ctx, clientMeta, taskClient, task, row, column, result)

				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return result, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("activated").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when transition to ACTIVE status completed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when user was last updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transitioning_to_status").ColumnType(schema.ColumnTypeString).Description("Target status of an in-progress asynchronous status transition.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_groups").ColumnType(schema.ColumnTypeJSON).Description("List of groups of which the user is a member.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 002
				r, err := listUserGroups(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				if reflect_util.IsNil(r) {
					return nil, nil
				}

				r, err = transformUserGroups(ctx, clientMeta, taskClient, task, row, column, r)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}

				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("login").ColumnType(schema.ColumnTypeString).Description("Unique identifier for the user (username).").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 002
				r, err := userProfile(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("Primary email address of user.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 002
				r, err := userProfile(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when user was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password_changed").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when password last changed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("Current status of user. Can be one of the STAGED, PROVISIONED, ACTIVE, RECOVERY, LOCKED_OUT, PASSWORD_EXPIRED, SUSPENDED, or DEPROVISIONED.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("The title of the resource.").
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
				// 002
				r, err := userProfile(ctx, clientMeta, taskClient, task, row, column, result)
				if err != nil {
					return nil, schema.NewDiagnosticsErrorColumnValueExtractor(task.Table, column, err)
				}
				return r, nil
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique key for user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_login").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp of last login.").Build(),
	}
}

func (x *TableOktaUserGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableOktaFactorGenerator{}),
	}
}
