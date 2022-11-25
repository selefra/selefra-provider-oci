package tables

import (
	"context"

	"github.com/oracle/oci-go-sdk/v44/identity"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciRegionGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciRegionGenerator{}

func (x *TableOciRegionGenerator) GetTableName() string {
	return "oci_region"
}

func (x *TableOciRegionGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciRegionGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciRegionGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

type ociRegion struct {
	identity.Region
	identity.RegionSubscription
}

func (x *TableOciRegionGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.IdentityService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			regions, err := session.IdentityClient.ListRegions(ctx)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := identity.ListRegionSubscriptionsRequest{
				TenancyId: &session.TenancyID,
			}

			subscribedRegions, err := session.IdentityClient.ListRegionSubscriptions(ctx, request)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, region := range regions.Items {
				isSubscribed := false
				for _, subscribedRegion := range subscribedRegions.Items {
					if *region.Name == *subscribedRegion.RegionName {
						resultChannel <- ociRegion{region, subscribedRegion}
						isSubscribed = true
						break
					}
				}
				if isSubscribed {
					continue
				}
				resultChannel <- ociRegion{region, identity.RegionSubscription{}}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, err)

		},
	}
}

func (x *TableOciRegionGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableOciRegionGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("is_home_region").ColumnType(schema.ColumnTypeBool).Description("Indicates if the region is the home region or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Description("The region subscription status.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key").ColumnType(schema.ColumnTypeString).Description("The key of the region.").Build(),
	}
}

func (x *TableOciRegionGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableOciIdentityAvailabilityDomainGenerator{}),
	}
}
