# Table: oci_database_db_home

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| time_created | timestamp | X | √ | The date and time the database home was created. | 
| kms_key_id | string | X | √ | The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| id | string | X | √ | The OCID of the database home. | 
| lifecycle_state | string | X | √ | The current state of the database home. | 
| db_system_id | string | X | √ | The OCID of the DB system. | 
| database_software_image_id | string | X | √ | The database software image OCID. | 
| db_version | string | X | √ | The oracle database version. | 
| lifecycle_details | string | X | √ | Additional information about the current lifecycle state. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| display_name | string | X | √ | The user-friendly name for the database home. It does not have to be unique. | 
| db_home_location | string | X | √ | The location of the oracle database home. | 
| vm_cluster_id | string | X | √ | The OCID of the VM cluster. | 
| tags | json | X | √ | A map of tags for the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| last_patch_history_entry_id | string | X | √ | The OCID of the last patch history. | 
| one_off_patches | json | X | √ | List of one-off patches for database homes. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| title | string | X | √ | Title of the resource. | 


