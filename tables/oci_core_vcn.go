package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/core"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciCoreVcnGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCoreVcnGenerator{}

func (x *TableOciCoreVcnGenerator) GetTableName() string {
	return "oci_core_vcn"
}

func (x *TableOciCoreVcnGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCoreVcnGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCoreVcnGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCoreVcnGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.CoreVirtualNetworkService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := core.ListVcnsRequest{
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.VirtualNetworkClient.ListVcns(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, network := range response.Items {
					resultChannel <- network

				}
				if response.OpcNextPage != nil {
					request.Page = response.OpcNextPage
				} else {
					pagesLeft = false
				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, err)

		},
	}
}

func (x *TableOciCoreVcnGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciCoreVcnGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the VCN was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cidr_block").ColumnType(schema.ColumnTypeCIDR).Description("The first CIDR IP address from cidrBlocks.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_dhcp_options_id").ColumnType(schema.ColumnTypeString).Description("The OCID for the VCN's default set of DHCP options.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_security_list_id").ColumnType(schema.ColumnTypeString).Description("The OCID for the VCN's default security list.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cidr_blocks").ColumnType(schema.ColumnTypeJSON).Description("The list of IPv4 CIDR blocks the VCN will use.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_cidr_blocks").ColumnType(schema.ColumnTypeJSON).Description("For an IPv6-enabled VCN, this is the list of IPv6 CIDR blocks for the VCN's IP address space. The CIDRs are provided by Oracle and the sizes are always /56.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The VCN's current state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vcn_domain_name").ColumnType(schema.ColumnTypeString).Description("The VCN's domain name, which consists of the VCN's DNS label, and the oraclevcn.com domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name. Does not have to be unique, and it's changeable.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The VCN's Oracle ID (OCID).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_route_table_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dns_label").ColumnType(schema.ColumnTypeString).Description("A DNS label for the VCN, used in conjunction with the VNIC's hostname and subnet's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_cidr_block").ColumnType(schema.ColumnTypeCIDR).Description("For an IPv6-enabled VCN, this is the IPv6 CIDR block for the VCN's private IP address space.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_public_cidr_block").ColumnType(schema.ColumnTypeString).Description("For an IPv6-enabled VCN, this is the IPv6 CIDR block for the VCN's public IP address space.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
	}
}

func (x *TableOciCoreVcnGenerator) GetSubTables() []*schema.Table {
	return nil
}
