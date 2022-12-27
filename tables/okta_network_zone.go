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

type TableOktaNetworkZoneGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOktaNetworkZoneGenerator{}

func (x *TableOktaNetworkZoneGenerator) GetTableName() string {
	return "okta_network_zone"
}

func (x *TableOktaNetworkZoneGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOktaNetworkZoneGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOktaNetworkZoneGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOktaNetworkZoneGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			client, err := okta_client.Connect(ctx, taskClient.(*okta_client.Client).Config)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			input := query.Params{
				Limit: 200,
			}

			networkZones, resp, err := client.NetworkZone.ListNetworkZones(ctx, &input)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, networkZone := range networkZones {
				resultChannel <- networkZone
			}

			for resp.HasNextPage() {
				var nextZoneSet []*okta.NetworkZone
				resp, err = resp.Next(ctx, &nextZoneSet)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, networkZone := range nextZoneSet {
					resultChannel <- networkZone
				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, err)
		},
	}
}

func (x *TableOktaNetworkZoneGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOktaNetworkZoneGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("created").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when the network zone was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("asns").ColumnType(schema.ColumnTypeJSON).Description("Format of each array value: a string representation of an ASN numeric value.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("proxies").ColumnType(schema.ColumnTypeJSON).Description("IP addresses (range or CIDR form) that are allowed to forward a request from gateway addresses. These proxies are automatically trusted by Threat Insights. These proxies are used to identify the client IP of a request.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Unique name for the zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_updated").ColumnType(schema.ColumnTypeTimestamp).Description("Timestamp when the network zone was last modified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("The title of the resource.").
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("system").ColumnType(schema.ColumnTypeBool).Description("Indicates if this is a system network zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gateways").ColumnType(schema.ColumnTypeJSON).Description("IP addresses (range or CIDR form) of the zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locations").ColumnType(schema.ColumnTypeJSON).Description("The geolocations of the zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Identifier of the network zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("proxy_type").ColumnType(schema.ColumnTypeString).Description("One of: '' or null (when not specified), Any (meaning any proxy), Tor, NotTorAnonymizer.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("Status of the network zone: ACTIVE or INACTIVE.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of the network zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("usage").ColumnType(schema.ColumnTypeString).Description("Usage of Zone: POLICY, BLOCKLIST.").Build(),
	}
}

func (x *TableOktaNetworkZoneGenerator) GetSubTables() []*schema.Table {
	return nil
}
