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

type TableOciCoreBootVolumeGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCoreBootVolumeGenerator{}

func (x *TableOciCoreBootVolumeGenerator) GetTableName() string {
	return "oci_core_boot_volume"
}

func (x *TableOciCoreBootVolumeGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCoreBootVolumeGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCoreBootVolumeGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCoreBootVolumeGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.CoreBlockStorageService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := core.ListBootVolumesRequest{
				CompartmentId:      &session.TenancyID,
				AvailabilityDomain: pointer.ToStringPointer(taskClient.(*oci_client.OciClient).Zone),
				Limit:              pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.BlockstorageClient.ListBootVolumes(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, volume := range response.Items {
					resultChannel <- volume

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

func (x *TableOciCoreBootVolumeGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildCompartementZonalList()
}

func (x *TableOciCoreBootVolumeGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Key Management key which is the master encryption key for the boot volume.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size_in_mbs").ColumnType(schema.ColumnTypeInt).Description("The size of the boot volume in MBs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("volume_backup_policy_assignment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the volume backup policy assignment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("system_tags").ColumnType(schema.ColumnTypeJSON).Description("System tags for resource. System tags can be viewed by users, but can only be created by the system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("volume_group_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the source volume group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_domain").ColumnType(schema.ColumnTypeString).Description("The availability domain of the boot volume.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_hydrated").ColumnType(schema.ColumnTypeBool).Description("Specifies whether the boot volume's data has finished copying from the source boot volume or boot volume backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpus_per_gb").ColumnType(schema.ColumnTypeInt).Description("The number of volume performance units (VPUs) that will be applied to this boot volume per GB,representing the Block Volume service's elastic performance options.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_tuned_vpus_per_gb").ColumnType(schema.ColumnTypeInt).Description("The number of Volume Performance Units per GB that this boot volume is effectively tuned to when it's idle.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_auto_tune_enabled").ColumnType(schema.ColumnTypeBool).Description("Specifies whether the auto-tune performance is enabled for this boot volume.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the boot volume.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the boot volume was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("volume_backup_policy_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the volume backup policy that has been assigned to the volume.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of a boot volume.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size_in_gbs").ColumnType(schema.ColumnTypeInt).Description("The size of the boot volume in GBs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_details").ColumnType(schema.ColumnTypeJSON).Description("The boot volume source, either an existing volume in the same availability domain or a volume backup.").Build(),
	}
}

func (x *TableOciCoreBootVolumeGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableOciCoreBootVolumeMetricWriteOpsDailyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciCoreBootVolumeMetricWriteOpsHourlyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciCoreBootVolumeMetricReadOpsHourlyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciCoreBootVolumeMetricReadOpsDailyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciCoreBootVolumeMetricWriteOpsGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciCoreBootVolumeMetricReadOpsGenerator{}),
	}
}
