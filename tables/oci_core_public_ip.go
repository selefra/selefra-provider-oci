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

type TableOciCorePublicIpGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCorePublicIpGenerator{}

func (x *TableOciCorePublicIpGenerator) GetTableName() string {
	return "oci_core_public_ip"
}

func (x *TableOciCorePublicIpGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCorePublicIpGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCorePublicIpGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCorePublicIpGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.CoreVirtualNetworkService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildCorePublicIPFilters()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}
			request.Scope = "REGION"

			pagesLeft := true
			for pagesLeft {
				response, err := session.VirtualNetworkClient.ListPublicIps(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, ip := range response.Items {
					resultChannel <- ip

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

func buildCorePublicIPFilters() core.ListPublicIpsRequest {
	request := core.ListPublicIpsRequest{}

	return request
}

func (x *TableOciCorePublicIpGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciCorePublicIpGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assigned_entity_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the entity the public IP is assigned to, or in the process of being assigned to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assigned_entity_type").ColumnType(schema.ColumnTypeString).Description("The type of entity the public IP is assigned to, or in the process of being assigned to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the public IP was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_domain").ColumnType(schema.ColumnTypeString).Description("The public IP's availability domain. This property is set only for ephemeral public IPs that are assigned to a private IP.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_address").ColumnType(schema.ColumnTypeIp).Description("The public IP address of the publicIp object.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_ip_pool_id").ColumnType(schema.ColumnTypeString).Description("The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pool object created in the current tenancy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The public IP's Oracle ID (OCID).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The public IP's current state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scope").ColumnType(schema.ColumnTypeString).Description("Indicates whether the public IP is regional or specific to a particular availability domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifetime").ColumnType(schema.ColumnTypeString).Description("Defines when the public IP is deleted and released back to Oracle's public IP pool.").Build(),
	}
}

func (x *TableOciCorePublicIpGenerator) GetSubTables() []*schema.Table {
	return nil
}
