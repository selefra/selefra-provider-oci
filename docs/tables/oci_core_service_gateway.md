# Table: oci_core_service_gateway

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| vcn_id | string | X | √ | The OCID of the VCN the service gateway belongs to. | 
| tags | json | X | √ | A map of tags for the resource. | 
| id | string | X | √ | The OCID of the service gateway. | 
| block_traffic | bool | X | √ | Specifies whether the service gateway blocks traffic through it. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| title | string | X | √ | Title of the resource. | 
| lifecycle_state | string | X | √ | The service gateway's current state. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| display_name | string | X | √ | A user-friendly name of the service gateway. | 
| route_table_id | string | X | √ | The OCID of the route table the service gateway is using. | 
| time_created | timestamp | X | √ | The date and time the service gateway was created | 
| services | json | X | √ | List of the Service objects enabled for this service gateway. | 


