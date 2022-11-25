package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/database"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableOciDatabaseDbSystemGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciDatabaseDbSystemGenerator{}

func (x *TableOciDatabaseDbSystemGenerator) GetTableName() string {
	return "oci_database_db_system"
}

func (x *TableOciDatabaseDbSystemGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciDatabaseDbSystemGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciDatabaseDbSystemGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciDatabaseDbSystemGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.DatabaseService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildDatabaseDBSystemFilters()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.DatabaseClient.ListDbSystems(ctx, request)
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

func buildDatabaseDBSystemFilters() database.ListDbSystemsRequest {
	request := database.ListDbSystemsRequest{}

	return request
}

func (x *TableOciDatabaseDbSystemGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciDatabaseDbSystemGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("scan_ip_ids").ColumnType(schema.ColumnTypeJSON).Description("The OCID of the single client access name (SCAN) IP addresses associated with the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_maintenance_run_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the next maintenance run.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeString).Description("The oracle database version of the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the zone the DB system is associated with.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sparse_diskgroup").ColumnType(schema.ColumnTypeBool).Description("True, If sparse diskgroup is configured for exadata DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance_window").ColumnType(schema.ColumnTypeJSON).Description("The maintenance window of the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ssh_public_keys").ColumnType(schema.ColumnTypeJSON).Description("The public key portion of one or more key pairs used for SSH access to the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain").ColumnType(schema.ColumnTypeString).Description("The domain name for the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("license_model").ColumnType(schema.ColumnTypeString).Description("The oracle license model that applies to all the databases on the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scan_dns_name").ColumnType(schema.ColumnTypeString).Description("The FQDN of the DNS record for the SCAN IP addresses that are associated with the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_db_system_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the DB system from where the DB system is created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reco_storage_size_in_gb").ColumnType(schema.ColumnTypeInt).Description("The RECO/REDO storage size, in gigabytes, that is currently allocated to the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iorm_config_cache").ColumnType(schema.ColumnTypeJSON).Description("The IORM configuration of the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nsg_ids").ColumnType(schema.ColumnTypeJSON).Description("A list of the OCIDs of the network security groups (NSGs) that this resource belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_storage_percentage").ColumnType(schema.ColumnTypeInt).Description("The percentage assigned to data storage.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("host_name").ColumnType(schema.ColumnTypeString).Description("The hostname for the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("listener_port").ColumnType(schema.ColumnTypeInt).Description("The port number configured for the listener on the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the subnet the DB system is associated with.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the DB system was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_storage_size_in_gbs").ColumnType(schema.ColumnTypeInt).Description("The data storage size, in gigabytes, that is currently available to the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_system_options_storage_management").ColumnType(schema.ColumnTypeString).Description("The storage option used in DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_count").ColumnType(schema.ColumnTypeInt).Description("The number of nodes in the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_details").ColumnType(schema.ColumnTypeString).Description("Additional information about the current lifecycle state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fault_domains").ColumnType(schema.ColumnTypeJSON).Description("List of the fault domains in which this DB system is provisioned.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vip_ids").ColumnType(schema.ColumnTypeJSON).Description("A list of the OCIDs of the virtual IP (VIP) addresses associated with the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_domain").ColumnType(schema.ColumnTypeString).Description("The name of the availability domain that the DB system is located in.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup_subnet_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the backup network subnet the DB system is associated with.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_redundancy").ColumnType(schema.ColumnTypeString).Description("The type of redundancy configured for the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_patch_history_entry_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the last patch history.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The user-friendly name for the DB system. The name does not have to be unique.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cpu_core_count").ColumnType(schema.ColumnTypeInt).Description("The number of CPU cores enabled on the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("point_in_time_data_disk_clone_timestamp").ColumnType(schema.ColumnTypeTimestamp).Description("The point in time for a cloned database system when the data disks were cloned from the source database system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup_network_nsg_ids").ColumnType(schema.ColumnTypeJSON).Description("A list of the OCIDs of the network security groups (NSGs) that the backup network of this DB system belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_zone").ColumnType(schema.ColumnTypeString).Description("The time zone of the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_name").ColumnType(schema.ColumnTypeString).Description("The cluster name for exadata and 2-node RAC virtual machine DB systems.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_maintenance_run_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the last maintenance run.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scan_dns_record_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the DNS record for the SCAN IP addresses that are associated with the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_edition").ColumnType(schema.ColumnTypeString).Description("The oracle database edition that applies to all the databases on the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
	}
}

func (x *TableOciDatabaseDbSystemGenerator) GetSubTables() []*schema.Table {
	return nil
}
