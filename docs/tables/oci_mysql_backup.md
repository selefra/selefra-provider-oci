# Table: oci_mysql_backup

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| mysql_version | string | X | √ | The version of the DB System used for backup. | 
| retention_in_days | int | X | √ | Number of days to retain this backup. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| time_created | timestamp | X | √ | The time the backup record was created. | 
| backup_type | string | X | √ | The type of backup. | 
| data_storage_size_in_gbs | int | X | √ | Initial size of the data volume in GiBs. | 
| db_system_snapshot | json | X | √ | Snapshot of the DbSystem details at the time of the backup. | 
| title | string | X | √ | Title of the resource. | 
| time_updated | timestamp | X | √ | The time at which the backup was updated. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| display_name | string | X | √ | A user-supplied display name for the backup. | 
| id | string | X | √ | The OCID of the backup. | 
| lifecycle_state | string | X | √ | The current state of the Backup. | 
| creation_type | string | X | √ | If the backup was created automatically, or by a manual request. | 
| description | string | X | √ | A user-supplied description of the backup. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| db_system_id | string | X | √ | The OCID of the DB System the Backup is associated with. | 
| backup_size_in_gbs | int | X | √ | The size of the backup in GiBs. | 
| lifecycle_details | string | X | √ | Additional information about the current lifecycleState. | 
| shape_name | string | X | √ | The shape of the DB System instance used for backup. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 


