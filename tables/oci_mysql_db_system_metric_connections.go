package tables

import (
	"context"

	"github.com/oracle/oci-go-sdk/v44/mysql"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciMysqlDbSystemMetricConnectionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciMysqlDbSystemMetricConnectionsGenerator{}

func (x *TableOciMysqlDbSystemMetricConnectionsGenerator) GetTableName() string {
	return "oci_mysql_db_system_metric_connections"
}

func (x *TableOciMysqlDbSystemMetricConnectionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciMysqlDbSystemMetricConnectionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciMysqlDbSystemMetricConnectionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciMysqlDbSystemMetricConnectionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			dbSystem := task.ParentRawResult.(mysql.DbSystemSummary)
			region := taskClient.(*oci_client.OciClient).Region

			if dbSystem.LifecycleState == "DELETING" || dbSystem.LifecycleState == "DELETED" {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}

			_, err := listMonitoringMetricStatistics(ctx, clientMeta, taskClient, task, resultChannel, "5_MIN", "oci_mysql_database", "ActiveConnections", "resourceId", *dbSystem.Id, *dbSystem.CompartmentId, region)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			_, err = listMonitoringMetricStatistics(ctx, clientMeta, taskClient, task, resultChannel, "5_MIN", "oci_mysql_database", "CurrentConnections", "resourceId", *dbSystem.Id, *dbSystem.CompartmentId, region)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func (x *TableOciMysqlDbSystemMetricConnectionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciMysqlDbSystemMetricConnectionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metric_name").ColumnType(schema.ColumnTypeString).Description("The name of the metric.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).Description("The metric namespace.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("minimum").ColumnType(schema.ColumnTypeFloat).Description("The minimum metric value for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("unit").ColumnType(schema.ColumnTypeString).Description("The standard unit for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timestamp").ColumnType(schema.ColumnTypeTimestamp).Description("The time stamp used for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The ID of the compartment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the DB System.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("average").ColumnType(schema.ColumnTypeFloat).Description("The average of the metric values that correspond to the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maximum").ColumnType(schema.ColumnTypeFloat).Description("The maximum metric value for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sample_count").ColumnType(schema.ColumnTypeFloat).Description("The number of metric values that contributed to the aggregate value of this data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sum").ColumnType(schema.ColumnTypeFloat).Description("The sum of the metric values for the data point.").Build(),
	}
}

func (x *TableOciMysqlDbSystemMetricConnectionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
