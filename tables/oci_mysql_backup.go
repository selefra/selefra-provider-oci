package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/mysql"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciMysqlBackupGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciMysqlBackupGenerator{}

func (x *TableOciMysqlBackupGenerator) GetTableName() string {
	return "oci_mysql_backup"
}

func (x *TableOciMysqlBackupGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciMysqlBackupGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciMysqlBackupGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciMysqlBackupGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.MySQLBackupService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildMySQLBackupFilters()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.MySQLBackupClient.ListBackups(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, dbBackup := range response.Items {
					resultChannel <- dbBackup

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

func buildMySQLBackupFilters() mysql.ListBackupsRequest {
	request := mysql.ListBackupsRequest{}

	return request
}

func (x *TableOciMysqlBackupGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciMysqlBackupGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("mysql_version").ColumnType(schema.ColumnTypeString).Description("The version of the DB System used for backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retention_in_days").ColumnType(schema.ColumnTypeInt).Description("Number of days to retain this backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The time the backup record was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup_type").ColumnType(schema.ColumnTypeString).Description("The type of backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_storage_size_in_gbs").ColumnType(schema.ColumnTypeInt).Description("Initial size of the data volume in GiBs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_system_snapshot").ColumnType(schema.ColumnTypeJSON).Description("Snapshot of the DbSystem details at the time of the backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_updated").ColumnType(schema.ColumnTypeTimestamp).Description("The time at which the backup was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-supplied display name for the backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the Backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_type").ColumnType(schema.ColumnTypeString).Description("If the backup was created automatically, or by a manual request.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("A user-supplied description of the backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_system_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the DB System the Backup is associated with.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup_size_in_gbs").ColumnType(schema.ColumnTypeInt).Description("The size of the backup in GiBs.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_details").ColumnType(schema.ColumnTypeString).Description("Additional information about the current lifecycleState.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shape_name").ColumnType(schema.ColumnTypeString).Description("The shape of the DB System instance used for backup.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
	}
}

func (x *TableOciMysqlBackupGenerator) GetSubTables() []*schema.Table {
	return nil
}
