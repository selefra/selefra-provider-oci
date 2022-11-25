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

type TableOciCoreVolumeAttachmentGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCoreVolumeAttachmentGenerator{}

func (x *TableOciCoreVolumeAttachmentGenerator) GetTableName() string {
	return "oci_core_volume_attachment"
}

func (x *TableOciCoreVolumeAttachmentGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCoreVolumeAttachmentGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCoreVolumeAttachmentGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCoreVolumeAttachmentGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.CoreComputeService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildCoreVolumeAttachmentFilters()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.ComputeClient.ListVolumeAttachments(ctx, request)
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

func buildCoreVolumeAttachmentFilters() core.ListVolumeAttachmentsRequest {
	request := core.ListVolumeAttachmentsRequest{}

	return request
}

func (x *TableOciCoreVolumeAttachmentGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciCoreVolumeAttachmentGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the volume was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iscsi_login_state").ColumnType(schema.ColumnTypeString).Description("The iscsi login state of the volume attachment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_read_only").ColumnType(schema.ColumnTypeBool).Description("Whether the attachment was created in read-only mode.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_pv_encryption_in_transit_enabled").ColumnType(schema.ColumnTypeBool).Description("Whether in-transit encryption for the data volume's paravirtualized attachment is enabled or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the volume attachment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("volume_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the volume.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_domain").ColumnType(schema.ColumnTypeString).Description("The availability domain of an instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("device").ColumnType(schema.ColumnTypeString).Description("The device name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_shareable").ColumnType(schema.ColumnTypeBool).Description("Whether the attachment should be created in shareable mode.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_multipath").ColumnType(schema.ColumnTypeBool).Description("Whether the attachment is multipath or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the instance the volume is attached to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the volume attachment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name. Does not have to be unique, and it cannot be changed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
	}
}

func (x *TableOciCoreVolumeAttachmentGenerator) GetSubTables() []*schema.Table {
	return nil
}
