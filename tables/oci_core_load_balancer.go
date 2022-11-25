package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/loadbalancer"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciCoreLoadBalancerGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCoreLoadBalancerGenerator{}

func (x *TableOciCoreLoadBalancerGenerator) GetTableName() string {
	return "oci_core_load_balancer"
}

func (x *TableOciCoreLoadBalancerGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCoreLoadBalancerGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCoreLoadBalancerGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCoreLoadBalancerGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.LoadBalancerService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := loadbalancer.ListLoadBalancersRequest{
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToInt64Pointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.LoadBalancerClient.ListLoadBalancers(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, loadBalancer := range response.Items {
					resultChannel <- loadBalancer

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

func (x *TableOciCoreLoadBalancerGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciCoreLoadBalancerGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the load balancer.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name of the load balancer.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backend_sets").ColumnType(schema.ColumnTypeJSON).Description("The configuration of a load balancer backend set.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hostnames").ColumnType(schema.ColumnTypeJSON).Description("A hostname resource associated with a load balancer for use by one or more listeners.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ssl_cipher_suites").ColumnType(schema.ColumnTypeJSON).Description("The configuration details of an SSL cipher suite.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_ids").ColumnType(schema.ColumnTypeJSON).Description("An array of subnet OCIDs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shape_name").ColumnType(schema.ColumnTypeString).Description("A template that determines the total pre-provisioned bandwidth (ingress plus egress).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_private").ColumnType(schema.ColumnTypeBool).Description("Whether the load balancer has a VCN-local (private) IP address.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_addresses").ColumnType(schema.ColumnTypeJSON).Description("An array of IP addresses.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_security_group_ids").ColumnType(schema.ColumnTypeJSON).Description("An array of NSG OCIDs associated with the load balancer.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shape_details").ColumnType(schema.ColumnTypeJSON).Description("The configuration details to update load balancer to a different shape.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("system_tags").ColumnType(schema.ColumnTypeJSON).Description("System tags for this resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The load balancer's current state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the load balancer was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("path_route_sets").ColumnType(schema.ColumnTypeJSON).Description("A named set of path route rules.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rule_sets").ColumnType(schema.ColumnTypeJSON).Description("A named set of rules associated with a load balancer.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificates").ColumnType(schema.ColumnTypeJSON).Description("The configuration details of a certificate bundle.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("listeners").ColumnType(schema.ColumnTypeJSON).Description("The listener's configuration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("routing_policies").ColumnType(schema.ColumnTypeJSON).Description("A named ordered list of routing rules that is applied to a listener.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
	}
}

func (x *TableOciCoreLoadBalancerGenerator) GetSubTables() []*schema.Table {
	return nil
}
