# Table: oci_identity_tenancy

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ | The OCID of the tenancy. | 
| retention_period_days | int | X | √ | The retention period setting, specified in days. | 
| home_region_key | string | X | √ | The region key for the tenancy's home region. | 
| upi_idcs_compatibility_layer_endpoint | string | X | √ | Url which refers to the UPI IDCS compatibility layer endpoint configured for this Tenant's home region. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| name | string | X | √ | The name of the tenancy. | 
| description | string | X | √ | The description of the tenancy. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 


