# Table: oci_core_boot_volume_replica

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| total_data_transferred_in_gbs | int | X | √ | The total size of the data transferred from the source boot volume to the boot volume replica, in GBs. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| id | string | X | √ | The boot volume replica's Oracle ID (OCID). | 
| display_name | string | X | √ | A user-friendly name. | 
| boot_volume_id | string | X | √ | The OCID of the source boot volume. | 
| time_created | timestamp | X | √ | The date and time the boot volume replica was created. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| availability_domain | string | X | √ | The availability domain of the boot volume replica. | 
| image_id | string | X | √ | The image OCID used to create the boot volume the replica is replicated from. | 
| time_last_synced | timestamp | X | √ | The date and time the boot volume replica was last synced from the source boot volume. | 
| lifecycle_state | string | X | √ | The current state of a boot volume replica. | 
| size_in_gbs | int | X | √ | The size of the source boot volume, in GBs. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 


