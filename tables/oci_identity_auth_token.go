package tables

import (
	"context"
	"github.com/oracle/oci-go-sdk/v44/common"

	"github.com/oracle/oci-go-sdk/v44/identity"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciIdentityAuthTokenGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciIdentityAuthTokenGenerator{}

func (x *TableOciIdentityAuthTokenGenerator) GetTableName() string {
	return "oci_identity_auth_token"
}

func (x *TableOciIdentityAuthTokenGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciIdentityAuthTokenGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciIdentityAuthTokenGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciIdentityAuthTokenGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			user := task.ParentRawResult.(identity.User)

			session, err := oci_client.IdentityService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := identity.ListAuthTokensRequest{
				UserId: user.Id,
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			item, err := session.IdentityClient.ListAuthTokens(ctx, request)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, authToken := range item.Items {
				resultChannel <- authTokenInfo{authToken, *user.Name}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type authTokenInfo struct {
	identity.AuthToken
	UserName string
}

func (x *TableOciIdentityAuthTokenGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOciIdentityAuthTokenGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the auth token.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_name").ColumnType(schema.ColumnTypeString).Description("The name of the user the auth token belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_expires").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time when this auth token will expire.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inactive_status").ColumnType(schema.ColumnTypeInt).Description("The detailed status of INACTIVE lifecycleState.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the user the auth token belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("token").ColumnType(schema.ColumnTypeString).Description("The auth token. The value is available only in the response for `CreateAuthToken`, and not for `ListAuthTokens` or `UpdateAuthToken`.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The token's current state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time the `AuthToken` object was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description you assign to the auth token.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
	}
}

func (x *TableOciIdentityAuthTokenGenerator) GetSubTables() []*schema.Table {
	return nil
}
