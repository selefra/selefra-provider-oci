package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/dns"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciDnsZoneGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciDnsZoneGenerator{}

func (x *TableOciDnsZoneGenerator) GetTableName() string {
	return "oci_dns_zone"
}

func (x *TableOciDnsZoneGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciDnsZoneGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciDnsZoneGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciDnsZoneGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.DnsService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildDnsZoneFilters()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToInt64Pointer(50)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				zones, err := session.DnsClient.ListZones(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, zone := range zones.Items {
					resultChannel <- zone

				}
				if zones.OpcNextPage != nil {
					request.Page = zones.OpcNextPage
				} else {
					pagesLeft = false
				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, err)

		},
	}
}

func buildDnsZoneFilters() dns.ListZonesRequest {
	request := dns.ListZonesRequest{}

	return request
}

func (x *TableOciDnsZoneGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildCompartmentList()
}

func (x *TableOciDnsZoneGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("is_protected").ColumnType(schema.ColumnTypeBool).Description("A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nameservers").ColumnType(schema.ColumnTypeJSON).Description("The authoritative nameservers for the zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self").ColumnType(schema.ColumnTypeString).Description("The canonical absolute URL of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the zone resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the zone was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("serial").ColumnType(schema.ColumnTypeInt).Description("The current serial of the zone. As seen in the zone's SOA record.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeString).Description("Version is the never-repeating, totally-orderable, version of the zone, from which the serial field of the zone's SOA record is derived.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_type").ColumnType(schema.ColumnTypeString).Description("The type of the zone. Must be either `PRIMARY` or `SECONDARY`. `SECONDARY` is only supported for GLOBAL zones.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scope").ColumnType(schema.ColumnTypeString).Description("The scope of the zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("view_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the private view containing the zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external_masters").ColumnType(schema.ColumnTypeJSON).Description("External master servers for the zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
	}
}

func (x *TableOciDnsZoneGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableOciDnsRrsetGenerator{}),
	}
}
