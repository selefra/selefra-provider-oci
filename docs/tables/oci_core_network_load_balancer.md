# Table: oci_core_network_load_balancer

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| lifecycle_details | string | X | √ | A message describing the current state in more detail. | 
| network_security_group_ids | json | X | √ | An array of network security groups OCIDs. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| lifecycle_state | string | X | √ | The current state of the network load balancer. | 
| time_updated | timestamp | X | √ | The date and time the network load balancer was created. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| time_created | timestamp | X | √ | The date and time the network load balancer was created. | 
| is_private | bool | X | √ | Whether the network load balancer has a virtual cloud network-local (private) IP address. | 
| listeners | json | X | √ | Listeners associated with the network load balancer. | 
| system_tags | json | X | √ | System tags for resource. System tags can be viewed by users, but can only be created by the system. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| display_name | string | X | √ | A user-friendly name. Does not have to be unique. | 
| id | string | X | √ | The OCID of the network load balancer. | 
| subnet_id | string | X | √ | The subnet in which the network load balancer is spawned OCIDs. | 
| is_preserve_source_destination | bool | X | √ | When enabled, the skipSourceDestinationCheck parameter is automatically enabled on the load balancer VNIC. | 
| network_load_balancer_health | json | X | √ | The overall health status of the network load balancer. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 


