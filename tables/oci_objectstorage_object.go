package tables

import (
	"context"
	"github.com/selefra/selefra-utils/pkg/pointer"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/objectstorage"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableOciObjectstorageObjectGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciObjectstorageObjectGenerator{}

func (x *TableOciObjectstorageObjectGenerator) GetTableName() string {
	return "oci_objectstorage_object"
}

func (x *TableOciObjectstorageObjectGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciObjectstorageObjectGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciObjectstorageObjectGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciObjectstorageObjectGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			region := taskClient.(*oci_client.OciClient).Region

			bucketName := *task.ParentRawResult.(bucketInfo).Name

			objectNameSpace, err := getNamespace(ctx, clientMeta, taskClient, task)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			session, err := oci_client.ObjectStorageService(ctx, clientMeta, taskClient, task)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			request := objectstorage.ListObjectsRequest{
				BucketName:    &bucketName,
				NamespaceName: &objectNameSpace.Value,
				Fields:        pointer.ToStringPointer("name,size,etag,timeCreated,md5,timeModified,storageTier,archivalState"),
				Limit:         pointer.ToIntPointer(1000),
				RequestMetadata: common.RequestMetadata{
					RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
				},
			}

			response, err := session.ObjectStorageClient.ListObjects(ctx, request)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			for _, objectSummary := range response.Objects {
				resultChannel <- objectInfo{bucketName, objectNameSpace.Value, region, objectSummary}
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

type bucketInfo struct {
	Region string
	objectstorage.BucketSummary
}
type nameSpace struct {
	Value string
}
type objectInfo struct {
	BucketName string
	Namespace  string
	Region     string
	objectstorage.ObjectSummary
}

func getNamespace(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*nameSpace, error) {

	session, err := oci_client.ObjectStorageService(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}
	request := objectstorage.GetNamespaceRequest{}

	response, err := session.ObjectStorageClient.GetNamespace(ctx, request)
	if err != nil {
		return nil, err
	}
	name := &nameSpace{
		Value: *response.Value,
	}

	return name, err
}

func (x *TableOciObjectstorageObjectGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildRegionList()
}

func (x *TableOciObjectstorageObjectGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("The name of the object.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("content_disposition").ColumnType(schema.ColumnTypeString).Description("The Content-Disposition header.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("content_type").ColumnType(schema.ColumnTypeString).Description("The Content-Type header.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_tier").ColumnType(schema.ColumnTypeString).Description("The storage tier that the object is stored in.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_of_archival").ColumnType(schema.ColumnTypeTimestamp).Description("Time that the object is returned to the archived state. This field is only present for restored objects.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("opc_meta").ColumnType(schema.ColumnTypeJSON).Description("The user-defined metadata for the object.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("content_language").ColumnType(schema.ColumnTypeString).Description("The Content-Language header.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("content_range").ColumnType(schema.ColumnTypeString).Description("Content-Range header for range requests.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).Description("The current entity tag (ETag) for the object.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_not_modified").ColumnType(schema.ColumnTypeBool).Description("Flag to indicate whether or not the object was modified. If this is true, the getter for the object itself will return null.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("opc_multipart_md5").ColumnType(schema.ColumnTypeString).Description("Base-64 representation of the multipart object hash. Only applicable to objects uploaded using multipart upload.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size").ColumnType(schema.ColumnTypeInt).Description("Size of the object in bytes.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).Description("Title of the resource.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bucket_name").ColumnType(schema.ColumnTypeString).Description("The name of the bucket.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).Description("The Object Storage namespace used for the request.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cache_control").ColumnType(schema.ColumnTypeString).Description("The Cache-Control header.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("content_encoding").ColumnType(schema.ColumnTypeString).Description("The Content-Encoding header.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("md5").ColumnType(schema.ColumnTypeString).Description("Base64-encoded MD5 hash of the object data.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the object was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_modified").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time the object was modified.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("archival_state").ColumnType(schema.ColumnTypeString).Description("Archival state of an object. This field is set only for objects in Archive tier.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expires").ColumnType(schema.ColumnTypeTimestamp).Description("The date and time after which the object is no longer cached by a browser, proxy, or other caching entity.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_id").ColumnType(schema.ColumnTypeString).Description("The version ID of the object requested.").Build(),
	}
}

func (x *TableOciObjectstorageObjectGenerator) GetSubTables() []*schema.Table {
	return nil
}
