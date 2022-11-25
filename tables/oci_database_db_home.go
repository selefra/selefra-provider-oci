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

type TableOciDatabaseDbHomeGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciDatabaseDbHomeGenerator{}

func (x *TableOciDatabaseDbHomeGenerator) GetTableName() string {
	return "oci_database_db_home"
}

func (x *TableOciDatabaseDbHomeGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciDatabaseDbHomeGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciDatabaseDbHomeGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciDatabaseDbHomeGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.DatabaseService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildDatabaseDBHomeFilters()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.DatabaseClient.ListDbHomes(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, dbHome := range response.Items {
					resultChannel <- dbHome

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

func buildDatabaseDBHomeFilters() database.ListDbHomesRequest {
	request := database.ListDbHomesRequest{}

	return request
}

func (x *TableOciDatabaseDbHomeGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciDatabaseDbHomeGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the database home was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the database home.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the database home.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_system_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the DB system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_software_image_id").ColumnType(schema.ColumnTypeString).Description("The database software image OCID.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_version").ColumnType(schema.ColumnTypeString).Description("The oracle database version.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_details").ColumnType(schema.ColumnTypeString).Description("Additional information about the current lifecycle state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The user-friendly name for the database home. It does not have to be unique.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("db_home_location").ColumnType(schema.ColumnTypeString).Description("The location of the oracle database home.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vm_cluster_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the VM cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_patch_history_entry_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the last patch history.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("one_off_patches").ColumnType(schema.ColumnTypeJSON).Description("List of one-off patches for database homes.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
	}
}

func (x *TableOciDatabaseDbHomeGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableOciDatabaseDbGenerator{}),
	}
}
