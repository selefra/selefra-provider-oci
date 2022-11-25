# Table: oci_core_volume_backup_policy

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| display_name | string | X | √ | A user-friendly name for volume backup policy. | 
| time_created | timestamp | X | √ | The date and time the volume backup policy was created. | 
| destination_region | string | X | √ | The paired destination region for copying scheduled backups to. | 
| schedules | json | X | √ | The collection of schedules that this policy will apply. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| id | string | X | √ | The OCID of the volume backup policy. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 


