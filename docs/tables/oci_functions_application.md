# Table: oci_functions_application

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| time_updated | timestamp | X | √ | The time the application was updated. | 
| title | string | X | √ | Title of the resource. | 
| id | string | X | √ | The OCID of the application. | 
| syslog_url | string | X | √ | A syslog URL to which to send all function logs. Supports tcp, udp, and tcp+tls. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| display_name | string | X | √ | The display name of the application. | 
| lifecycle_state | string | X | √ | The current state of the application. | 
| time_created | timestamp | X | √ | The time the application was created. | 
| config | json | X | √ | Application configuration for functions in this application. | 
| subnet_ids | json | X | √ | The OCIDs of the subnets in which to run functions in the application. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| tags | json | X | √ | A map of tags for the resource. | 


