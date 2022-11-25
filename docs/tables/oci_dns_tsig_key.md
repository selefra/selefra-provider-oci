# Table: oci_dns_tsig_key

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | A globally unique domain name identifying the key for a given pair of hosts. | 
| time_created | timestamp | X | √ | The date and time the resource was created. | 
| lifecycle_state | string | X | √ | The current state of the resource. | 
| title | string | X | √ | Title of the resource. | 
| tags | json | X | √ | A map of tags for the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| algorithm | string | X | √ | TSIG key algorithms are encoded as domain names, but most consist of only one non-empty label, which is not required to be explicitly absolute. | 
| secret | string | X | √ | A base64 string encoding the binary shared secret. | 
| self | string | X | √ | The canonical absolute URL of the resource. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| id | string | X | √ | The OCID of the resource. | 
| time_updated | timestamp | X | √ | The date and time the resource was last updated. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 


