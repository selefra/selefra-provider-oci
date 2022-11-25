package tables

import (
	"context"
	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/objectstorage"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
)

type TableOciObjectstorageBucketGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciObjectstorageBucketGenerator{}

func (x *TableOciObjectstorageBucketGenerator) GetTableName() string {
	return "oci_objectstorage_bucket"
}

func (x *TableOciObjectstorageBucketGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciObjectstorageBucketGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciObjectstorageBucketGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciObjectstorageBucketGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			region := taskClient.(*oci_client.OciClient).Region

			session, err := oci_client.ObjectStorageService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			nameSpace, err := getNamespace(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := objectstorage.ListBucketsRequest{
				CompartmentId: &session.TenancyID,
				NamespaceName: &nameSpace.Value,
				Limit:         pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			pagesLeft := true
			for pagesLeft {
				response, err := session.ObjectStorageClient.ListBuckets(ctx, request)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)
				}

				for _, bucketSummary := range response.Items {
					resultChannel <- bucketInfo{region, bucketSummary}

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

func (x *TableOciObjectstorageBucketGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciObjectstorageBucketGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("approximate_size").ColumnType(schema.ColumnTypeInt).Description("The approximate total size in bytes of all objects in the bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_tier").ColumnType(schema.ColumnTypeString).Description("The storage tier type assigned to the bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the bucket was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).Description("The Object Storage namespace in which the bucket lives.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("approximate_count").ColumnType(schema.ColumnTypeInt).Description("The approximate number of objects in the bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_read_only").ColumnType(schema.ColumnTypeBool).Description("Whether or not this bucket is read only.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_id").ColumnType(schema.ColumnTypeString).Description("The OCID of a master encryption key used to call the Key Management service to generate a data encryption key or to encrypt or decrypt a data encryption key.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("object_events_enabled").ColumnType(schema.ColumnTypeBool).Description("Whether or not events are emitted for object state changes in this bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("freeform_tags").ColumnType(schema.ColumnTypeJSON).Description("Free-form tags for resource. This tags can be applied by any user with permissions on the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).Description("The entity tag (ETag) for the bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_access_type").ColumnType(schema.ColumnTypeString).Description("The type of public access enabled on this bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("versioning").ColumnType(schema.ColumnTypeString).Description("The versioning status on the bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("defined_tags").ColumnType(schema.ColumnTypeJSON).Description("Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Description("A map of tags for the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the compartment in Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by").ColumnType(schema.ColumnTypeString).Description("The OCID of the user who created the bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("object_lifecycle_policy_etag").ColumnType(schema.ColumnTypeString).Description("The entity tag (ETag) for the live object lifecycle policy on the bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication_enabled").ColumnType(schema.ColumnTypeBool).Description("Whether or not this bucket is a replication source.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).Description("Arbitrary string keys and values for user-defined metadata.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("object_lifecycle_policy").ColumnType(schema.ColumnTypeJSON).Description("Specifies the object lifecycle policy for the bucket.").Build(),
	}
}

func (x *TableOciObjectstorageBucketGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableOciObjectstorageObjectGenerator{}),
	}
}
