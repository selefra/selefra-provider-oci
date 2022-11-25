# Table: oci_dns_zone

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| is_protected | bool | X | √ | A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed. | 
| nameservers | json | X | √ | The authoritative nameservers for the zone. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| name | string | X | √ | The name of the zone. | 
| self | string | X | √ | The canonical absolute URL of the resource. | 
| title | string | X | √ | Title of the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| lifecycle_state | string | X | √ | The current state of the zone resource. | 
| time_created | timestamp | X | √ | The date and time the zone was created. | 
| serial | int | X | √ | The current serial of the zone. As seen in the zone's SOA record. | 
| version | string | X | √ | Version is the never-repeating, totally-orderable, version of the zone, from which the serial field of the zone's SOA record is derived. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 
| id | string | X | √ | The OCID of the zone. | 
| zone_type | string | X | √ | The type of the zone. Must be either `PRIMARY` or `SECONDARY`. `SECONDARY` is only supported for GLOBAL zones. | 
| scope | string | X | √ | The scope of the zone. | 
| view_id | string | X | √ | The OCID of the private view containing the zone. | 
| external_masters | json | X | √ | External master servers for the zone. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 


