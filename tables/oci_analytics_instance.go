package tables

import (
	"context"
	"github.com/oracle/oci-go-sdk/v44/analytics"
	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-utils/pkg/pointer"
)

type TableOciAnalyticsInstanceGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciAnalyticsInstanceGenerator{}

func (x *TableOciAnalyticsInstanceGenerator) GetTableName() string {
	return "oci_analytics_instance"
}

func (x *TableOciAnalyticsInstanceGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciAnalyticsInstanceGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciAnalyticsInstanceGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciAnalyticsInstanceGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.AnalyticsService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildAnalyticsInstanceFilters()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.AnalyticsClient.ListAnalyticsInstances(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, instance := range response.Items {
					resultChannel <- instance

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

func buildAnalyticsInstanceFilters() analytics.ListAnalyticsInstancesRequest {
	request := analytics.ListAnalyticsInstancesRequest{}

	return request
}

func (x *TableOciAnalyticsInstanceGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciAnalyticsInstanceGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The analytics instance's optional description.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email_notification").ColumnType(schema.ColumnTypeString).Description("The email address receiving notifications.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the instance was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capacity_type").ColumnType(schema.ColumnTypeString).Description("The analytics instance's capacity model to use.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capacity_value").ColumnType(schema.ColumnTypeInt).Description("The analytics instance's capacity value selected.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_access_channels").ColumnType(schema.ColumnTypeJSON).Description("The private access channels of the analytics instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_updated").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the instance was last updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The analytics instance's current state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("feature_set").ColumnType(schema.ColumnTypeString).Description("The analytics instance's feature set.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("license_type").ColumnType(schema.ColumnTypeString).Description("The license used for the service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_url").ColumnType(schema.ColumnTypeString).Description("The URL of the Analytics service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name. Does not have to be unique, and it's changeable.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_endpoint_details").ColumnType(schema.ColumnTypeJSON).Description("The base representation of a network endpoint.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vanity_url_details").ColumnType(schema.ColumnTypeJSON).Description("The vanity url configuration details of the analytic instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The analytics instance's Oracle ID (OCID).").Build(),
	}
}

func (x *TableOciAnalyticsInstanceGenerator) GetSubTables() []*schema.Table {
	return nil
}
