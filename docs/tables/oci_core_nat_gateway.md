# Table: oci_core_nat_gateway

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| time_created | timestamp | X | √ | The date and time the NAT gateway was created. | 
| nat_ip | ip | X | √ | The IP address associated with the NAT gateway. | 
| tags | json | X | √ | A map of tags for the resource. | 
| display_name | string | X | √ | A user-friendly name of the NAT gateway. | 
| lifecycle_state | string | X | √ | The NAT gateway's current state. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| id | string | X | √ | The OCID of the NAT gateway. | 
| vcn_id | string | X | √ | The OCID of the VCN the NAT gateway belongs to. | 
| block_traffic | bool | X | √ | Specifies whether the NAT gateway blocks traffic through it. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| public_ip_id | string | X | √ | The OCID of the public IP address associated with the NAT gateway. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| title | string | X | √ | Title of the resource. | 


