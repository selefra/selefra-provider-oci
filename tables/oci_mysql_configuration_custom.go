package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/mysql"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciMysqlConfigurationCustomGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciMysqlConfigurationCustomGenerator{}

func (x *TableOciMysqlConfigurationCustomGenerator) GetTableName() string {
	return "oci_mysql_configuration_custom"
}

func (x *TableOciMysqlConfigurationCustomGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciMysqlConfigurationCustomGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciMysqlConfigurationCustomGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciMysqlConfigurationCustomGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.MySQLConfigurationService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildMySQLConfigurationFilters()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.Type = []mysql.ListConfigurationsTypeEnum{"CUSTOM"}
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.MySQLConfigurationClient.ListConfigurations(ctx, request)
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

func (x *TableOciMysqlConfigurationCustomGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciMysqlConfigurationCustomGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The display name of the configuration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The configuration type, DEFAULT or CUSTOM.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the configuration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the configuration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parent_configuration_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the configuration from which this configuration is derived.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("User-provided data about the configuration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("variables").ColumnType(schema.ColumnTypeString).Description("User controllable service variables of the configuration.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the configuration was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shape_name").ColumnType(schema.ColumnTypeString).Description("The name of the associated shape.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_updated").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the configuration was last updated.").Build(),
	}
}

func (x *TableOciMysqlConfigurationCustomGenerator) GetSubTables() []*schema.Table {
	return nil
}
