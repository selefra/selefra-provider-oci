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

type TableOciCoreBootVolumeAttachmentGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCoreBootVolumeAttachmentGenerator{}

func (x *TableOciCoreBootVolumeAttachmentGenerator) GetTableName() string {
	return "oci_core_boot_volume_attachment"
}

func (x *TableOciCoreBootVolumeAttachmentGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCoreBootVolumeAttachmentGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCoreBootVolumeAttachmentGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCoreBootVolumeAttachmentGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.CoreComputeService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := core.ListBootVolumeAttachmentsRequest{
				AvailabilityDomain: pointer.ToStringPointer(taskClient.(*oci_client.OciClient).Zone),
				CompartmentId:      &session.TenancyID,
				Limit:              pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.ComputeClient.ListBootVolumeAttachments(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, volumeAttachment := range response.Items {
					resultChannel <- volumeAttachment

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

func (x *TableOciCoreBootVolumeAttachmentGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildCompartementZonalList()
}

func (x *TableOciCoreBootVolumeAttachmentGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_in_transit_type").ColumnType(schema.ColumnTypeString).Description("The type of the encryption in transit for the boot volume.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the boot volume attachment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_domain").ColumnType(schema.ColumnTypeString).Description("The availability domain of an instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the boot volume was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_pv_encryption_in_transit_enabled").ColumnType(schema.ColumnTypeBool).Description("Whether in-transit encryption for the boot volume's paravirtualized attachment is enabled or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the boot volume attachment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name. Does not have to be unique, and it cannot be changed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("boot_volume_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the boot volume.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the instance the boot volume is attached to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
	}
}

func (x *TableOciCoreBootVolumeAttachmentGenerator) GetSubTables() []*schema.Table {
	return nil
}
