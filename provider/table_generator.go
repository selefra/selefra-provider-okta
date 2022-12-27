package provider

import (
	"github.com/selefra/selefra-provider-okta/table_schema_generator"
	"github.com/selefra/selefra-provider-okta/tables"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableOktaSignonPolicyGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableOktaUserGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableOktaTrustedOriginGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableOktaPasswordPolicyGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableOktaMfaPolicyGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableOktaAuthServerGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableOktaNetworkZoneGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableOktaIdpDiscoveryPolicyGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableOktaFactorGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableOktaUserTypeGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableOktaGroupGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableOktaApplicationGenerator{}),
	}
}
