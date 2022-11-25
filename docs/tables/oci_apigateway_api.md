# Table: oci_apigateway_api

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ | A map of tags for the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| time_created | timestamp | X | √ | The time this resource was created. | 
| lifecycle_details | string | X | √ | A message describing the current lifecycleState. | 
| time_updated | timestamp | X | √ | The time this resource was last updated. | 
| id | string | X | √ | The OCID of the resource. | 
| validation_results | json | X | √ | Status of each feature available from the API. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| display_name | string | X | √ | A user-friendly name. Does not have to be unique, and it's changeable. | 
| lifecycle_state | string | X | √ | The current state of the API. | 
| specification_type | string | X | √ | Type of API Specification file. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| title | string | X | √ | Title of the resource. | 


