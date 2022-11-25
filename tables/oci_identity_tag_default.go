package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/identity"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciIdentityTagDefaultGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciIdentityTagDefaultGenerator{}

func (x *TableOciIdentityTagDefaultGenerator) GetTableName() string {
	return "oci_identity_tag_default"
}

func (x *TableOciIdentityTagDefaultGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciIdentityTagDefaultGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciIdentityTagDefaultGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciIdentityTagDefaultGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.IdentityService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := identity.ListTagDefaultsRequest{
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.IdentityClient.ListTagDefaults(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}
				for _, tagDefault := range response.Items {
					resultChannel <- tagDefault

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

func (x *TableOciIdentityTagDefaultGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildCompartmentList()
}

func (x *TableOciIdentityTagDefaultGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tag_definition_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the tag definition.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_required").ColumnType(schema.ColumnTypeBool).Description("If you specify that a value is required, a value is set during resource creation (either by the user creating the resource or another tag default). If no value is set, resource creation is blocked.If the `isRequired` flag is set to true, the value is set during resource creation.If the `isRequired` flag is set to false, the value you enter is set during resource creation.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("value").ColumnType(schema.ColumnTypeString).Description("The default value for the tag definition.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The tag default's current state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time the tagDefault was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the tag default.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tag_definition_name").ColumnType(schema.ColumnTypeString).Description("The name used in the tag definition.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tag_namespace_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the tag namespace that contains the tag definition.").Build(),
	}
}

func (x *TableOciIdentityTagDefaultGenerator) GetSubTables() []*schema.Table {
	return nil
}
