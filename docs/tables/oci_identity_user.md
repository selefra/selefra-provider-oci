# Table: oci_identity_user

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| can_use_auth_tokens | bool | X | √ | Indicates if the user can use SWIFT passwords/auth tokens. | 
| email_verified | bool | X | √ | Whether the email address has been validated. | 
| identity_provider_id | string | X | √ | The OCID of the `IdentityProvider` this user belongs to. | 
| inactive_status | int | X | √ | Applicable only if the user's `lifecycleState` is INACTIVE. A 16-bit value showing the reason why the user is inactive. 0: SUSPENDED; 1: DISABLED; 2: BLOCKED (the user has exceeded the maximum number of failed login attempts for the Console) | 
| tags | json | X | √ | A map of tags for the resource. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| id | string | X | √ | The OCID of the user. | 
| time_created | timestamp | X | √ | Date and time the user was created. | 
| description | string | X | √ | The description assigned to the user. | 
| is_mfa_activated | bool | X | √ | The user's current state. | 
| email | string | X | √ | The email address you assign to the user. | 
| user_groups | json | X | √ | List of groups associated with the user. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| user_type | string | X | √ | Type of the user. Value can be IDCS or IAM. Oracle Identity Cloud Service(IDCS) users authenticate through single sign-on and can be granted access to all services included in your account. IAM users can access Oracle Cloud Infrastructure services, but not all Cloud Platform services. | 
| can_use_api_keys | bool | X | √ | Indicates if the user can use API keys. | 
| can_use_smtp_credentials | bool | X | √ | Indicates if the user can use SMTP passwords. | 
| external_identifier | string | X | √ | Identifier of the user in the identity provider. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| title | string | X | √ | Title of the resource. | 
| name | string | X | √ | The user's login for the Console. | 
| lifecycle_state | string | X | √ | The user's current state. | 
| can_use_console_password | bool | X | √ | Indicates if the user can log in to the console. | 
| can_use_customer_secret_keys | bool | X | √ | Indicates if the user can use SigV4 symmetric keys. | 
| can_use_o_auth2_client_credentials | bool | X | √ | Indicates if the user can use OAuth2 credentials and tokens. | 


