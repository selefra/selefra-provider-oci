# Table: oci_core_public_ip_pool

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| title | string | X | √ | Title of the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| id | string | X | √ | The OCID of the public IP pool. | 
| lifecycle_state | string | X | √ | The public IP pool's current state. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| display_name | string | X | √ | A user-friendly name. Does not have to be unique. | 
| time_created | timestamp | X | √ | The date and time the public IP pool was created. | 
| cidr_blocks | json | X | √ | The CIDR blocks added to this pool. This could be all or a portion of a BYOIP CIDR block. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 


