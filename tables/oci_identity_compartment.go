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

type TableOciIdentityCompartmentGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciIdentityCompartmentGenerator{}

func (x *TableOciIdentityCompartmentGenerator) GetTableName() string {
	return "oci_identity_compartment"
}

func (x *TableOciIdentityCompartmentGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciIdentityCompartmentGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciIdentityCompartmentGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciIdentityCompartmentGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.IdentityService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			rootRequest := identity.GetCompartmentRequest{
				CompartmentId: &session.TenancyID,
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			responseRoot, err := session.IdentityClient.GetCompartment(ctx, rootRequest)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			if responseRoot.CompartmentId != nil {
				resultChannel <- responseRoot.Compartment
			}

			request := identity.ListCompartmentsRequest{
				CompartmentId:          &session.TenancyID,
				CompartmentIdInSubtree: pointer.ToBoolPointer(true),
				Limit:                  pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.IdentityClient.ListCompartments(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, compartment := range response.Items {
					resultChannel <- compartment

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

func (x *TableOciIdentityCompartmentGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOciIdentityCompartmentGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name assigned to the compartment during creation").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("Date and time the user was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The compartment's current state.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("The description you assign to the compartment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inactive_status").ColumnType(schema.ColumnTypeInt).Description("The detailed status of INACTIVE lifecycleState").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_accessible").ColumnType(schema.ColumnTypeBool).Description("Indicates whether or not the compartment is accessible for the user making the request.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
	}
}

func (x *TableOciIdentityCompartmentGenerator) GetSubTables() []*schema.Table {
	return nil
}
