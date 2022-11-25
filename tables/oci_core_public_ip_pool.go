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

type TableOciCorePublicIpPoolGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCorePublicIpPoolGenerator{}

func (x *TableOciCorePublicIpPoolGenerator) GetTableName() string {
	return "oci_core_public_ip_pool"
}

func (x *TableOciCorePublicIpPoolGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCorePublicIpPoolGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCorePublicIpPoolGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCorePublicIpPoolGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.CoreVirtualNetworkService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := core.ListPublicIpPoolsRequest{
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.VirtualNetworkClient.ListPublicIpPools(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, ipPool := range response.Items {
					resultChannel <- ipPool

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

func (x *TableOciCorePublicIpPoolGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciCorePublicIpPoolGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the public IP pool.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The public IP pool's current state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name. Does not have to be unique.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the public IP pool was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cidr_blocks").ColumnType(schema.ColumnTypeJSON).Description("The CIDR blocks added to this pool. This could be all or a portion of a BYOIP CIDR block.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
	}
}

func (x *TableOciCorePublicIpPoolGenerator) GetSubTables() []*schema.Table {
	return nil
}
