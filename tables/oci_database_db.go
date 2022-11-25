package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/database"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciDatabaseDbGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciDatabaseDbGenerator{}

func (x *TableOciDatabaseDbGenerator) GetTableName() string {
	return "oci_database_db"
}

func (x *TableOciDatabaseDbGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciDatabaseDbGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciDatabaseDbGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciDatabaseDbGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			homeId := task.ParentRawResult.(database.DbHomeSummary).Id

			session, err := oci_client.DatabaseService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := database.ListDatabasesRequest{
				CompartmentId: &session.TenancyID,
				DbHomeId:      homeId,
				Limit:         pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.DatabaseClient.ListDatabases(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, database := range response.Items {
					resultChannel <- database

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

func (x *TableOciDatabaseDbGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciDatabaseDbGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_name").ColumnType(schema.ColumnTypeString).Description("The database name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_software_image_id").ColumnType(schema.ColumnTypeString).Description("The database software image OCID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pdb_name").ColumnType(schema.ColumnTypeString).Description("The name of the pluggable database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ncharacter_set").ColumnType(schema.ColumnTypeString).Description("The national character set for the database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_database_point_in_time_recovery_timestamp").ColumnType(schema.ColumnTypeTimestamp).Description("Point in time recovery timeStamp of the source database at which cloned database system is cloned from the source database system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_unique_name").ColumnType(schema.ColumnTypeString).Description("A system-generated name for the database to ensure uniqueness within an oracle data guard group.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_home_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the database home.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_system_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_workload").ColumnType(schema.ColumnTypeString).Description("The database workload type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the database was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vm_cluster_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the vm cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connection_strings").ColumnType(schema.ColumnTypeJSON).Description("The connection strings used to connect to the oracle database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_backup_config").ColumnType(schema.ColumnTypeJSON).Description("Database backup configuration details.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("character_set").ColumnType(schema.ColumnTypeString).Description("The character set for the database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_backup_timestamp").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time when the latest database backup was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_details").ColumnType(schema.ColumnTypeString).Description("Additional information about the current lifecycle state.").Build(),
	}
}

func (x *TableOciDatabaseDbGenerator) GetSubTables() []*schema.Table {
	return nil
}
