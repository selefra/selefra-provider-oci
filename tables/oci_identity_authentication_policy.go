package tables

import (
	"context"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/identity"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciIdentityAuthenticationPolicyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciIdentityAuthenticationPolicyGenerator{}

func (x *TableOciIdentityAuthenticationPolicyGenerator) GetTableName() string {
	return "oci_identity_authentication_policy"
}

func (x *TableOciIdentityAuthenticationPolicyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciIdentityAuthenticationPolicyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciIdentityAuthenticationPolicyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciIdentityAuthenticationPolicyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.IdentityService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := identity.GetAuthenticationPolicyRequest{
				CompartmentId: &session.TenancyID,
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			response, err := session.IdentityClient.GetAuthenticationPolicy(ctx, request)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- response.AuthenticationPolicy
			return schema.NewDiagnosticsErrorPullTable(task.Table, err)

		},
	}
}

func (x *TableOciIdentityAuthenticationPolicyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOciIdentityAuthenticationPolicyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_lowercase_characters_required").ColumnType(schema.ColumnTypeBool).Description("At least one lower case character required.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_numeric_characters_required").ColumnType(schema.ColumnTypeBool).Description("At least one numeric character required.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_special_characters_required").ColumnType(schema.ColumnTypeBool).Description("At least one special character required.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_uppercase_characters_required").ColumnType(schema.ColumnTypeBool).Description("At least one uppercase character required.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_username_containment_allowed").ColumnType(schema.ColumnTypeBool).Description("User name is allowed to be part of the password.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("minimum_password_length").ColumnType(schema.ColumnTypeInt).Description("Minimum password length required.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_source_ids").ColumnType(schema.ColumnTypeString).Description("List of IP ranges from which users can sign in to the Console.").Build(),
	}
}

func (x *TableOciIdentityAuthenticationPolicyGenerator) GetSubTables() []*schema.Table {
	return nil
}
