# Table: oci_identity_authentication_policy

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| is_lowercase_characters_required | bool | X | √ | At least one lower case character required. | 
| is_numeric_characters_required | bool | X | √ | At least one numeric character required. | 
| is_special_characters_required | bool | X | √ | At least one special character required. | 
| is_uppercase_characters_required | bool | X | √ | At least one uppercase character required. | 
| is_username_containment_allowed | bool | X | √ | User name is allowed to be part of the password. | 
| minimum_password_length | int | X | √ | Minimum password length required. | 
| network_source_ids | string | X | √ | List of IP ranges from which users can sign in to the Console. | 


