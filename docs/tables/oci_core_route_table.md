# Table: oci_core_route_table

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| time_created | timestamp | X | √ | The date and time the route table was created. | 
| route_rules | json | X | √ | The collection of rules for routing destination IPs to network devices. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| display_name | string | X | √ | A user-friendly name. Does not have to be unique, and it's changeable. | 
| id | string | X | √ | The route table's Oracle ID (OCID). | 
| vcn_id | string | X | √ | The OCID of the VCN the route table list belongs to. | 
| lifecycle_state | string | X | √ | The route table's current state. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 


