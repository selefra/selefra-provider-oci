# Table: oci_core_volume_attachment

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| time_created | timestamp | X | √ | The date and time the volume was created. | 
| iscsi_login_state | string | X | √ | The iscsi login state of the volume attachment. | 
| is_read_only | bool | X | √ | Whether the attachment was created in read-only mode. | 
| is_pv_encryption_in_transit_enabled | bool | X | √ | Whether in-transit encryption for the data volume's paravirtualized attachment is enabled or not. | 
| title | string | X | √ | Title of the resource. | 
| id | string | X | √ | The OCID of the volume attachment. | 
| volume_id | string | X | √ | The OCID of the volume. | 
| availability_domain | string | X | √ | The availability domain of an instance. | 
| device | string | X | √ | The device name. | 
| is_shareable | bool | X | √ | Whether the attachment should be created in shareable mode. | 
| is_multipath | bool | X | √ | Whether the attachment is multipath or not. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| instance_id | string | X | √ | The OCID of the instance the volume is attached to. | 
| lifecycle_state | string | X | √ | The current state of the volume attachment. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| display_name | string | X | √ | A user-friendly name. Does not have to be unique, and it cannot be changed. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 


