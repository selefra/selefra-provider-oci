package tables

import (
	"context"
	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/identity"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableOciIdentityApiKeyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciIdentityApiKeyGenerator{}

func (x *TableOciIdentityApiKeyGenerator) GetTableName() string {
	return "oci_identity_api_key"
}

func (x *TableOciIdentityApiKeyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciIdentityApiKeyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciIdentityApiKeyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciIdentityApiKeyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			user := task.ParentRawResult.(identity.User)

			session, err := oci_client.IdentityService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := identity.ListApiKeysRequest{
				UserId: user.Id,
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			item, err := session.IdentityClient.ListApiKeys(ctx, request)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, apiKey := range item.Items {
				resultChannel <- apiKeyInfo{apiKey, *user.Name}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type apiKeyInfo struct {
	identity.ApiKey
	UserName string
}

func (x *TableOciIdentityApiKeyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOciIdentityApiKeyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_id").ColumnType(schema.ColumnTypeString).Description("An Oracle-assigned identifier for the key.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the user the key belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_name").ColumnType(schema.ColumnTypeString).Description("The name of the user the key belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The API key's current state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fingerprint").ColumnType(schema.ColumnTypeString).Description("The key's fingerprint.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_value").ColumnType(schema.ColumnTypeString).Description("The key's value.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time the `ApiKey` object was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inactive_status").ColumnType(schema.ColumnTypeInt).Description("The detailed status of INACTIVE lifecycleState.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
	}
}

func (x *TableOciIdentityApiKeyGenerator) GetSubTables() []*schema.Table {
	return nil
}

