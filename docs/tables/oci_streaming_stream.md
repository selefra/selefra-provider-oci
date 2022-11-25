# Table: oci_streaming_stream

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | The name of the stream. | 
| time_created | timestamp | X | √ | The date and time the stream was created. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| title | string | X | √ | Title of the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| partitions | int | X | √ | The number of partitions in the stream. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| id | string | X | √ | The OCID of the stream. | 
| lifecycle_state_details | string | X | √ | Any additional details about the current state of the stream. | 
| messages_endpoint | string | X | √ | The endpoint to use when creating the StreamClient to consume or publish messages in the stream. | 
| lifecycle_state | string | X | √ | The current state of the stream. | 
| retention_in_hours | int | X | √ | The retention period of the stream, in hours. This property is read-only. | 
| stream_pool_id | string | X | √ | The OCID of the stream pool that contains the stream. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 


