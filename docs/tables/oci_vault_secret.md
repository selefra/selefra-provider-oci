# Table: oci_vault_secret

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| title | string | X | √ | Title of the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| key_id | string | X | √ | The OCID of the master encryption key that is used to encrypt the secret. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| current_version_number | int | X | √ | The version number of the secret that's currently in use. | 
| secret_rules | json | X | √ | A list of rules that control how the secret is used and managed. | 
| time_of_current_version_expiry | string | X | √ | An optional property indicating when the current secret version will expire. | 
| metadata | json | X | √ | Additional metadata that you can use to provide context about how to use the secret or during rotation or other administrative tasks. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| id | string | X | √ | The OCID of the secret. | 
| description | string | X | √ | A brief description of the secret. | 
| vault_id | string | X | √ | The OCID of the Vault in which the secret exists. | 
| lifecycle_details | string | X | √ | Additional information about the secret's current lifecycle state. | 
| time_created | string | X | √ | A property indicating when the secret was created. | 
| time_of_deletion | string | X | √ | An optional property indicating when to delete the secret. | 
| name | string | X | √ | The name of the secret. | 
| lifecycle_state | string | X | √ | The current lifecycle state of the secret. | 


