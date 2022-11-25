package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/cloudguard"
	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
)

type TableOciCloudGuardTargetGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCloudGuardTargetGenerator{}

func (x *TableOciCloudGuardTargetGenerator) GetTableName() string {
	return "oci_cloud_guard_target"
}

func (x *TableOciCloudGuardTargetGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCloudGuardTargetGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCloudGuardTargetGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCloudGuardTargetGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.CloudGuardService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := cloudguard.ListTargetsRequest{
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.CloudGuardClient.ListTargets(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, target := range response.Items {
					resultChannel <- target

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

func (x *TableOciCloudGuardTargetGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildCompartmentList()
}

func (x *TableOciCloudGuardTargetGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("time_updated").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the target was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inherited_by_compartments").ColumnType(schema.ColumnTypeJSON).Description("List of inherited compartments.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_resource_id").ColumnType(schema.ColumnTypeString).Description("Resource ID which the target uses to monitor.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The target description.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("OCID for target.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the target was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_resource_type").ColumnType(schema.ColumnTypeString).Description("Possible type of targets(compartment/HCMCloud/ERPCloud).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecyle_details").ColumnType(schema.ColumnTypeString).Description("A message describing the current state in more detail.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recipe_count").ColumnType(schema.ColumnTypeInt).Description("Total number of recipes attached to target.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_detector_recipes").ColumnType(schema.ColumnTypeJSON).Description("List of detector recipes associated with target.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_responder_recipes").ColumnType(schema.ColumnTypeJSON).Description("List of responder recipes associated with target.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("system_tags").ColumnType(schema.ColumnTypeJSON).Description("System tags for resource. System tags can be viewed by users, but can only be created by the system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Target display name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
	}
}

func (x *TableOciCloudGuardTargetGenerator) GetSubTables() []*schema.Table {
	return nil
}
