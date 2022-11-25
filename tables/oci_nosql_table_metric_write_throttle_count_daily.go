package tables

import (
	"context"

	"github.com/oracle/oci-go-sdk/v44/nosql"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciNosqlTableMetricWriteThrottleCountDailyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciNosqlTableMetricWriteThrottleCountDailyGenerator{}

func (x *TableOciNosqlTableMetricWriteThrottleCountDailyGenerator) GetTableName() string {
	return "oci_nosql_table_metric_write_throttle_count_daily"
}

func (x *TableOciNosqlTableMetricWriteThrottleCountDailyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciNosqlTableMetricWriteThrottleCountDailyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciNosqlTableMetricWriteThrottleCountDailyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciNosqlTableMetricWriteThrottleCountDailyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			table := task.ParentRawResult.(nosql.TableSummary)
			region := taskClient.(*oci_client.OciClient).Region

			_, err := listMonitoringMetricStatistics(ctx, clientMeta, taskClient, task, resultChannel, "DAILY", "oci_nosql", "WriteThrottleCount", "tableName", *table.Name, *table.CompartmentId, region)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func (x *TableOciNosqlTableMetricWriteThrottleCountDailyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciNosqlTableMetricWriteThrottleCountDailyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("metric_name").ColumnType(schema.ColumnTypeString).Description("The name of the metric.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maximum").ColumnType(schema.ColumnTypeFloat).Description("The maximum metric value for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sample_count").ColumnType(schema.ColumnTypeFloat).Description("The number of metric values that contributed to the aggregate value of this data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sum").ColumnType(schema.ColumnTypeFloat).Description("The sum of the metric values for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("unit").ColumnType(schema.ColumnTypeString).Description("The standard unit for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timestamp").ColumnType(schema.ColumnTypeTimestamp).Description("The time stamp used for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The ID of the compartment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the NoSQL table.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).Description("The metric namespace.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("average").ColumnType(schema.ColumnTypeFloat).Description("The average of the metric values that correspond to the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("minimum").ColumnType(schema.ColumnTypeFloat).Description("The minimum metric value for the data point.").Build(),
	}
}

func (x *TableOciNosqlTableMetricWriteThrottleCountDailyGenerator) GetSubTables() []*schema.Table {
	return nil
}
