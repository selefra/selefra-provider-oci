# Table: oci_objectstorage_bucket

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| approximate_size | int | X | √ | The approximate total size in bytes of all objects in the bucket. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| storage_tier | string | X | √ | The storage tier type assigned to the bucket. | 
| time_created | timestamp | X | √ | The date and time the bucket was created. | 
| name | string | X | √ | The name of the bucket. | 
| namespace | string | X | √ | The Object Storage namespace in which the bucket lives. | 
| approximate_count | int | X | √ | The approximate number of objects in the bucket. | 
| is_read_only | bool | X | √ | Whether or not this bucket is read only. | 
| kms_key_id | string | X | √ | The OCID of a master encryption key used to call the Key Management service to generate a data encryption key or to encrypt or decrypt a data encryption key. | 
| object_events_enabled | bool | X | √ | Whether or not events are emitted for object state changes in this bucket. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| etag | string | X | √ | The entity tag (ETag) for the bucket. | 
| public_access_type | string | X | √ | The type of public access enabled on this bucket. | 
| versioning | string | X | √ | The versioning status on the bucket. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| id | string | X | √ | The OCID of the bucket. | 
| created_by | string | X | √ | The OCID of the user who created the bucket. | 
| object_lifecycle_policy_etag | string | X | √ | The entity tag (ETag) for the live object lifecycle policy on the bucket. | 
| replication_enabled | bool | X | √ | Whether or not this bucket is a replication source. | 
| metadata | json | X | √ | Arbitrary string keys and values for user-defined metadata. | 
| object_lifecycle_policy | json | X | √ | Specifies the object lifecycle policy for the bucket. | 


