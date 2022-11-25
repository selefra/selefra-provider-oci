# Table: oci_logging_log

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| name | string | X | √ | A user-friendly name. | 
| lifecycle_state | string | X | √ | The log object state. | 
| time_created | timestamp | X | √ | Time the resource was created. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| log_group_id | string | X | √ | The OCID of the log group. | 
| is_enabled | bool | X | √ | Whether or not this resource is currently enabled. | 
| configuration | json | X | √ | Log object configuration. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| tags | json | X | √ | A map of tags for the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| id | string | X | √ | The OCID of the log. | 
| log_type | string | X | √ | The logType that the log object is for, whether custom or service. | 
| time_last_modified | timestamp | X | √ | Time the resource was last modified. | 
| retention_duration | int | X | √ | Log retention duration in 30-day increments (30, 60, 90 and so on). | 


