package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/filestorage"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciFileStorageMountTargetGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciFileStorageMountTargetGenerator{}

func (x *TableOciFileStorageMountTargetGenerator) GetTableName() string {
	return "oci_file_storage_mount_target"
}

func (x *TableOciFileStorageMountTargetGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciFileStorageMountTargetGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciFileStorageMountTargetGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciFileStorageMountTargetGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.FileStorageService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildFileStorageMountTargetFilters()
			request.CompartmentId = &session.TenancyID
			request.AvailabilityDomain = pointer.ToStringPointer(taskClient.(*oci_client.OciClient).Zone)
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.FileStorageClient.ListMountTargets(ctx, request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, mountTarget := range response.Items {
					resultChannel <- mountTarget

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

func buildFileStorageMountTargetFilters() filestorage.ListMountTargetsRequest {
	request := filestorage.ListMountTargetsRequest{}

	return request
}

func (x *TableOciFileStorageMountTargetGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildCompartementZonalList()
}

func (x *TableOciFileStorageMountTargetGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name of the Mount Target.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_domain").ColumnType(schema.ColumnTypeString).Description("The availability domain the mount target is in.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("export_set_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the associated export set.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nsg_ids").ColumnType(schema.ColumnTypeJSON).Description("A list of Network Security Group OCIDs associated with this mount target.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_ip_ids").ColumnType(schema.ColumnTypeJSON).Description("The OCIDs of the private IP addresses associated with this mount target.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Mount Target.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the Mount Target.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_id").ColumnType(schema.ColumnTypeString).Description("The OCIDs of the subnet the mount target is in.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the Mount Target was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_details").ColumnType(schema.ColumnTypeString).Description("Additional information about the current 'lifecycleState'.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
	}
}

func (x *TableOciFileStorageMountTargetGenerator) GetSubTables() []*schema.Table {
	return nil
}
