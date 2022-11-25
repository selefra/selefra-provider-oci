# Table: oci_functions_function

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| display_name | string | X | √ | The display name of the function. | 
| id | string | X | √ | The OCID of the function. | 
| memory_in_mbs | int | X | √ | Maximum usable memory for the function (MiB). | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 
| timeout_in_seconds | int | X | √ | Timeout for executions of the function. Value in seconds. | 
| config | json | X | √ | The function configuration. Overrides application configuration. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| lifecycle_state | string | X | √ | The current state of the function. | 
| image_digest | string | X | √ | The image digest for the version of the image that will be pulled when invoking this function. If no value is specified, the digest currently associated with the image in the OCI Registry will be used. | 
| time_updated | timestamp | X | √ | The time the function was updated. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| application_id | string | X | √ | The OCID of the application the function belongs to. | 
| image | string | X | √ | The qualified name of the Docker image to use in the function, including the image tag. The image should be in the OCI Registry that is in the same region as the function itself. | 
| invoke_endpoint | string | X | √ | The base https invoke URL to set on a client in order to invoke a function. This URL will never change over the lifetime of the function and can be cached. | 
| time_created | timestamp | X | √ | The time the function was created. | 
| trace_config | json | X | √ | The trace configuration of the function. | 


