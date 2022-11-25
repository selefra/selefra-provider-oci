package tables

import (
	"context"

	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableOciCloudGuardConfigurationGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCloudGuardConfigurationGenerator{}

func (x *TableOciCloudGuardConfigurationGenerator) GetTableName() string {
	return "oci_cloud_guard_configuration"
}

func (x *TableOciCloudGuardConfigurationGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCloudGuardConfigurationGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCloudGuardConfigurationGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCloudGuardConfigurationGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

func (x *TableOciCloudGuardConfigurationGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOciCloudGuardConfigurationGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("reporting_region").ColumnType(schema.ColumnTypeString).Description("The reporting region value.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("Status of Cloud Guard Tenant.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_manage_resources").ColumnType(schema.ColumnTypeBool).Description("Identifies if Oracle managed resources were created by customers.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
	}
}

func (x *TableOciCloudGuardConfigurationGenerator) GetSubTables() []*schema.Table {
	return nil
}
