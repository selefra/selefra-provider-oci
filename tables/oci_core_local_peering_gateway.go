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

type TableOciCoreLocalPeeringGatewayGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCoreLocalPeeringGatewayGenerator{}

func (x *TableOciCoreLocalPeeringGatewayGenerator) GetTableName() string {
	return "oci_core_local_peering_gateway"
}

func (x *TableOciCoreLocalPeeringGatewayGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCoreLocalPeeringGatewayGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCoreLocalPeeringGatewayGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCoreLocalPeeringGatewayGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.CoreVirtualNetworkService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := core.ListLocalPeeringGatewaysRequest{
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				gateways, err := session.VirtualNetworkClient.ListLocalPeeringGateways(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, gateway := range gateways.Items {
					resultChannel <- gateway

				}
				if gateways.OpcNextPage != nil {
					request.Page = gateways.OpcNextPage
				} else {
					pagesLeft = false
				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, err)

		},
	}
}

func (x *TableOciCoreLocalPeeringGatewayGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciCoreLocalPeeringGatewayGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name. Does not have to be unique, and it's changeable.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vcn_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the VCN that uses the LPG.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("peer_advertised_cidr").ColumnType(schema.ColumnTypeCIDR).Description("The smallest aggregate CIDR that contains all the CIDR routes advertised by the VCN at the other end of the peering from this LPG.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("peer_advertised_cidr_details").ColumnType(schema.ColumnTypeJSON).Description("The specific ranges of IP addresses available on or via the VCN at the other end of the peering from this LPG.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("peering_status").ColumnType(schema.ColumnTypeString).Description("Whether the LPG is peered with another LPG.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("peering_status_details").ColumnType(schema.ColumnTypeString).Description("Additional information regarding the peering status.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The LPG's Oracle ID").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The LPG's current lifecycle state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the LPG was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("route_table_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the route table the LPG is using.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_cross_tenancy_peering").ColumnType(schema.ColumnTypeBool).Description("Whether the VCN at the other end of the peering is in a different tenancy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
	}
}

func (x *TableOciCoreLocalPeeringGatewayGenerator) GetSubTables() []*schema.Table {
	return nil
}
