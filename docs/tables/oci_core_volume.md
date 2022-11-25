# Table: oci_core_volume

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ | A map of tags for the resource. | 
| display_name | string | X | √ | A user-friendly name. | 
| lifecycle_state | string | X | √ | The current state of a volume. | 
| size_in_mbs | int | X | √ | The size of the volume in MBs. | 
| volume_backup_policy_id | string | X | √ | The OCID of the volume backup policy that has been assigned to the volume. | 
| system_tags | json | X | √ | System tags to volume by the service. | 
| source_details | json | X | √ | The volume source, either an existing volume in the same availability domain or a volume backup. | 
| id | string | X | √ | The OCID of the volume. | 
| auto_tuned_vpus_per_gb | int | X | √ | The number of Volume Performance Units per GB that this volume is effectively tuned to when it's idle. | 
| is_hydrated | bool | X | √ | Specifies whether the cloned volume's data has finished copying from the source volume or backup. | 
| kms_key_id | string | X | √ | The OCID of the Key Management key which is the master encryption key for the volume. | 
| size_in_gbs | int | X | √ | The size of the volume in GBs. | 
| availability_domain | string | X | √ | The availability domain of the volume. | 
| time_created | timestamp | X | √ | The date and time the volume was created. | 
| is_auto_tune_enabled | bool | X | √ | Specifies whether the auto-tune performance is enabled for this volume. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| title | string | X | √ | Title of the resource. | 
| compartment_id | string | X | √ | ColumnDescriptionCompartment | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| volume_group_id | string | X | √ | The OCID of the source volume group. | 
| volume_backup_policy_assignment_id | string | X | √ | The OCID of the volume backup policy assignment. | 
| vpus_per_gb | int | X | √ | The number of volume performance units (VPUs) that will be applied to this volume per GB,representing the Block Volume service's elastic performance options. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 


