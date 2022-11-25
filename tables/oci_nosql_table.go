package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/nosql"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
)

type TableOciNosqlTableGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciNosqlTableGenerator{}

func (x *TableOciNosqlTableGenerator) GetTableName() string {
	return "oci_nosql_table"
}

func (x *TableOciNosqlTableGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciNosqlTableGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciNosqlTableGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciNosqlTableGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.NoSQLDatabaseService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := nosql.ListTablesRequest{
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.NoSQLClient.ListTables(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, table := range response.Items {
					resultChannel <- table

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

func (x *TableOciNosqlTableGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciNosqlTableGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The state of a table.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_auto_reclaimable").ColumnType(schema.ColumnTypeBool).Description("True if this table can be reclaimed after an idle period.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Immutable human-friendly table name.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier that is immutable.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_of_expiration").ColumnType(schema.ColumnTypeTimestamp).Description("If lifecycleState is INACTIVE, indicates when this table will be automatically removed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("system_tags").ColumnType(schema.ColumnTypeJSON).Description("System tags for resource. System tags can be viewed by users, but can only be created by the system.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The time the the table was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ddl_statement").ColumnType(schema.ColumnTypeString).Description("A DDL statement representing the schema.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_details").ColumnType(schema.ColumnTypeString).Description("A message describing the current state in more detail.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_updated").ColumnType(schema.ColumnTypeTimestamp).Description("The time the the table was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schema").ColumnType(schema.ColumnTypeJSON).Description("The schema of the table.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("table_limits").ColumnType(schema.ColumnTypeJSON).Description("Various limit for the table.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
	}
}

func (x *TableOciNosqlTableGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableOciNosqlTableMetricReadThrottleCountHourlyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciNosqlTableMetricReadThrottleCountDailyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciNosqlTableMetricStorageUtilizationHourlyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciNosqlTableMetricReadThrottleCountGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciNosqlTableMetricStorageUtilizationDailyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciNosqlTableMetricStorageUtilizationGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciNosqlTableMetricWriteThrottleCountGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciNosqlTableMetricWriteThrottleCountHourlyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciNosqlTableMetricWriteThrottleCountDailyGenerator{}),
	}
}
