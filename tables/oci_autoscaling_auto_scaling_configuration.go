package tables

import (
	"context"
	"github.com/oracle/oci-go-sdk/v44/autoscaling"
	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-utils/pkg/pointer"
)

type TableOciAutoscalingAutoScalingConfigurationGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciAutoscalingAutoScalingConfigurationGenerator{}

func (x *TableOciAutoscalingAutoScalingConfigurationGenerator) GetTableName() string {
	return "oci_autoscaling_auto_scaling_configuration"
}

func (x *TableOciAutoscalingAutoScalingConfigurationGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciAutoscalingAutoScalingConfigurationGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciAutoscalingAutoScalingConfigurationGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciAutoscalingAutoScalingConfigurationGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.AutoScalingService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := autoscaling.ListAutoScalingConfigurationsRequest{
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.AutoScalingClient.ListAutoScalingConfigurations(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, configuration := range response.Items {
					resultChannel <- configuration

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

func (x *TableOciAutoscalingAutoScalingConfigurationGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciAutoscalingAutoScalingConfigurationGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("resource").ColumnType(schema.ColumnTypeJSON).Description("The resource details of this AutoScalingConfiguration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the autoscaling configuration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_enabled").ColumnType(schema.ColumnTypeBool).Description("Indicates whether the autoscaling configuration is enabled.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the AutoScalingConfiguration was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cool_down_in_seconds").ColumnType(schema.ColumnTypeInt).Description("The minimum period of time to wait between scaling actions.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_resource_count").ColumnType(schema.ColumnTypeInt).Description("The maximum number of resources to scale out to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("min_resource_count").ColumnType(schema.ColumnTypeInt).Description("The minimum number of resources to scale in to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policies").ColumnType(schema.ColumnTypeJSON).Description("Autoscaling policy definitions for the autoscaling configuration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
	}
}

func (x *TableOciAutoscalingAutoScalingConfigurationGenerator) GetSubTables() []*schema.Table {
	return nil
}
