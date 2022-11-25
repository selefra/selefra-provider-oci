package tables

import (
	"context"
	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/identity"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciIdentityAvailabilityDomainGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciIdentityAvailabilityDomainGenerator{}

func (x *TableOciIdentityAvailabilityDomainGenerator) GetTableName() string {
	return "oci_identity_availability_domain"
}

func (x *TableOciIdentityAvailabilityDomainGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciIdentityAvailabilityDomainGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciIdentityAvailabilityDomainGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

type availabilityDomainInfo struct {
	identity.AvailabilityDomain
	Region string
}

func (x *TableOciIdentityAvailabilityDomainGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			region := taskClient.(*oci_client.OciClient).Region

			status := task.ParentRawResult.(ociRegion).Status

			if status != "READY" {
				return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
			}

			session, err := oci_client.IdentityServiceRegional(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := identity.ListAvailabilityDomainsRequest{
				CompartmentId: &session.TenancyID,
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			response, err := session.IdentityClient.ListAvailabilityDomains(ctx, request)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, availabilityDomain := range response.Items {
				resultChannel <- availabilityDomainInfo{availabilityDomain, region}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableOciIdentityAvailabilityDomainGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOciIdentityAvailabilityDomainGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the Availability Domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Availability Domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
	}
}

func (x *TableOciIdentityAvailabilityDomainGenerator) GetSubTables() []*schema.Table {
	return nil
}
