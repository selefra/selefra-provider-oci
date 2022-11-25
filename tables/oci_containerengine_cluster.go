package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/containerengine"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
)

type TableOciContainerengineClusterGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciContainerengineClusterGenerator{}

func (x *TableOciContainerengineClusterGenerator) GetTableName() string {
	return "oci_containerengine_cluster"
}

func (x *TableOciContainerengineClusterGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciContainerengineClusterGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciContainerengineClusterGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciContainerengineClusterGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.ContainerEngineService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildContainerEngineClusterFilters()

			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.ContainerEngineClient.ListClusters(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, cluster := range response.Items {
					resultChannel <- cluster

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

func buildContainerEngineClusterFilters() containerengine.ListClustersRequest {
	request := containerengine.ListClustersRequest{}

	return request
}
func isValidContainerEngineClusterLifecycleState(state string) bool {
	stateType := containerengine.ClusterLifecycleStateEnum(state)
	switch stateType {
	case containerengine.ClusterLifecycleStateActive, containerengine.ClusterLifecycleStateCreating, containerengine.ClusterLifecycleStateDeleted, containerengine.ClusterLifecycleStateDeleting, containerengine.ClusterLifecycleStateFailed, containerengine.ClusterLifecycleStateUpdating:
		return true
	}
	return false
}

func (x *TableOciContainerengineClusterGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciContainerengineClusterGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("image_policy_config_enabled").ColumnType(schema.ColumnTypeBool).Description("Whether the image verification policy is enabled. Defaults to false. If set to true, the images will be verified against the policy at runtime.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoints").ColumnType(schema.ColumnTypeJSON).Description("Endpoints served up by the cluster masters.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint_config").ColumnType(schema.ColumnTypeJSON).Description("The network configuration for access to the Cluster control plane.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).Description("Metadata about the cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name. It does not have to be unique, and it is changeable.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The state of the cluster masters.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vcn_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the virtual cloud network (VCN) in which the cluster exists.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("options").ColumnType(schema.ColumnTypeJSON).Description("Optional attributes for the cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the cluster.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the KMS key to be used as the master encryption key for Kubernetes secret encryption.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kubernetes_version").ColumnType(schema.ColumnTypeString).Description("The version of Kubernetes running on the cluster masters.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_details").ColumnType(schema.ColumnTypeString).Description("Additional information about the current 'lifecycleState'.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("available_kubernetes_upgrades").ColumnType(schema.ColumnTypeJSON).Description("Available Kubernetes versions to which the clusters masters may be upgraded.").Build(),
	}
}

func (x *TableOciContainerengineClusterGenerator) GetSubTables() []*schema.Table {
	return nil
}
