package tables

import (
	"context"
	"github.com/oracle/oci-go-sdk/v44/common"

	"github.com/oracle/oci-go-sdk/v44/identity"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciIdentityTenancyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciIdentityTenancyGenerator{}

func (x *TableOciIdentityTenancyGenerator) GetTableName() string {
	return "oci_identity_tenancy"
}

func (x *TableOciIdentityTenancyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciIdentityTenancyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciIdentityTenancyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciIdentityTenancyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.IdentityService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := identity.GetTenancyRequest{
				TenancyId: &session.TenancyID,
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			response, err := session.IdentityClient.GetTenancy(ctx, request)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- response.Tenancy

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableOciIdentityTenancyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOciIdentityTenancyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the tenancy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retention_period_days").ColumnType(schema.ColumnTypeInt).Description("The retention period setting, specified in days.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("home_region_key").ColumnType(schema.ColumnTypeString).Description("The region key for the tenancy's home region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("upi_idcs_compatibility_layer_endpoint").ColumnType(schema.ColumnTypeString).Description("Url which refers to the UPI IDCS compatibility layer endpoint configured for this Tenant's home region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the tenancy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description of the tenancy.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
	}
}

func (x *TableOciIdentityTenancyGenerator) GetSubTables() []*schema.Table {
	return nil
}
