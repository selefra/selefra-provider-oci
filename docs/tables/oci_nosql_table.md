# Table: oci_nosql_table

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| lifecycle_state | string | X | √ | The state of a table. | 
| is_auto_reclaimable | bool | X | √ | True if this table can be reclaimed after an idle period. | 
| tags | json | X | √ | A map of tags for the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| name | string | X | √ | Immutable human-friendly table name. | 
| id | string | X | √ | Unique identifier that is immutable. | 
| time_of_expiration | timestamp | X | √ | If lifecycleState is INACTIVE, indicates when this table will be automatically removed. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| system_tags | json | X | √ | System tags for resource. System tags can be viewed by users, but can only be created by the system. | 
| time_created | timestamp | X | √ | The time the the table was created. | 
| ddl_statement | string | X | √ | A DDL statement representing the schema. | 
| lifecycle_details | string | X | √ | A message describing the current state in more detail. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| title | string | X | √ | Title of the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| time_updated | timestamp | X | √ | The time the the table was updated. | 
| schema | json | X | √ | The schema of the table. | 
| table_limits | json | X | √ | Various limit for the table. | 
| region | string | X | √ | The OCI region in which the resource is located. | 


