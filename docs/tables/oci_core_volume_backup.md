# Table: oci_core_volume_backup

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| size_in_gbs | int | X | √ | The size of the volume, in GBs. | 
| unique_size_in_mbs | int | X | √ | The size used by the backup, in MBs. | 
| volume_id | string | X | √ | The OCID of the volume. | 
| title | string | X | √ | Title of the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| expiration_time | timestamp | X | √ | The date and time the volume backup will expire and be automatically deleted. | 
| kms_key_id | string | X | √ | The OCID of the Key Management key which is the master encryption key for the volume backup. | 
| size_in_mbs | int | X | √ | The size of the volume in MBs. | 
| source_volume_backup_id | string | X | √ | The OCID of the source volume backup. | 
| time_created | timestamp | X | √ | The date and time the volume backup was created. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| system_tags | json | X | √ | System tags to volume by the service. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| id | string | X | √ | The OCID of the volume backup. | 
| display_name | string | X | √ | A user-friendly name for the volume backup. | 
| lifecycle_state | string | X | √ | The current state of a volume backup. | 
| source_type | string | X | √ | Specifies whether the backup was created manually, or via scheduled backup policy. | 
| type | string | X | √ | The type of a volume backup. | 
| time_request_received | timestamp | X | √ | The date and time the request to create the volume backup was received. | 
| unique_size_in_gbs | int | X | √ | The size used by the backup, in GBs. | 
| tags | json | X | √ | A map of tags for the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 


