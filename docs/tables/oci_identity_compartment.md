# Table: oci_identity_compartment

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | The name assigned to the compartment during creation | 
| id | string | X | √ | The OCID of the compartment. | 
| time_created | timestamp | X | √ | Date and time the user was created. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| lifecycle_state | string | X | √ | The compartment's current state. | 
| description | string | X | √ | The description you assign to the compartment. | 
| inactive_status | int | X | √ | The detailed status of INACTIVE lifecycleState | 
| is_accessible | bool | X | √ | Indicates whether or not the compartment is accessible for the user making the request. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 


