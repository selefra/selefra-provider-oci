# Table: oci_core_local_peering_gateway

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | A user-friendly name. Does not have to be unique, and it's changeable. | 
| vcn_id | string | X | √ | The OCID of the VCN that uses the LPG. | 
| peer_advertised_cidr | cidr | X | √ | The smallest aggregate CIDR that contains all the CIDR routes advertised by the VCN at the other end of the peering from this LPG. | 
| peer_advertised_cidr_details | json | X | √ | The specific ranges of IP addresses available on or via the VCN at the other end of the peering from this LPG. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| title | string | X | √ | Title of the resource. | 
| peering_status | string | X | √ | Whether the LPG is peered with another LPG. | 
| peering_status_details | string | X | √ | Additional information regarding the peering status. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| id | string | X | √ | The LPG's Oracle ID | 
| lifecycle_state | string | X | √ | The LPG's current lifecycle state. | 
| time_created | timestamp | X | √ | The date and time the LPG was created. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| route_table_id | string | X | √ | The OCID of the route table the LPG is using. | 
| is_cross_tenancy_peering | bool | X | √ | Whether the VCN at the other end of the peering is in a different tenancy. | 
| tags | json | X | √ | A map of tags for the resource. | 


