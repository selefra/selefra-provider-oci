package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/streaming"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableOciStreamingStreamGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciStreamingStreamGenerator{}

func (x *TableOciStreamingStreamGenerator) GetTableName() string {
	return "oci_streaming_stream"
}

func (x *TableOciStreamingStreamGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciStreamingStreamGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciStreamingStreamGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciStreamingStreamGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.StreamAdminService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := streaming.ListStreamsRequest{
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToIntPointer(50),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.StreamAdminClient.ListStreams(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, stream := range response.Items {
					resultChannel <- stream

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

func (x *TableOciStreamingStreamGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciStreamingStreamGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the stream.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the stream was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("partitions").ColumnType(schema.ColumnTypeInt).Description("The number of partitions in the stream.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the stream.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state_details").ColumnType(schema.ColumnTypeString).Description("Any additional details about the current state of the stream.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("messages_endpoint").ColumnType(schema.ColumnTypeString).Description("The endpoint to use when creating the StreamClient to consume or publish messages in the stream.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the stream.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retention_in_hours").ColumnType(schema.ColumnTypeInt).Description("The retention period of the stream, in hours. This property is read-only.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stream_pool_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the stream pool that contains the stream.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
	}
}

func (x *TableOciStreamingStreamGenerator) GetSubTables() []*schema.Table {
	return nil
}
