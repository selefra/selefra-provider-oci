package tables

import (
	"context"
	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-utils/pkg/pointer"
)

type TableOciCoreImageGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCoreImageGenerator{}

func (x *TableOciCoreImageGenerator) GetTableName() string {
	return "oci_core_image"
}

func (x *TableOciCoreImageGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCoreImageGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCoreImageGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCoreImageGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.CoreComputeService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildImageFilter()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.ComputeClient.ListImages(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, image := range response.Items {
					if image.BaseImageId == nil {
						resultChannel <- image

					}
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

func (x *TableOciCoreImageGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciCoreImageGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("operating_system").ColumnType(schema.ColumnTypeString).Description("The image's operating system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("agent_features").ColumnType(schema.ColumnTypeJSON).Description("Oracle Cloud Agent features supported on the image.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name for the image. It does not have to be unique, and it's changeable.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the image.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_image_allowed").ColumnType(schema.ColumnTypeBool).Description("Indicates whether instances launched with this image can be used to create new images.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operating_system_version").ColumnType(schema.ColumnTypeString).Description("The image's operating system version.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size_in_mbs").ColumnType(schema.ColumnTypeInt).Description("The boot volume size for an instance launched from this image.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The image's current state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the image was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("launch_mode").ColumnType(schema.ColumnTypeString).Description("Specifies the configuration mode for launching virtual machine (VM) instances.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("launch_options").ColumnType(schema.ColumnTypeJSON).Description("LaunchOptions Options for tuning the compatibility and performance of VM shapes.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("base_image_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the image originally used to launch the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
	}
}

func (x *TableOciCoreImageGenerator) GetSubTables() []*schema.Table {
	return nil
}
