package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/resourcemanager"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableOciResourcemanagerStackGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciResourcemanagerStackGenerator{}

func (x *TableOciResourcemanagerStackGenerator) GetTableName() string {
	return "oci_resourcemanager_stack"
}

func (x *TableOciResourcemanagerStackGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciResourcemanagerStackGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciResourcemanagerStackGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciResourcemanagerStackGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.ResourceManagerService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := resourcemanager.ListStacksRequest{
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.ResourceManagerClient.ListStacks(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, resource := range response.Items {
					resultChannel <- resource

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

func (x *TableOciResourcemanagerStackGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciResourcemanagerStackGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of the specified stack.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("terraform_version").ColumnType(schema.ColumnTypeString).Description("The version of Terraform specified for the stack.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_drift_last_checked").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time when the drift detection was last executed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("Human-readable display name for the stack.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stack_drift_status").ColumnType(schema.ColumnTypeString).Description("Drift status of the stack.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time when the stack was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("variables").ColumnType(schema.ColumnTypeJSON).Description("Terraform variables associated with this resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current lifecycle state of the stack.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("General description of the stack.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("config_source").ColumnType(schema.ColumnTypeJSON).Description("The version of Terraform specified for the stack.").Build(),
	}
}

func (x *TableOciResourcemanagerStackGenerator) GetSubTables() []*schema.Table {
	return nil
}
