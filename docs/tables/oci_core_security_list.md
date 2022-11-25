# Table: oci_core_security_list

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| egress_security_rules | json | X | √ | Rules for allowing egress IP packets. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| lifecycle_state | string | X | √ | The security list's current state. | 
| ingress_security_rules | json | X | √ | Rules for allowing ingress IP packets. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| id | string | X | √ | The security list's Oracle Cloud ID (OCID). | 
| vcn_id | string | X | √ | The OCID of the VCN the security list belongs to. | 
| time_created | timestamp | X | √ | The date and time the security list was created. | 
| tags | json | X | √ | A map of tags for the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| display_name | string | X | √ | A user-friendly name. Does not have to be unique, and it's changeable. | 


