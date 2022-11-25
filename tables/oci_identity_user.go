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

type TableOciIdentityUserGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciIdentityUserGenerator{}

func (x *TableOciIdentityUserGenerator) GetTableName() string {
	return "oci_identity_user"
}

func (x *TableOciIdentityUserGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciIdentityUserGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciIdentityUserGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciIdentityUserGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.IdentityService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildUserGroupFilters()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.IdentityClient.ListUsers(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, user := range response.Items {
					resultChannel <- user

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

func buildUserGroupFilters() identity.ListUsersRequest {
	request := identity.ListUsersRequest{}

	return request
}

func (x *TableOciIdentityUserGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOciIdentityUserGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("can_use_auth_tokens").ColumnType(schema.ColumnTypeBool).Description("Indicates if the user can use SWIFT passwords/auth tokens.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email_verified").ColumnType(schema.ColumnTypeBool).Description("Whether the email address has been validated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity_provider_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the `IdentityProvider` this user belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inactive_status").ColumnType(schema.ColumnTypeInt).Description("Applicable only if the user's `lifecycleState` is INACTIVE. A 16-bit value showing the reason why the user is inactive. 0: SUSPENDED; 1: DISABLED; 2: BLOCKED (the user has exceeded the maximum number of failed login attempts for the Console)").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time the user was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description assigned to the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_mfa_activated").ColumnType(schema.ColumnTypeBool).Description("The user's current state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("The email address you assign to the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_groups").ColumnType(schema.ColumnTypeJSON).Description("List of groups associated with the user.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_type").ColumnType(schema.ColumnTypeString).Description("Type of the user. Value can be IDCS or IAM. Oracle Identity Cloud Service(IDCS) users authenticate through single sign-on and can be granted access to all services included in your account. IAM users can access Oracle Cloud Infrastructure services, but not all Cloud Platform services.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("can_use_api_keys").ColumnType(schema.ColumnTypeBool).Description("Indicates if the user can use API keys.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("can_use_smtp_credentials").ColumnType(schema.ColumnTypeBool).Description("Indicates if the user can use SMTP passwords.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external_identifier").ColumnType(schema.ColumnTypeString).Description("Identifier of the user in the identity provider.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The user's login for the Console.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The user's current state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("can_use_console_password").ColumnType(schema.ColumnTypeBool).Description("Indicates if the user can log in to the console.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("can_use_customer_secret_keys").ColumnType(schema.ColumnTypeBool).Description("Indicates if the user can use SigV4 symmetric keys.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("can_use_o_auth2_client_credentials").ColumnType(schema.ColumnTypeBool).Description("Indicates if the user can use OAuth2 credentials and tokens.").Build(),
	}
}

func (x *TableOciIdentityUserGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableOciIdentityAuthTokenGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciIdentityCustomerSecretKeyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciIdentityApiKeyGenerator{}),
	}
}
