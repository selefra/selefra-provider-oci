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

type TableOciCoreSubnetGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCoreSubnetGenerator{}

func (x *TableOciCoreSubnetGenerator) GetTableName() string {
	return "oci_core_subnet"
}

func (x *TableOciCoreSubnetGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCoreSubnetGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCoreSubnetGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCoreSubnetGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.CoreVirtualNetworkService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildCoreSubnetFilters()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.VirtualNetworkClient.ListSubnets(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, subnet := range response.Items {
					resultChannel <- subnet

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

func buildCoreSubnetFilters() core.ListSubnetsRequest {
	request := core.ListSubnetsRequest{}

	return request
}

func (x *TableOciCoreSubnetGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciCoreSubnetGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The subnet's Oracle ID (OCID).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("route_table_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the route table that the subnet uses.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_virtual_router_ip").ColumnType(schema.ColumnTypeIp).Description("For an IPv6-enabled subnet, this is the IPv6 address of the virtual router.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("prohibit_public_ip_on_vnic").ColumnType(schema.ColumnTypeBool).Description("Indicates whether VNICs within this subnet can have public IP addresses.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dns_label").ColumnType(schema.ColumnTypeString).Description("A DNS label for the subnet, used in conjunction with the VNIC's hostname and VCN's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_list_ids").ColumnType(schema.ColumnTypeJSON).Description("The OCIDs of the security list or lists that the subnet uses.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vcn_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the VCN the subnet is in.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The subnet's current state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_domain").ColumnType(schema.ColumnTypeString).Description("The subnet's availability domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cidr_block").ColumnType(schema.ColumnTypeCIDR).Description("The subnet's CIDR block.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dhcp_options_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the set of DHCP options that the subnet uses.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the subnet was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_cidr_block").ColumnType(schema.ColumnTypeCIDR).Description("For an IPv6-enabled subnet, this is the IPv6 CIDR block for the subnet's private IP address space.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_public_cidr_block").ColumnType(schema.ColumnTypeCIDR).Description("For an IPv6-enabled subnet, this is the IPv6 CIDR block for the subnet's public IP address space.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("virtual_router_ip").ColumnType(schema.ColumnTypeIp).Description("The IP address of the virtual router.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name. Does not have to be unique, and it's changeable.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_domain_name").ColumnType(schema.ColumnTypeString).Description("The subnet's domain name, which consists of the subnet's DNS label, the VCN's DNS label, and the `oraclevcn.com` domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("virtual_router_mac").ColumnType(schema.ColumnTypeString).Description("The MAC address of the virtual router.").Build(),
	}
}

func (x *TableOciCoreSubnetGenerator) GetSubTables() []*schema.Table {
	return nil
}
