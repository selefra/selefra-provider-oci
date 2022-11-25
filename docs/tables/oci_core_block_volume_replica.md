# Table: oci_core_block_volume_replica

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ | The block volume replica's Oracle ID (OCID). | 
| size_in_gbs | int | X | √ | The size of the source block volume, in GBs. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| display_name | string | X | √ | A user-friendly name. | 
| lifecycle_state | string | X | √ | The current state of a block volume replica. | 
| time_created | timestamp | X | √ | The date and time the block volume replica was created. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| block_volume_id | string | X | √ | The OCID of the source block volume. | 
| availability_domain | string | X | √ | The availability domain of the block volume replica. | 
| time_last_synced | timestamp | X | √ | The date and time the block volume replica was last synced from the source block volume. | 
| total_data_transferred_in_gbs | int | X | √ | The total size of the data transferred from the source block volume to the block volume replica, in GBs. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 


