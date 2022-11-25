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

type TableOciCoreBootVolumeBackupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCoreBootVolumeBackupGenerator{}

func (x *TableOciCoreBootVolumeBackupGenerator) GetTableName() string {
	return "oci_core_boot_volume_backup"
}

func (x *TableOciCoreBootVolumeBackupGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCoreBootVolumeBackupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCoreBootVolumeBackupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCoreBootVolumeBackupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.CoreBlockStorageService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildBootVolumeBackupFilters()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.BlockstorageClient.ListBootVolumeBackups(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, backup := range response.Items {
					resultChannel <- backup

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

func buildBootVolumeBackupFilters() core.ListBootVolumeBackupsRequest {
	request := core.ListBootVolumeBackupsRequest{}

	return request
}

func (x *TableOciCoreBootVolumeBackupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciCoreBootVolumeBackupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("boot_volume_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the boot volume.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the boot volume backup was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size_in_gbs").ColumnType(schema.ColumnTypeInt).Description("The size of the boot volume, in GBs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name for the boot volume backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expiration_time").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the volume backup will expire and be automatically deleted.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_id").ColumnType(schema.ColumnTypeString).Description("The image OCID used to create the boot volume the backup is taken from.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of a volume backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_type").ColumnType(schema.ColumnTypeString).Description("Specifies whether the backup was created manually, or via scheduled backup policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_request_received").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the request to create the boot volume backup was received.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_boot_volume_backup_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the source boot volume backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("unique_size_in_gbs").ColumnType(schema.ColumnTypeInt).Description("The size used by the backup, in GBs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the boot volume backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of a boot volume backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Key Management master encryption assigned to the boot volume backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("system_tags").ColumnType(schema.ColumnTypeJSON).Description("System tags for this resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
	}
}

func (x *TableOciCoreBootVolumeBackupGenerator) GetSubTables() []*schema.Table {
	return nil
}
