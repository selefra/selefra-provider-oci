# Table: oci_objectstorage_object

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | The name of the object. | 
| content_disposition | string | X | √ | The Content-Disposition header. | 
| content_type | string | X | √ | The Content-Type header. | 
| storage_tier | string | X | √ | The storage tier that the object is stored in. | 
| time_of_archival | timestamp | X | √ | Time that the object is returned to the archived state. This field is only present for restored objects. | 
| opc_meta | json | X | √ | The user-defined metadata for the object. | 
| content_language | string | X | √ | The Content-Language header. | 
| content_range | string | X | √ | Content-Range header for range requests. | 
| etag | string | X | √ | The current entity tag (ETag) for the object. | 
| is_not_modified | bool | X | √ | Flag to indicate whether or not the object was modified. If this is true, the getter for the object itself will return null. | 
| opc_multipart_md5 | string | X | √ | Base-64 representation of the multipart object hash. Only applicable to objects uploaded using multipart upload. | 
| size | int | X | √ | Size of the object in bytes. | 
| title | string | X | √ | Title of the resource. | 
| bucket_name | string | X | √ | The name of the bucket. | 
| namespace | string | X | √ | The Object Storage namespace used for the request. | 
| cache_control | string | X | √ | The Cache-Control header. | 
| content_encoding | string | X | √ | The Content-Encoding header. | 
| md5 | string | X | √ | Base64-encoded MD5 hash of the object data. | 
| time_created | timestamp | X | √ | The date and time the object was created. | 
| time_modified | timestamp | X | √ | The date and time the object was modified. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| archival_state | string | X | √ | Archival state of an object. This field is set only for objects in Archive tier. | 
| expires | timestamp | X | √ | The date and time after which the object is no longer cached by a browser, proxy, or other caching entity. | 
| version_id | string | X | √ | The version ID of the object requested. | 


