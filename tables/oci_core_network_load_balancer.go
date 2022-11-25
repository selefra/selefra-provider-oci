package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/networkloadbalancer"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableOciCoreNetworkLoadBalancerGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCoreNetworkLoadBalancerGenerator{}

func (x *TableOciCoreNetworkLoadBalancerGenerator) GetTableName() string {
	return "oci_core_network_load_balancer"
}

func (x *TableOciCoreNetworkLoadBalancerGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCoreNetworkLoadBalancerGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCoreNetworkLoadBalancerGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCoreNetworkLoadBalancerGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.NetworkLoadBalancerService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := networkloadbalancer.ListNetworkLoadBalancersRequest{
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.NetworkLoadBalancerClient.ListNetworkLoadBalancers(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, networkLoadBalancer := range response.Items {
					resultChannel <- networkLoadBalancer

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

func (x *TableOciCoreNetworkLoadBalancerGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciCoreNetworkLoadBalancerGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_details").ColumnType(schema.ColumnTypeString).Description("A message describing the current state in more detail.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_security_group_ids").ColumnType(schema.ColumnTypeJSON).Description("An array of network security groups OCIDs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the network load balancer.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_updated").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the network load balancer was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the network load balancer was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_private").ColumnType(schema.ColumnTypeBool).Description("Whether the network load balancer has a virtual cloud network-local (private) IP address.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("listeners").ColumnType(schema.ColumnTypeJSON).Description("Listeners associated with the network load balancer.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("system_tags").ColumnType(schema.ColumnTypeJSON).Description("System tags for resource. System tags can be viewed by users, but can only be created by the system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name. Does not have to be unique.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the network load balancer.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_id").ColumnType(schema.ColumnTypeString).Description("The subnet in which the network load balancer is spawned OCIDs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_preserve_source_destination").ColumnType(schema.ColumnTypeBool).Description("When enabled, the skipSourceDestinationCheck parameter is automatically enabled on the load balancer VNIC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_load_balancer_health").ColumnType(schema.ColumnTypeJSON).Description("The overall health status of the network load balancer.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
	}
}

func (x *TableOciCoreNetworkLoadBalancerGenerator) GetSubTables() []*schema.Table {
	return nil
}
