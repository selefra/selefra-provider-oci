package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/core"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableOciCoreVolumeBackupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCoreVolumeBackupGenerator{}

func (x *TableOciCoreVolumeBackupGenerator) GetTableName() string {
	return "oci_core_volume_backup"
}

func (x *TableOciCoreVolumeBackupGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCoreVolumeBackupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCoreVolumeBackupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCoreVolumeBackupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.CoreBlockStorageService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildCoreVolumeBackupFilters()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.BlockstorageClient.ListVolumeBackups(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, volumeBackups := range response.Items {
					resultChannel <- volumeBackups

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

func buildCoreVolumeBackupFilters() core.ListVolumeBackupsRequest {
	request := core.ListVolumeBackupsRequest{}

	return request
}

func (x *TableOciCoreVolumeBackupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciCoreVolumeBackupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("size_in_gbs").ColumnType(schema.ColumnTypeInt).Description("The size of the volume, in GBs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("unique_size_in_mbs").ColumnType(schema.ColumnTypeInt).Description("The size used by the backup, in MBs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("volume_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the volume.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expiration_time").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the volume backup will expire and be automatically deleted.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Key Management key which is the master encryption key for the volume backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size_in_mbs").ColumnType(schema.ColumnTypeInt).Description("The size of the volume in MBs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_volume_backup_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the source volume backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the volume backup was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("system_tags").ColumnType(schema.ColumnTypeJSON).Description("System tags to volume by the service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the volume backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name for the volume backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of a volume backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_type").ColumnType(schema.ColumnTypeString).Description("Specifies whether the backup was created manually, or via scheduled backup policy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of a volume backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_request_received").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the request to create the volume backup was received.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("unique_size_in_gbs").ColumnType(schema.ColumnTypeInt).Description("The size used by the backup, in GBs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
	}
}

func (x *TableOciCoreVolumeBackupGenerator) GetSubTables() []*schema.Table {
	return nil
}
