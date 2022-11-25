package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/logging"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciLoggingLogGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciLoggingLogGenerator{}

func (x *TableOciLoggingLogGenerator) GetTableName() string {
	return "oci_logging_log"
}

func (x *TableOciLoggingLogGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciLoggingLogGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciLoggingLogGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciLoggingLogGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.LoggingManagementService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			logGroupId := task.ParentRawResult.(logging.LogGroupSummary).Id

			request := buildLoggingLogFilters()
			request.LogGroupId = logGroupId
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.LoggingManagementClient.ListLogs(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, log := range response.Items {
					resultChannel <- log

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

func buildLoggingLogFilters() logging.ListLogsRequest {
	request := logging.ListLogsRequest{}

	return request
}

func (x *TableOciLoggingLogGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciLoggingLogGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The log object state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("Time the resource was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_group_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the log group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_enabled").ColumnType(schema.ColumnTypeBool).Description("Whether or not this resource is currently enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("configuration").ColumnType(schema.ColumnTypeJSON).Description("Log object configuration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the log.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_type").ColumnType(schema.ColumnTypeString).Description("The logType that the log object is for, whether custom or service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_last_modified").ColumnType(schema.ColumnTypeTimestamp).Description("Time the resource was last modified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retention_duration").ColumnType(schema.ColumnTypeInt).Description("Log retention duration in 30-day increments (30, 60, 90 and so on).").Build(),
	}
}

func (x *TableOciLoggingLogGenerator) GetSubTables() []*schema.Table {
	return nil
}
