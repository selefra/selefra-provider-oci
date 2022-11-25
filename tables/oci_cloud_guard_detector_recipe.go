package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/cloudguard"
	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciCloudGuardDetectorRecipeGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCloudGuardDetectorRecipeGenerator{}

func (x *TableOciCloudGuardDetectorRecipeGenerator) GetTableName() string {
	return "oci_cloud_guard_detector_recipe"
}

func (x *TableOciCloudGuardDetectorRecipeGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCloudGuardDetectorRecipeGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCloudGuardDetectorRecipeGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCloudGuardDetectorRecipeGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.CloudGuardService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := cloudguard.ListDetectorRecipesRequest{
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.CloudGuardClient.ListDetectorRecipes(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, detectorRecipe := range response.Items {
					resultChannel <- detectorRecipe

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

func (x *TableOciCloudGuardDetectorRecipeGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildCompartmentList()
}

func (x *TableOciCloudGuardDetectorRecipeGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_updated").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the detector recipe was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("system_tags").ColumnType(schema.ColumnTypeJSON).Description("System tags for resource. System tags can be viewed by users, but can only be created by the system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Ocid for detector recipe.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the detector recipe.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("DisplayName of detector recipe.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the detector recipe was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeString).Description("Owner of detector recipe.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("detector").ColumnType(schema.ColumnTypeString).Description("Type of detector.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("detector_rules").ColumnType(schema.ColumnTypeJSON).Description("List of detector rules for the detector type for recipe.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("effective_detector_rules").ColumnType(schema.ColumnTypeJSON).Description("List of detector rules for the detector type for recipe.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_detector_recipe_id").ColumnType(schema.ColumnTypeString).Description("Recipe Ocid of the Source Recipe to be cloned.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("Detector recipe description.").Build(),
	}
}

func (x *TableOciCloudGuardDetectorRecipeGenerator) GetSubTables() []*schema.Table {
	return nil
}
