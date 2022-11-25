# Table: oci_resourcemanager_stack

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ | A map of tags for the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| id | string | X | √ | Unique identifier of the specified stack. | 
| terraform_version | string | X | √ | The version of Terraform specified for the stack. | 
| time_drift_last_checked | timestamp | X | √ | The date and time when the drift detection was last executed. | 
| display_name | string | X | √ | Human-readable display name for the stack. | 
| stack_drift_status | string | X | √ | Drift status of the stack. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| time_created | timestamp | X | √ | The date and time when the stack was created. | 
| variables | json | X | √ | Terraform variables associated with this resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| title | string | X | √ | Title of the resource. | 
| lifecycle_state | string | X | √ | The current lifecycle state of the stack. | 
| description | string | X | √ | General description of the stack. | 
| config_source | json | X | √ | The version of Terraform specified for the stack. | 


