# Table: oci_identity_tag_namespace

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| title | string | X | √ | Title of the resource. | 
| name | string | X | √ | The name of the tag namespace. It must be unique across all tag namespaces in the tenancy and cannot be changed. | 
| id | string | X | √ | The OCID of the tag namespace. | 
| lifecycle_state | string | X | √ | The tagnamespace's current state. | 
| time_created | timestamp | X | √ | Date and time the tagNamespace was created. | 
| description | string | X | √ | The description you assign to the tag namespace. | 
| tags | json | X | √ | A map of tags for the resource. | 
| is_retired | bool | X | √ | Whether the tag namespace is retired. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 


