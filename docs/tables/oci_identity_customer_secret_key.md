# Table: oci_identity_customer_secret_key

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| title | string | X | √ | Title of the resource. | 
| id | string | X | √ | The OCID of the secret key. | 
| lifecycle_state | string | X | √ | The secret key's current state. | 
| inactive_status | int | X | √ | The detailed status of INACTIVE lifecycleState. | 
| time_created | timestamp | X | √ | Date and time the CustomerSecretKey object was created. | 
| time_expires | timestamp | X | √ | Date and time when this password will expire. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| display_name | string | X | √ | The displayName you assign to the secret key. | 
| user_id | string | X | √ | The OCID of the user the password belongs to. | 
| user_name | string | X | √ | The name of the user the password belongs to. | 


