package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/database"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciDatabaseSoftwareImageGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciDatabaseSoftwareImageGenerator{}

func (x *TableOciDatabaseSoftwareImageGenerator) GetTableName() string {
	return "oci_database_software_image"
}

func (x *TableOciDatabaseSoftwareImageGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciDatabaseSoftwareImageGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciDatabaseSoftwareImageGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciDatabaseSoftwareImageGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.DatabaseService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildDatabaseSoftwareImageFilters()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.DatabaseClient.ListDatabaseSoftwareImages(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, image := range response.Items {
					resultChannel <- image

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

func buildDatabaseSoftwareImageFilters() database.ListDatabaseSoftwareImagesRequest {
	request := database.ListDatabaseSoftwareImagesRequest{}

	return request
}

func (x *TableOciDatabaseSoftwareImageGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciDatabaseSoftwareImageGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("included_patches_summary").ColumnType(schema.ColumnTypeString).Description("The patches included in the image and the version of the image.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The user-friendly name for the database software image. The name does not have to be unique.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the database software image.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the database software image was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_type").ColumnType(schema.ColumnTypeString).Description("The type of software image. It can be grid or database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the database software image.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_shape_family").ColumnType(schema.ColumnTypeString).Description("The shape that the image is meant for.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_details").ColumnType(schema.ColumnTypeString).Description("Detailed message for the lifecycle state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_upgrade_supported").ColumnType(schema.ColumnTypeBool).Description("True if this database software image is supported for upgrade.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_software_image_one_off_patches").ColumnType(schema.ColumnTypeJSON).Description("List of one-off patches for database homes.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_version").ColumnType(schema.ColumnTypeString).Description("The database version with which the database software image is to be built.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ls_inventory").ColumnType(schema.ColumnTypeString).Description("The output from lsinventory which will get passed as a string.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("patch_set").ColumnType(schema.ColumnTypeString).Description("The PSU or PBP or release updates.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_software_image_included_patches").ColumnType(schema.ColumnTypeJSON).Description("List of one-off patches for database homes.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
	}
}

func (x *TableOciDatabaseSoftwareImageGenerator) GetSubTables() []*schema.Table {
	return nil
}
