package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/vault"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableOciVaultSecretGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciVaultSecretGenerator{}

func (x *TableOciVaultSecretGenerator) GetTableName() string {
	return "oci_vault_secret"
}

func (x *TableOciVaultSecretGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciVaultSecretGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciVaultSecretGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciVaultSecretGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.VaultService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildVaultSecretFilters()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.VaultClient.ListSecrets(ctx, request)
				if err != nil {

					if ociErr, ok := err.(common.ServiceError); ok {
						if ociErr.GetCode() == "InvalidParameter" || ociErr.GetCode() == "BadErrorResponse" {
							return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
						}
					}
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, vault := range response.Items {
					resultChannel <- vault

				}
				if response.OpcNextPage != nil {
					request.Page = response.OpcNextPage
				} else {
					pagesLeft = false
				}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func buildVaultSecretFilters() vault.ListSecretsRequest {
	request := vault.ListSecretsRequest{}

	return request
}

func (x *TableOciVaultSecretGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciVaultSecretGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the master encryption key that is used to encrypt the secret.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_version_number").ColumnType(schema.ColumnTypeInt).Description("The version number of the secret that's currently in use.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secret_rules").ColumnType(schema.ColumnTypeJSON).Description("A list of rules that control how the secret is used and managed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_of_current_version_expiry").ColumnType(schema.ColumnTypeString).Description("An optional property indicating when the current secret version will expire.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).Description("Additional metadata that you can use to provide context about how to use the secret or during rotation or other administrative tasks.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the secret.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("A brief description of the secret.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vault_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Vault in which the secret exists.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_details").ColumnType(schema.ColumnTypeString).Description("Additional information about the secret's current lifecycle state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeString).Description("A property indicating when the secret was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_of_deletion").ColumnType(schema.ColumnTypeString).Description("An optional property indicating when to delete the secret.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the secret.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current lifecycle state of the secret.").Build(),
	}
}

func (x *TableOciVaultSecretGenerator) GetSubTables() []*schema.Table {
	return nil
}
