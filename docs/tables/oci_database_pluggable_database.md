# Table: oci_database_pluggable_database

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| open_mode | string | X | √ | The mode that pluggableDatabase is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend. | 
| connection_strings | json | X | √ | The connection strings used to connect to the oracle pluggable database. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| time_created | timestamp | X | √ | The date and time the pluggable database was created. | 
| container_database_id | string | X | √ | The OCID of the CDB. | 
| is_restricted | bool | X | √ | The restricted mode of pluggableDatabase. If a pluggableDatabase is opened in restricted mode, the user needs both Create a session and restricted session privileges to connect to it. | 
| lifecycle_details | string | X | √ | Detailed message for the lifecycle state. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| pdb_name | string | X | √ | The name for the pluggable database. The name is unique in the context of a Database. The name must begin with an alphabetic character and can contain a maximum of thirty alphanumeric characters. Special characters are not permitted. The pluggable database name should not be same as the container database name. | 
| id | string | X | √ | The OCID of the pluggable database. | 
| lifecycle_state | string | X | √ | The current state of the pluggable database. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 


