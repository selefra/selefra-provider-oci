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

type TableOciCoreBlockVolumeReplicaGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCoreBlockVolumeReplicaGenerator{}

func (x *TableOciCoreBlockVolumeReplicaGenerator) GetTableName() string {
	return "oci_core_block_volume_replica"
}

func (x *TableOciCoreBlockVolumeReplicaGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCoreBlockVolumeReplicaGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCoreBlockVolumeReplicaGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCoreBlockVolumeReplicaGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.CoreBlockStorageService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := core.ListBlockVolumeReplicasRequest{
				AvailabilityDomain: pointer.ToStringPointer(taskClient.(*oci_client.OciClient).Zone),
				CompartmentId:      &session.TenancyID,
				Limit:              pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.BlockstorageClient.ListBlockVolumeReplicas(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, volumeReplica := range response.Items {
					resultChannel <- volumeReplica

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

func (x *TableOciCoreBlockVolumeReplicaGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildCompartementZonalList()
}

func (x *TableOciCoreBlockVolumeReplicaGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The block volume replica's Oracle ID (OCID).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size_in_gbs").ColumnType(schema.ColumnTypeInt).Description("The size of the source block volume, in GBs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of a block volume replica.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the block volume replica was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("block_volume_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the source block volume.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_domain").ColumnType(schema.ColumnTypeString).Description("The availability domain of the block volume replica.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_last_synced").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the block volume replica was last synced from the source block volume.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("total_data_transferred_in_gbs").ColumnType(schema.ColumnTypeInt).Description("The total size of the data transferred from the source block volume to the block volume replica, in GBs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
	}
}

func (x *TableOciCoreBlockVolumeReplicaGenerator) GetSubTables() []*schema.Table {
	return nil
}
