# Table: oci_core_internet_gateway

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| display_name | string | X | √ | A user-friendly name. | 
| vcn_id | string | X | √ | The OCID of the VCN the internet gateway belongs to. | 
| is_enabled | bool | X | √ | Whether the gateway is enabled. | 
| lifecycle_state | string | X | √ | The internet gateway's current state. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| id | string | X | √ | The internet gateway's Oracle ID (OCID). | 
| time_created | timestamp | X | √ | The date and time the internet gateway was created. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| tags | json | X | √ | A map of tags for the resource. | 


