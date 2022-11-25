# Table: oci_core_boot_volume_attachment

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| encryption_in_transit_type | string | X | √ | The type of the encryption in transit for the boot volume. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| lifecycle_state | string | X | √ | The current state of the boot volume attachment. | 
| availability_domain | string | X | √ | The availability domain of an instance. | 
| time_created | timestamp | X | √ | The date and time the boot volume was created. | 
| is_pv_encryption_in_transit_enabled | bool | X | √ | Whether in-transit encryption for the boot volume's paravirtualized attachment is enabled or not. | 
| id | string | X | √ | The OCID of the boot volume attachment. | 
| display_name | string | X | √ | A user-friendly name. Does not have to be unique, and it cannot be changed. | 
| boot_volume_id | string | X | √ | The OCID of the boot volume. | 
| instance_id | string | X | √ | The OCID of the instance the boot volume is attached to. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 


