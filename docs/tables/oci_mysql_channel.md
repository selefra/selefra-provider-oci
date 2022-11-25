# Table: oci_mysql_channel

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | A user-supplied description of the backup. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| id | string | X | √ | The OCID of the Channel. | 
| is_enabled | bool | X | √ | Whether the Channel has been enabled by the user. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| title | string | X | √ | Title of the resource. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| tags | json | X | √ | A map of tags for the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| lifecycle_details | string | X | √ | A message describing the state of the Channel. | 
| time_updated | timestamp | X | √ | The time the Channel was last updated. | 
| source | json | X | √ | Parameters detailing how to provision the source for the given Channel. | 
| target | json | X | √ | Parameters detailing how to provision the target for the given Channel. | 
| display_name | string | X | √ | The user-friendly name for the Channel. It does not have to be unique. | 
| lifecycle_state | string | X | √ | The current state of the Channel. | 
| time_created | timestamp | X | √ | The date and time the Channel was created. | 


