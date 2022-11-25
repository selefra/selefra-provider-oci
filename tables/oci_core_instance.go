package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/core"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableOciCoreInstanceGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCoreInstanceGenerator{}

func (x *TableOciCoreInstanceGenerator) GetTableName() string {
	return "oci_core_instance"
}

func (x *TableOciCoreInstanceGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCoreInstanceGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCoreInstanceGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCoreInstanceGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.CoreComputeService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := buildCoreInstanceFilters()
			request.CompartmentId = &session.TenancyID
			request.Limit = pointer.ToIntPointer(1000)
			request.RequestMetadata = common.RequestMetadata{
				RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.ComputeClient.ListInstances(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, instance := range response.Items {
					resultChannel <- instance

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

func buildCoreInstanceFilters() core.ListInstancesRequest {
	request := core.ListInstancesRequest{}

	return request
}

func (x *TableOciCoreInstanceGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciCoreInstanceGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("shape_config_baseline_ocpu_utilization").ColumnType(schema.ColumnTypeString).Description("The baseline OCPU utilization for a subcore burstable VM instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dedicated_vm_host_id").ColumnType(schema.ColumnTypeString).Description("The OCID of dedicated VM host.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipxe_script").ColumnType(schema.ColumnTypeString).Description("When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shape_config_memory_in_gbs").ColumnType(schema.ColumnTypeFloat).Description("The total amount of memory available to the instance, in gigabytes.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_config").ColumnType(schema.ColumnTypeJSON).Description("The platform configuration for the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_details").ColumnType(schema.ColumnTypeJSON).Description("Contains the details of the source image for the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("launch_mode").ColumnType(schema.ColumnTypeString).Description("Specifies the configuration mode for launching virtual machine (VM) instances.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shape_config_max_vnic_attachments").ColumnType(schema.ColumnTypeInt).Description("The maximum number of VNIC attachments for the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shape_config_ocpus").ColumnType(schema.ColumnTypeFloat).Description("The total number of OCPUs available to the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("extended_metadata").ColumnType(schema.ColumnTypeJSON).Description("Additional metadata key/value pairs that user provided to instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the instance was created").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("capacity_reservation_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compute capacity reservation this instance is launched under.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shape_config_gpus").ColumnType(schema.ColumnTypeInt).Description("The number of GPUs available to the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shape_config_networking_bandwidth_in_gbps").ColumnType(schema.ColumnTypeFloat).Description("The networking bandwidth available to the instance, in gigabits per second.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_config").ColumnType(schema.ColumnTypeJSON).Description("Options for defining the availability of a VM instance after a maintenance event that impacts the underlying hardware.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shape_config_local_disks_total_size_in_gbs").ColumnType(schema.ColumnTypeFloat).Description("The aggregate size of all local disks, in gigabytes.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_maintenance_reboot_due").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the instance is expected to be stopped/started. After that time if instance hasn't been rebooted, Oracle will reboot the instance within 24 hours of the due time.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("agent_config").ColumnType(schema.ColumnTypeJSON).Description("Options for the Oracle Cloud Agent software running on the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shape_config").ColumnType(schema.ColumnTypeJSON).Description("The shape configuration for an instance. The shape configuration determines the resources allocated to an instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("system_tags").ColumnType(schema.ColumnTypeJSON).Description("Tags added to instances by the service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_domain").ColumnType(schema.ColumnTypeString).Description("The availability domain the instance is running in.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shape").ColumnType(schema.ColumnTypeString).Description("The shape of the instance. The shape determines the number of CPUs and the amount of memory allocated to the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("launch_options").ColumnType(schema.ColumnTypeJSON).Description("LaunchOptions Options for tuning the compatibility and performance of VM shapes.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_options").ColumnType(schema.ColumnTypeJSON).Description("Optional mutable instance options.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name. Does not have to be unique, and it's changeable.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fault_domain").ColumnType(schema.ColumnTypeString).Description("The name of the fault domain the instance is running in. A fault domain is a grouping of hardware and infrastructure within an availability domain.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shape_config_local_disks").ColumnType(schema.ColumnTypeInt).Description("The number of local disks available to the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).Description("Custom metadata that you provided to instance.").Build(),
	}
}

func (x *TableOciCoreInstanceGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableOciCoreInstanceMetricCpuUtilizationDailyGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciCoreInstanceMetricCpuUtilizationGenerator{}),
		table_schema_generator.GenTableSchema(&TableOciCoreInstanceMetricCpuUtilizationHourlyGenerator{}),
	}
}
