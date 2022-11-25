# Table: oci_core_boot_volume_backup

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| boot_volume_id | string | X | √ | The OCID of the boot volume. | 
| time_created | timestamp | X | √ | The date and time the boot volume backup was created. | 
| size_in_gbs | int | X | √ | The size of the boot volume, in GBs. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| display_name | string | X | √ | A user-friendly name for the boot volume backup. | 
| expiration_time | timestamp | X | √ | The date and time the volume backup will expire and be automatically deleted. | 
| image_id | string | X | √ | The image OCID used to create the boot volume the backup is taken from. | 
| type | string | X | √ | The type of a volume backup. | 
| source_type | string | X | √ | Specifies whether the backup was created manually, or via scheduled backup policy. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| time_request_received | timestamp | X | √ | The date and time the request to create the boot volume backup was received. | 
| source_boot_volume_backup_id | string | X | √ | The OCID of the source boot volume backup. | 
| unique_size_in_gbs | int | X | √ | The size used by the backup, in GBs. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| id | string | X | √ | The OCID of the boot volume backup. | 
| lifecycle_state | string | X | √ | The current state of a boot volume backup. | 
| kms_key_id | string | X | √ | The OCID of the Key Management master encryption assigned to the boot volume backup. | 
| system_tags | json | X | √ | System tags for this resource. | 
| title | string | X | √ | Title of the resource. | 


