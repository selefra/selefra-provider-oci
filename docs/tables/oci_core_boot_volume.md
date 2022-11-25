# Table: oci_core_boot_volume

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kms_key_id | string | X | √ | The OCID of the Key Management key which is the master encryption key for the boot volume. | 
| size_in_mbs | int | X | √ | The size of the boot volume in MBs. | 
| volume_backup_policy_assignment_id | string | X | √ | The OCID of the volume backup policy assignment. | 
| system_tags | json | X | √ | System tags for resource. System tags can be viewed by users, but can only be created by the system. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| volume_group_id | string | X | √ | The OCID of the source volume group. | 
| availability_domain | string | X | √ | The availability domain of the boot volume. | 
| is_hydrated | bool | X | √ | Specifies whether the boot volume's data has finished copying from the source boot volume or boot volume backup. | 
| vpus_per_gb | int | X | √ | The number of volume performance units (VPUs) that will be applied to this boot volume per GB,representing the Block Volume service's elastic performance options. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| title | string | X | √ | Title of the resource. | 
| display_name | string | X | √ | A user-friendly name. | 
| auto_tuned_vpus_per_gb | int | X | √ | The number of Volume Performance Units per GB that this boot volume is effectively tuned to when it's idle. | 
| is_auto_tune_enabled | bool | X | √ | Specifies whether the auto-tune performance is enabled for this boot volume. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| id | string | X | √ | The OCID of the boot volume. | 
| time_created | timestamp | X | √ | The date and time the boot volume was created. | 
| volume_backup_policy_id | string | X | √ | The OCID of the volume backup policy that has been assigned to the volume. | 
| tags | json | X | √ | A map of tags for the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| lifecycle_state | string | X | √ | The current state of a boot volume. | 
| size_in_gbs | int | X | √ | The size of the boot volume in GBs. | 
| source_details | json | X | √ | The boot volume source, either an existing volume in the same availability domain or a volume backup. | 


