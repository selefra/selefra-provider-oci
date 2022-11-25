package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/dns"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableOciDnsRrsetGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciDnsRrsetGenerator{}

func (x *TableOciDnsRrsetGenerator) GetTableName() string {
	return "oci_dns_rrset"
}

func (x *TableOciDnsRrsetGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciDnsRrsetGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciDnsRrsetGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciDnsRrsetGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.DnsService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			zone := task.ParentRawResult.(dns.ZoneSummary)

			request := dns.GetZoneRecordsRequest{
				ZoneNameOrId:  zone.Id,
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToInt64Pointer(50),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.DnsClient.GetZoneRecords(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, record := range response.Items {
					resultChannel <- dnsRecordInfo{record, session.TenancyID}

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

type dnsRecordInfo struct {
	dns.Record
	CompartmentId string
}

func (x *TableOciDnsRrsetGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildCompartmentList()
}

func (x *TableOciDnsRrsetGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("domain").ColumnType(schema.ColumnTypeString).Description("The fully qualified domain name where the record can be located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("record_hash").ColumnType(schema.ColumnTypeString).Description("A unique identifier for the record within its zone.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_protected").ColumnType(schema.ColumnTypeBool).Description("A Boolean flag indicating whether or not parts of the record are unable to be explicitly managed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rrset_version").ColumnType(schema.ColumnTypeString).Description("The latest version of the record's zone in which its RRSet differs from the preceding version.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rtype").ColumnType(schema.ColumnTypeString).Description("The type of DNS record, such as A or CNAME.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rdata").ColumnType(schema.ColumnTypeString).Description("The record's data.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ttl").ColumnType(schema.ColumnTypeString).Description("The Time To Live for the record, in seconds.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
	}
}

func (x *TableOciDnsRrsetGenerator) GetSubTables() []*schema.Table {
	return nil
}
