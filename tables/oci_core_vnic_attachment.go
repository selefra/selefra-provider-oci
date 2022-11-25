package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/core"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-oci/oci_client"
)

type TableOciCoreVnicAttachmentGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCoreVnicAttachmentGenerator{}

func (x *TableOciCoreVnicAttachmentGenerator) GetTableName() string {
	return "oci_core_vnic_attachment"
}

func (x *TableOciCoreVnicAttachmentGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCoreVnicAttachmentGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCoreVnicAttachmentGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCoreVnicAttachmentGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			session, err := oci_client.CoreComputeService(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := core.ListVnicAttachmentsRequest{
				CompartmentId: &session.TenancyID,
				Limit:         pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.ComputeClient.ListVnicAttachments(ctx, request)
				if err != nil {

					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, attachment := range response.Items {
					resultChannel <- attachment

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

func (x *TableOciCoreVnicAttachmentGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciCoreVnicAttachmentGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the VNIC attachment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("skip_source_dest_check").ColumnType(schema.ColumnTypeBool).Description("Whether the source/destination check is disabled on the VNIC. Defaults to `false`, which means the check is performed.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the subnet to create the VNIC in.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nsg_ids").ColumnType(schema.ColumnTypeJSON).Description("A list of the OCIDs of the network security groups that the VNIC belongs to.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vnic_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name for the VNIC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_primary").ColumnType(schema.ColumnTypeBool).Description("Whether the VNIC is the primary VNIC (the VNIC that is automatically created and attached during instance launch).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nic_index").ColumnType(schema.ColumnTypeInt).Description("The physical network interface card (NIC) the VNIC uses.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vnic_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the VNIC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hostname_label").ColumnType(schema.ColumnTypeString).Description("The hostname for the VNIC's primary private IP.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vlan_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the VLAN to create the VNIC in.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Description("A user-friendly name for the VNIC attachment.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifecycle_state").ColumnType(schema.ColumnTypeString).Description("The current state of the VNIC attachment. Possible values include: 'ATTACHING', 'ATTACHED', 'DETACHING', 'DETACHED'.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_ip").ColumnType(schema.ColumnTypeString).Description("The private IP address of the primary `privateIp` object on the VNIC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_ip").ColumnType(schema.ColumnTypeString).Description("The public IP address of the VNIC, if one is assigned.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability_domain").ColumnType(schema.ColumnTypeString).Description("The availability domain of the instance.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the VNIC attachment was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mac_address").ColumnType(schema.ColumnTypeString).Description("The MAC address of the VNIC.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vlan_tag").ColumnType(schema.ColumnTypeInt).Description("The OCID of the VNIC.").Build(),
	}
}

func (x *TableOciCoreVnicAttachmentGenerator) GetSubTables() []*schema.Table {
	return nil
}
