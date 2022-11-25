package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/filestorage"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableOciFileStorageFileSystemGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciFileStorageFileSystemGenerator{}

func (x *TableOciFileStorageFileSystemGenerator) GetTableName() string {
	return "oci_file_storage_file_system"
}

func (x *TableOciFileStorageFileSystemGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciFileStorageFileSystemGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciFileStorageFileSystemGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciFileStorageFileSystemGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.FileStorageService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildFileStorageFileSystemFilters()
			request.CompartmentId = &session.TenancyID
			request.AvailabilityDomain = pointer.ToStringPointer(taskClient.(*oci_client.OciClient).Zone)
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.FileStorageClient.ListFileSystems(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, fileSystems := range response.Items {
					resultChannel <- fileSystems

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

func buildFileStorageFileSystemFilters() filestorage.ListFileSystemsRequest {
	request := filestorage.ListFileSystemsRequest{}

	return request
}

func (x *TableOciFileStorageFileSystemGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildCompartementZonalList()
}

func (x *TableOciFileStorageFileSystemGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the file system was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_clone_parent").ColumnType(schema.ColumnTypeBool).Description("Specifies whether the file system has been cloned.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the file system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_details").ColumnType(schema.ColumnTypeString).Description("Additional information about the current 'lifecycleState'.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_details").ColumnType(schema.ColumnTypeJSON).Description("Source information for the file system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name. It does not have to be unique, and it is changeable.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the file system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_domain").ColumnType(schema.ColumnTypeString).Description("The availability domain the file system is in.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_hydrated").ColumnType(schema.ColumnTypeBool).Description("Specifies whether the data has finished copying from the source to the clone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the KMS key used to encrypt the encryption keys associated with this file system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metered_bytes").ColumnType(schema.ColumnTypeInt).Description("The number of bytes consumed by the file system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
	}
}

func (x *TableOciFileStorageFileSystemGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableOciFileStorageSnapshotGenerator{}),
	}
}
