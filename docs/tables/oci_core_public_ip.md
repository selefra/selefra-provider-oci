# Table: oci_core_public_ip

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| assigned_entity_id | string | X | √ | The OCID of the entity the public IP is assigned to, or in the process of being assigned to. | 
| assigned_entity_type | string | X | √ | The type of entity the public IP is assigned to, or in the process of being assigned to. | 
| time_created | timestamp | X | √ | The date and time the public IP was created. | 
| availability_domain | string | X | √ | The public IP's availability domain. This property is set only for ephemeral public IPs that are assigned to a private IP. | 
| ip_address | ip | X | √ | The public IP address of the publicIp object. | 
| public_ip_pool_id | string | X | √ | The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pool object created in the current tenancy. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| id | string | X | √ | The public IP's Oracle ID (OCID). | 
| lifecycle_state | string | X | √ | The public IP's current state. | 
| scope | string | X | √ | Indicates whether the public IP is regional or specific to a particular availability domain. | 
| display_name | string | X | √ | A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. | 
| lifetime | string | X | √ | Defines when the public IP is deleted and released back to Oracle's public IP pool. | 


