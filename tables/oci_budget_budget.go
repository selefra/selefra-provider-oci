package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/budget"
	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciBudgetBudgetGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciBudgetBudgetGenerator{}

func (x *TableOciBudgetBudgetGenerator) GetTableName() string {
	return "oci_budget_budget"
}

func (x *TableOciBudgetBudgetGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciBudgetBudgetGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciBudgetBudgetGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciBudgetBudgetGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.BudgetService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := budget.ListBudgetsRequest{
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToIntPointer(1000),
				TargetType:    "ALL",
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.BudgetClient.ListBudgets(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, budget := range response.Items {
					resultChannel <- budget

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

func (x *TableOciBudgetBudgetGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciBudgetBudgetGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the budget.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("amount").ColumnType(schema.ColumnTypeFloat).Description("The amount of the budget expressed in the currency of the customer's rate card.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reset_period").ColumnType(schema.ColumnTypeString).Description("The reset period for the budget.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("targets").ColumnType(schema.ColumnTypeJSON).Description("The list of targets on which the budget is applied.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the budget.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the budget.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("Time that budget was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alert_rule_count").ColumnType(schema.ColumnTypeInt).Description("Total number of alert rules in the budget.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("forecasted_spend").ColumnType(schema.ColumnTypeFloat).Description("The forecasted spend in currency by the end of the current budget cycle.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_updated").ColumnType(schema.ColumnTypeTimestamp).Description("Time that budget was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The display name of the budget.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_compartment_id").ColumnType(schema.ColumnTypeString).Description("This is DEPRECATED. For backwards compatability, the property will be populated when targetType is COMPARTMENT AND targets contains EXACT ONE target compartment ocid. For all other scenarios, this property will be left empty.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("budget_processing_period_start_offset").ColumnType(schema.ColumnTypeInt).Description("The number of days offset from the first day of the month, at which the budget processing period starts.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_type").ColumnType(schema.ColumnTypeString).Description("The type of target on which the budget is applied.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_spend_computed").ColumnType(schema.ColumnTypeTimestamp).Description("The time that the budget spend was last computed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeInt).Description("Version of the budget. Starts from 1 and increments by 1.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("actual_spend").ColumnType(schema.ColumnTypeFloat).Description("The actual spend in currency for the current budget cycle.").Build(),
	}
}

func (x *TableOciBudgetBudgetGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableOciBudgetAlertRuleGenerator{}),
	}
}
