package tables

import (
	"context"
	"github.com/oracle/oci-go-sdk/v44/common"

	"github.com/oracle/oci-go-sdk/v44/identity"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciIdentityCustomerSecretKeyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciIdentityCustomerSecretKeyGenerator{}

func (x *TableOciIdentityCustomerSecretKeyGenerator) GetTableName() string {
	return "oci_identity_customer_secret_key"
}

func (x *TableOciIdentityCustomerSecretKeyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciIdentityCustomerSecretKeyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciIdentityCustomerSecretKeyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciIdentityCustomerSecretKeyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			user := task.ParentRawResult.(identity.User)

			session, err := oci_client.IdentityService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := identity.ListCustomerSecretKeysRequest{
				UserId: user.Id,
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			item, err := session.IdentityClient.ListCustomerSecretKeys(ctx, request)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, secretKey := range item.Items {
				resultChannel <- customerSecretKeyInfo{secretKey, *user.Name}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type customerSecretKeyInfo struct {
	identity.CustomerSecretKeySummary
	UserName string
}

func (x *TableOciIdentityCustomerSecretKeyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOciIdentityCustomerSecretKeyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the secret key.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The secret key's current state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inactive_status").ColumnType(schema.ColumnTypeInt).Description("The detailed status of INACTIVE lifecycleState.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time the CustomerSecretKey object was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_expires").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time when this password will expire.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("The displayName you assign to the secret key.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the user the password belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_name").ColumnType(schema.ColumnTypeString).Description("The name of the user the password belongs to.").Build(),
	}
}

func (x *TableOciIdentityCustomerSecretKeyGenerator) GetSubTables() []*schema.Table {
	return nil
}
