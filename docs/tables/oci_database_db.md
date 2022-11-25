# Table: oci_database_db

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| db_name | string | X | √ | The database name. | 
| database_software_image_id | string | X | √ | The database software image OCID. | 
| pdb_name | string | X | √ | The name of the pluggable database. | 
| title | string | X | √ | Title of the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| kms_key_id | string | X | √ | The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. | 
| ncharacter_set | string | X | √ | The national character set for the database. | 
| source_database_point_in_time_recovery_timestamp | timestamp | X | √ | Point in time recovery timeStamp of the source database at which cloned database system is cloned from the source database system. | 
| db_unique_name | string | X | √ | A system-generated name for the database to ensure uniqueness within an oracle data guard group. | 
| id | string | X | √ | The OCID of the database. | 
| db_home_id | string | X | √ | The OCID of the database home. | 
| db_system_id | string | X | √ | The OCID of the DB system. | 
| db_workload | string | X | √ | The database workload type. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| time_created | timestamp | X | √ | The date and time the database was created. | 
| vm_cluster_id | string | X | √ | The OCID of the vm cluster. | 
| connection_strings | json | X | √ | The connection strings used to connect to the oracle database. | 
| db_backup_config | json | X | √ | Database backup configuration details. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| lifecycle_state | string | X | √ | The current state of the database. | 
| character_set | string | X | √ | The character set for the database. | 
| last_backup_timestamp | timestamp | X | √ | The date and time when the latest database backup was created. | 
| lifecycle_details | string | X | √ | Additional information about the current lifecycle state. | 


