# Table: oci_identity_auth_token

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ | The OCID of the auth token. | 
| user_name | string | X | √ | The name of the user the auth token belongs to. | 
| time_expires | timestamp | X | √ | Date and time when this auth token will expire. | 
| inactive_status | int | X | √ | The detailed status of INACTIVE lifecycleState. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| user_id | string | X | √ | The OCID of the user the auth token belongs to. | 
| token | string | X | √ | The auth token. The value is available only in the response for `CreateAuthToken`, and not for `ListAuthTokens` or `UpdateAuthToken`. | 
| lifecycle_state | string | X | √ | The token's current state. | 
| time_created | timestamp | X | √ | Date and time the `AuthToken` object was created. | 
| description | string | X | √ | The description you assign to the auth token. | 
| title | string | X | √ | Title of the resource. | 


