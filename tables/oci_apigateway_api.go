package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/apigateway"
	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableOciApigatewayApiGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciApigatewayApiGenerator{}

func (x *TableOciApigatewayApiGenerator) GetTableName() string {
	return "oci_apigateway_api"
}

func (x *TableOciApigatewayApiGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciApigatewayApiGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciApigatewayApiGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciApigatewayApiGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.ApiGatewayService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := apigateway.ListApisRequest{
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.ApiGatewayClient.ListApis(ctx, request)
				if err != nil {

					if ociErr, ok := err.(common.ServiceError); ok {
						if ociErr.GetCode() == "InvalidParameter" {
							return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
						}
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, api := range response.Items {
					resultChannel <- api

				}

				if response.OpcNextPage != nil {
					request.Page = response.OpcNextPage
				} else {
					pagesLeft = false
				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableOciApigatewayApiGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciApigatewayApiGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The time this resource was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_details").ColumnType(schema.ColumnTypeString).Description("A message describing the current lifecycleState.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_updated").ColumnType(schema.ColumnTypeTimestamp).Description("The time this resource was last updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("validation_results").ColumnType(schema.ColumnTypeJSON).Description("Status of each feature available from the API.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name. Does not have to be unique, and it's changeable.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the API.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("specification_type").ColumnType(schema.ColumnTypeString).Description("Type of API Specification file.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
	}
}

func (x *TableOciApigatewayApiGenerator) GetSubTables() []*schema.Table {
	return nil
}
