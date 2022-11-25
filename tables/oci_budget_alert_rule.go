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

type TableOciBudgetAlertRuleGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciBudgetAlertRuleGenerator{}

func (x *TableOciBudgetAlertRuleGenerator) GetTableName() string {
	return "oci_budget_alert_rule"
}

func (x *TableOciBudgetAlertRuleGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciBudgetAlertRuleGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciBudgetAlertRuleGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciBudgetAlertRuleGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.BudgetService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			id := task.ParentRawResult.(budget.BudgetSummary).Id

			request := budget.ListAlertRulesRequest{
				BudgetId: id,
				Limit:    pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.BudgetClient.ListAlertRules(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, rule := range response.Items {
					resultChannel <- AlertRuleInfo{rule.Id, rule.BudgetId, rule.DisplayName, rule.Type, rule.Threshold, rule.ThresholdType, rule.LifecycleState, rule.Recipients, rule.TimeCreated, rule.TimeUpdated, rule.Message, rule.Description, rule.Version, rule.FreeformTags, rule.DefinedTags, session.TenancyID}

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

type AlertRuleInfo struct {
	Id *string `mandatory:"true" json:"id"`

	BudgetId *string `mandatory:"true" json:"budgetId"`

	DisplayName *string `mandatory:"true" json:"displayName"`

	Type budget.AlertTypeEnum `mandatory:"true" json:"type"`

	Threshold *float32 `mandatory:"true" json:"threshold"`

	ThresholdType budget.ThresholdTypeEnum `mandatory:"true" json:"thresholdType"`

	LifecycleState budget.LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	Recipients *string `mandatory:"true" json:"recipients"`

	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	Message *string `mandatory:"false" json:"message"`

	Description *string `mandatory:"false" json:"description"`

	Version *int `mandatory:"false" json:"version"`

	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	CompartmentId string
}

func (x *TableOciBudgetAlertRuleGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciBudgetAlertRuleGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The name of the alert rule.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the alert rule.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the alert rule.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the alert rule.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("Time that budget was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message").ColumnType(schema.ColumnTypeString).Description("Custom message sent when alert is triggered.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recipients").ColumnType(schema.ColumnTypeString).Description("Delimited list of email addresses to receive the alert when it triggers.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_updated").ColumnType(schema.ColumnTypeTimestamp).Description("Time that budget was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Description("The type of alert.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("threshold").ColumnType(schema.ColumnTypeFloat).Description("The threshold for triggering the alert. If thresholdType is PERCENTAGE, the maximum value is 10000.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("threshold_type").ColumnType(schema.ColumnTypeString).Description("The type of threshold.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeInt).Description("Version of the alert rule. Starts from 1 and increments by 1.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("budget_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the budget").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
	}
}

func (x *TableOciBudgetAlertRuleGenerator) GetSubTables() []*schema.Table {
	return nil
}
