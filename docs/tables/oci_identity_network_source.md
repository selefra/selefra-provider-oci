# Table: oci_identity_network_source

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| time_created | timestamp | X | √ | Date and time the etwork source was created. | 
| inactive_status | int | X | √ | The detailed status of INACTIVE lifecycleState. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| name | string | X | √ | The name you assign to the network source during creation. | 
| id | string | X | √ | The OCID of the network source. | 
| public_source_list | json | X | √ | A list of allowed public IP addresses and CIDR ranges. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| lifecycle_state | string | X | √ | The network source object's current state. | 
| services | json | X | √ | A list of services allowed to make on-behalf-of requests. | 
| virtual_source_list | json | X | √ | A list of allowed VCN OCID and IP range pairs. | 
| description | string | X | √ | The description you assign to the network source. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 


