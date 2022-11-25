package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/database"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
)

type TableOciDatabaseAutonomousDatabaseGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciDatabaseAutonomousDatabaseGenerator{}

func (x *TableOciDatabaseAutonomousDatabaseGenerator) GetTableName() string {
	return "oci_database_autonomous_database"
}

func (x *TableOciDatabaseAutonomousDatabaseGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciDatabaseAutonomousDatabaseGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciDatabaseAutonomousDatabaseGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciDatabaseAutonomousDatabaseGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.DatabaseService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildAutonomousDatabaseFilter()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.DatabaseClient.ListAutonomousDatabases(ctx, request)
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

func buildAutonomousDatabaseFilter() database.ListAutonomousDatabasesRequest {
	request := database.ListAutonomousDatabasesRequest{}
	return request
}

func (x *TableOciDatabaseAutonomousDatabaseGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciDatabaseAutonomousDatabaseGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("infrastructure_type").ColumnType(schema.ColumnTypeString).Description("The infrastructure type this resource belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permission_level").ColumnType(schema.ColumnTypeString).Description("The Autonomous Database permission level.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("system_tags").ColumnType(schema.ColumnTypeJSON).Description("System tags for resource. System tags can be viewed by users, but can only be created by the system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_auto_scaling_enabled").ColumnType(schema.ColumnTypeBool).Description("Indicates if auto scaling is enabled for the Autonomous Database CPU core count.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("refreshable_status").ColumnType(schema.ColumnTypeString).Description("The refresh status of the clone. REFRESHING indicates that the clone is currently being refreshed with data from the source Autonomous Database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_of_last_switchover").ColumnType(schema.ColumnTypeTimestamp).Description("The timestamp of the last switchover operation for the Autonomous Database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("whitelisted_ips").ColumnType(schema.ColumnTypeJSON).Description("The client IP access control list (ACL).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Autonomous Database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_dedicated").ColumnType(schema.ColumnTypeBool).Description("True if the database uses dedicated Exadata infrastructure.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("license_model").ColumnType(schema.ColumnTypeString).Description("The Oracle license model that applies to the Oracle Autonomous Database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_of_last_refresh").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time when last refresh happened.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("apex_details").ColumnType(schema.ColumnTypeJSON).Description("Information about Oracle APEX Application Development.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("available_upgrade_versions").ColumnType(schema.ColumnTypeJSON).Description("List of Oracle Database versions available for a database upgrade. If there are no version upgrades available, this list is empty.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cpu_core_count").ColumnType(schema.ColumnTypeInt).Description("The number of OCPU cores to be made available to the database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_access_control_enabled").ColumnType(schema.ColumnTypeBool).Description("Indicates if the database-level access control is enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("open_mode").ColumnType(schema.ColumnTypeString).Description("The `DATABASE OPEN` mode. You can open the database in `READ_ONLY` or `READ_WRITE` mode.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the source Autonomous Database that was cloned to create the current Autonomous Database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_maintenance_begin").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time when maintenance will begin.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_storage_size_in_gbs").ColumnType(schema.ColumnTypeInt).Description("The quantity of data in the database, in gigabytes.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_refreshable_clone").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the Autonomous Database is a refreshable clone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup_config").ColumnType(schema.ColumnTypeJSON).Description("Autonomous Database configuration details for storing manual backups in the Object Storage service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_details").ColumnType(schema.ColumnTypeString).Description("Information about the current lifecycle state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the Autonomous Database was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_data_guard_enabled").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the Autonomous Database has Data Guard enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operations_insights_status").ColumnType(schema.ColumnTypeString).Description("Status of Operations Insights for this Autonomous Database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("standby_db").ColumnType(schema.ColumnTypeJSON).Description("Autonomous Data Guard standby database details.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("refreshable_mode").ColumnType(schema.ColumnTypeString).Description("The refresh mode of the clone. AUTOMATIC indicates that the clone is automatically being refreshed with data from the source Autonomous Database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_deletion_of_free_autonomous_database").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the Always Free database will be automatically deleted because of inactivity. If the database is in the STOPPED state and without activity until this time, it will be deleted.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_of_last_failover").ColumnType(schema.ColumnTypeTimestamp).Description("The timestamp of the last failover operation.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connection_strings").ColumnType(schema.ColumnTypeJSON).Description("The connection string used to connect to the Autonomous Database. The username for the Service Console is ADMIN. Use the password you entered when creating the Autonomous Database for the password value.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connection_urls").ColumnType(schema.ColumnTypeJSON).Description("The URLs for accessing Oracle Application Express (APEX) and SQL Developer Web with a browser from a Compute instance within your VCN or that has a direct connection to your VCN. Note that these URLs are provided by the console only for databases on dedicated Exadata infrastructure.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("standby_whitelisted_ips").ColumnType(schema.ColumnTypeJSON).Description("The client IP access control list (ACL). This feature is available for autonomous databases on shared Exadata infrastructure and on Exadata Cloud@Customer. Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. For shared Exadata infrastructure, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_name").ColumnType(schema.ColumnTypeString).Description("The database name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_safe_status").ColumnType(schema.ColumnTypeString).Description("Status of the Data Safe registration for this Autonomous Database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_workload").ColumnType(schema.ColumnTypeString).Description("The Autonomous Database workload type.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role").ColumnType(schema.ColumnTypeString).Description("The role of the Autonomous Data Guard-enabled Autonomous Container Database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_of_next_refresh").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time of next refresh.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failed_data_recovery_in_seconds").ColumnType(schema.ColumnTypeInt).Description("Indicates the number of seconds of data loss for a Data Guard failover.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_free_tier").ColumnType(schema.ColumnTypeBool).Description("Indicates if this is an Always Free resource. The default value is false.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_endpoint_ip").ColumnType(schema.ColumnTypeString).Description("The private endpoint Ip address for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_console_url").ColumnType(schema.ColumnTypeString).Description("The URL of the Service Console for the Autonomous Database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the subnet the resource is associated with.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("autonomous_container_database_id").ColumnType(schema.ColumnTypeString).Description("The Autonomous Container Database OCID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_store_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the key store.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_endpoint").ColumnType(schema.ColumnTypeString).Description("The private endpoint for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The user-friendly name for the Autonomous Database. The name does not have to be unique.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("are_primary_whitelisted_ips_used").ColumnType(schema.ColumnTypeBool).Description("This field will be null if the Autonomous Database is not Data Guard enabled or Access Control is disabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_endpoint_label").ColumnType(schema.ColumnTypeString).Description("The private endpoint label for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_of_last_refresh_point").ColumnType(schema.ColumnTypeTimestamp).Description("The refresh point timestamp (UTC). The refresh point is the time to which the database was most recently refreshed. Data created after the refresh point is not included in the refresh.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_reclamation_of_free_autonomous_database").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the Always Free database will be stopped because of inactivity. If this time is reached without any database activity, the database will automatically be put into the STOPPED state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nsg_ids").ColumnType(schema.ColumnTypeJSON).Description("A list of the OCIDs of the network security groups (NSGs) that this resource belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the Autonomous Database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_storage_size_in_tbs").ColumnType(schema.ColumnTypeInt).Description("The quantity of data in the database, in terabytes.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_version").ColumnType(schema.ColumnTypeString).Description("A valid Oracle Database version for Autonomous Database.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_preview").ColumnType(schema.ColumnTypeBool).Description("Indicates if the Autonomous Database version is a preview version.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_store_wallet_name").ColumnType(schema.ColumnTypeString).Description("The wallet name for Oracle Key Vault.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_maintenance_end").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time when maintenance will end.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("used_data_storage_size_in_tbs").ColumnType(schema.ColumnTypeInt).Description("The amount of storage that has been used, in terabytes.").Build(),
	}
}

func (x *TableOciDatabaseAutonomousDatabaseGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableOciDatabaseAutonomousDbMetricStorageUtilizationDailyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciDatabaseAutonomousDbMetricCpuUtilizationDailyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciDatabaseAutonomousDbMetricStorageUtilizationGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciDatabaseAutonomousDbMetricCpuUtilizationHourlyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciDatabaseAutonomousDbMetricStorageUtilizationHourlyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciDatabaseAutonomousDbMetricCpuUtilizationGenerator{}),
	}
}
