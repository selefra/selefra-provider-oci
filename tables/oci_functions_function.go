package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/functions"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciFunctionsFunctionGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciFunctionsFunctionGenerator{}

func (x *TableOciFunctionsFunctionGenerator) GetTableName() string {
	return "oci_functions_function"
}

func (x *TableOciFunctionsFunctionGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciFunctionsFunctionGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciFunctionsFunctionGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciFunctionsFunctionGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			var applicationId string

			if applicationId == "" {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}

			session, err := oci_client.FunctionsManagementService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildFunctionsFilters()
			request.ApplicationId = &applicationId
			request.Limit = pointer.ToIntPointer(50)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.FunctionsManagementClient.ListFunctions(ctx, request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, item := range response.Items {
					resultChannel <- item

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

func buildFunctionsFilters() functions.ListFunctionsRequest {
	request := functions.ListFunctionsRequest{}

	return request
}

func (x *TableOciFunctionsFunctionGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciFunctionsFunctionGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The display name of the function.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the function.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("memory_in_mbs").ColumnType(schema.ColumnTypeInt).Description("Maximum usable memory for the function (MiB).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timeout_in_seconds").ColumnType(schema.ColumnTypeInt).Description("Timeout for executions of the function. Value in seconds.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("config").ColumnType(schema.ColumnTypeJSON).Description("The function configuration. Overrides application configuration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the function.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_digest").ColumnType(schema.ColumnTypeString).Description("The image digest for the version of the image that will be pulled when invoking this function. If no value is specified, the digest currently associated with the image in the OCI Registry will be used.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_updated").ColumnType(schema.ColumnTypeTimestamp).Description("The time the function was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("application_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the application the function belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image").ColumnType(schema.ColumnTypeString).Description("The qualified name of the Docker image to use in the function, including the image tag. The image should be in the OCI Registry that is in the same region as the function itself.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("invoke_endpoint").ColumnType(schema.ColumnTypeString).Description("The base https invoke URL to set on a client in order to invoke a function. This URL will never change over the lifetime of the function and can be cached.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The time the function was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("trace_config").ColumnType(schema.ColumnTypeJSON).Description("The trace configuration of the function.").Build(),
	}
}

func (x *TableOciFunctionsFunctionGenerator) GetSubTables() []*schema.Table {
	return nil
}
