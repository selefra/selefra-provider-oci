# Table: oci_identity_policy

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| time_created | timestamp | X | √ | Date and time the policy was created. | 
| lifecycle_state | string | X | √ | The policy's current state. | 
| statements | json | X | √ | An array of one or more policy statements written in the policy language. | 
| description | string | X | √ | The description you assign to the policy. Does not have to be unique, and it's changeable. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| name | string | X | √ | The name you assign to the policy during creation. The name must be unique across all policies in the tenancy and cannot be changed. | 
| id | string | X | √ | The OCID of the policy. | 
| inactive_status | int | X | √ | The detailed status of INACTIVE lifecycleState. | 
| version_date | timestamp | X | √ | The version of the policy. If null or set to an empty string, when a request comes in for authorization, the policy will be evaluated according to the current behavior of the services at that moment. If set to a particular date (YYYY-MM-DD), the policy will be evaluated according to the behavior of the services on that date. | 
| title | string | X | √ | Title of the resource. | 


