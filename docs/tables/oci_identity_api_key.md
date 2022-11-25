# Table: oci_identity_api_key

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| key_id | string | X | √ | An Oracle-assigned identifier for the key. | 
| user_id | string | X | √ | The OCID of the user the key belongs to. | 
| user_name | string | X | √ | The name of the user the key belongs to. | 
| lifecycle_state | string | X | √ | The API key's current state. | 
| fingerprint | string | X | √ | The key's fingerprint. | 
| key_value | string | X | √ | The key's value. | 
| time_created | timestamp | X | √ | Date and time the `ApiKey` object was created. | 
| inactive_status | int | X | √ | The detailed status of INACTIVE lifecycleState. | 
| title | string | X | √ | Title of the resource. | 


