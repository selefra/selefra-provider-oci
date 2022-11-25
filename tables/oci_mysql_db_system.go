package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/mysql"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableOciMysqlDbSystemGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciMysqlDbSystemGenerator{}

func (x *TableOciMysqlDbSystemGenerator) GetTableName() string {
	return "oci_mysql_db_system"
}

func (x *TableOciMysqlDbSystemGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciMysqlDbSystemGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciMysqlDbSystemGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciMysqlDbSystemGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.MySQLDBSystemService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildMySQLDBSystemFilters()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.MySQLDBSystemClient.ListDbSystems(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, dbSystem := range response.Items {
					resultChannel <- dbSystem

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

func buildMySQLDBSystemFilters() mysql.ListDbSystemsRequest {
	request := mysql.ListDbSystemsRequest{}

	return request
}

func (x *TableOciMysqlDbSystemGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciMysqlDbSystemGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the DB System.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_domain").ColumnType(schema.ColumnTypeString).Description("The Availability Domain where the primary DB System should be located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_storage_size_in_gbs").ColumnType(schema.ColumnTypeInt).Description("Initial size of the data volume in GiBs that will be created and attached.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_address").ColumnType(schema.ColumnTypeIp).Description("The IP address the DB System is configured to listen on.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shape_name").ColumnType(schema.ColumnTypeString).Description("The shape of the primary instances of the DB System.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source").ColumnType(schema.ColumnTypeJSON).Description("DbSystemSource Parameters detailing how to provision the initial data of the DB System.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the DB System was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("User-provided data about the DB System.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup_policy").ColumnType(schema.ColumnTypeJSON).Description("BackupPolicy The Backup policy for the DB System.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The user-friendly name for the DB System. It does not have to be unique.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the subnet the DB System is associated with.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hostname_label").ColumnType(schema.ColumnTypeString).Description("The hostname for the primary endpoint of the DB System.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the DB System.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fault_domain").ColumnType(schema.ColumnTypeString).Description("The name of the fault domain the DB System is located in.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("analytics_cluster").ColumnType(schema.ColumnTypeJSON).Description("A summary of an Analytics Cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoints").ColumnType(schema.ColumnTypeJSON).Description("The network endpoints available for this DB System.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mysql_version").ColumnType(schema.ColumnTypeString).Description("Name of the MySQL Version in use for the DB System.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port_x").ColumnType(schema.ColumnTypeInt).Description("The network port on which X Plugin listens for TCP/IP connections. This is the X Plugin equivalent of port.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_analytics_cluster_attached").ColumnType(schema.ColumnTypeBool).Description("If the DB System has an Analytics Cluster attached.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_heat_wave_cluster_attached").ColumnType(schema.ColumnTypeBool).Description("Whether the DB System has a HeatWave cluster attached.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_details").ColumnType(schema.ColumnTypeString).Description("Additional information about the current lifecycleState.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("channels").ColumnType(schema.ColumnTypeJSON).Description("A list with a summary of all the Channels attached to the DB System.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("configuration_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Configuration to be used for Instances in this DB System.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_updated").ColumnType(schema.ColumnTypeTimestamp).Description("The time the DB System was last updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance").ColumnType(schema.ColumnTypeJSON).Description("The Maintenance Policy for the DB System.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port").ColumnType(schema.ColumnTypeInt).Description("The port for primary endpoint of the DB System to listen on.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
	}
}

func (x *TableOciMysqlDbSystemGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableOciMysqlDbSystemMetricConnectionsHourlyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciMysqlDbSystemMetricCpuUtilizationHourlyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciMysqlHeatWaveClusterGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciMysqlDbSystemMetricMemoryUtilizationDailyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciMysqlDbSystemMetricCpuUtilizationGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciMysqlDbSystemMetricCpuUtilizationDailyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciMysqlDbSystemMetricConnectionsGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciMysqlDbSystemMetricConnectionsDailyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciMysqlDbSystemMetricMemoryUtilizationGenerator{}),
	}
}
