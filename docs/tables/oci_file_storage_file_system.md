# Table: oci_file_storage_file_system

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| time_created | timestamp | X | √ | The date and time the file system was created. | 
| is_clone_parent | bool | X | √ | Specifies whether the file system has been cloned. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| id | string | X | √ | The OCID of the file system. | 
| lifecycle_details | string | X | √ | Additional information about the current 'lifecycleState'. | 
| source_details | json | X | √ | Source information for the file system. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| display_name | string | X | √ | A user-friendly name. It does not have to be unique, and it is changeable. | 
| lifecycle_state | string | X | √ | The current state of the file system. | 
| availability_domain | string | X | √ | The availability domain the file system is in. | 
| is_hydrated | bool | X | √ | Specifies whether the data has finished copying from the source to the clone. | 
| kms_key_id | string | X | √ | The OCID of the KMS key used to encrypt the encryption keys associated with this file system. | 
| metered_bytes | int | X | √ | The number of bytes consumed by the file system. | 
| tags | json | X | √ | A map of tags for the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 


