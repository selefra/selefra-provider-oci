package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/identity"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciIdentityNetworkSourceGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciIdentityNetworkSourceGenerator{}

func (x *TableOciIdentityNetworkSourceGenerator) GetTableName() string {
	return "oci_identity_network_source"
}

func (x *TableOciIdentityNetworkSourceGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciIdentityNetworkSourceGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciIdentityNetworkSourceGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciIdentityNetworkSourceGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.IdentityService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := identity.ListNetworkSourcesRequest{
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.IdentityClient.ListNetworkSources(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, networkSources := range response.Items {
					resultChannel <- networkSources

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

func (x *TableOciIdentityNetworkSourceGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOciIdentityNetworkSourceGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time the etwork source was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inactive_status").ColumnType(schema.ColumnTypeInt).Description("The detailed status of INACTIVE lifecycleState.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name you assign to the network source during creation.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the network source.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_source_list").ColumnType(schema.ColumnTypeJSON).Description("A list of allowed public IP addresses and CIDR ranges.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The network source object's current state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("services").ColumnType(schema.ColumnTypeJSON).Description("A list of services allowed to make on-behalf-of requests.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("virtual_source_list").ColumnType(schema.ColumnTypeJSON).Description("A list of allowed VCN OCID and IP range pairs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description you assign to the network source.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
	}
}

func (x *TableOciIdentityNetworkSourceGenerator) GetSubTables() []*schema.Table {
	return nil
}
