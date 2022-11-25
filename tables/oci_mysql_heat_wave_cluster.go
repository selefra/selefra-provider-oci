package tables

import (
	"context"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/mysql"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableOciMysqlHeatWaveClusterGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciMysqlHeatWaveClusterGenerator{}

func (x *TableOciMysqlHeatWaveClusterGenerator) GetTableName() string {
	return "oci_mysql_heat_wave_cluster"
}

func (x *TableOciMysqlHeatWaveClusterGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciMysqlHeatWaveClusterGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciMysqlHeatWaveClusterGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciMysqlHeatWaveClusterGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			dbSystem := task.ParentRawResult.(mysql.DbSystemSummary)

			session, err := oci_client.MySQLDBSystemService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			if !*dbSystem.IsHeatWaveClusterAttached {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}

			request := mysql.GetHeatWaveClusterRequest{
				DbSystemId: dbSystem.Id,
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			response, err := session.MySQLDBSystemClient.GetHeatWaveCluster(ctx, request)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}
			resultChannel <- response.HeatWaveCluster

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableOciMysqlHeatWaveClusterGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciMysqlHeatWaveClusterGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the HeatWave cluster was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_size").ColumnType(schema.ColumnTypeInt).Description("The number of analytics-processing compute instances, of the specified shape, in the HeatWave cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_details").ColumnType(schema.ColumnTypeString).Description("Additional information about the current lifecycleState.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the HeatWave cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shape_name").ColumnType(schema.ColumnTypeString).Description("The shape determines resources to allocate to the HeatWave nodes - CPU cores, memory.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_updated").ColumnType(schema.ColumnTypeTimestamp).Description("The time the HeatWave cluster was last updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_nodes").ColumnType(schema.ColumnTypeJSON).Description("A HeatWave node is a compute host that is part of a HeatWave cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_system_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the parent DB System this HeatWave cluster is attached to.").Build(),
	}
}

func (x *TableOciMysqlHeatWaveClusterGenerator) GetSubTables() []*schema.Table {
	return nil
}
