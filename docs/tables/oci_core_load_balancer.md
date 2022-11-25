# Table: oci_core_load_balancer

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ | A map of tags for the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| id | string | X | √ | The OCID of the load balancer. | 
| display_name | string | X | √ | A user-friendly name of the load balancer. | 
| backend_sets | json | X | √ | The configuration of a load balancer backend set. | 
| hostnames | json | X | √ | A hostname resource associated with a load balancer for use by one or more listeners. | 
| ssl_cipher_suites | json | X | √ | The configuration details of an SSL cipher suite. | 
| subnet_ids | json | X | √ | An array of subnet OCIDs. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| shape_name | string | X | √ | A template that determines the total pre-provisioned bandwidth (ingress plus egress). | 
| is_private | bool | X | √ | Whether the load balancer has a VCN-local (private) IP address. | 
| ip_addresses | json | X | √ | An array of IP addresses. | 
| network_security_group_ids | json | X | √ | An array of NSG OCIDs associated with the load balancer. | 
| shape_details | json | X | √ | The configuration details to update load balancer to a different shape. | 
| system_tags | json | X | √ | System tags for this resource. | 
| title | string | X | √ | Title of the resource. | 
| lifecycle_state | string | X | √ | The load balancer's current state. | 
| time_created | timestamp | X | √ | The date and time the load balancer was created. | 
| path_route_sets | json | X | √ | A named set of path route rules. | 
| rule_sets | json | X | √ | A named set of rules associated with a load balancer. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| certificates | json | X | √ | The configuration details of a certificate bundle. | 
| listeners | json | X | √ | The listener's configuration. | 
| routing_policies | json | X | √ | A named ordered list of routing rules that is applied to a listener. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 


