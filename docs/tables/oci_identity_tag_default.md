# Table: oci_identity_tag_default

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| title | string | X | √ | Title of the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| tag_definition_id | string | X | √ | The OCID of the tag definition. | 
| is_required | bool | X | √ | If you specify that a value is required, a value is set during resource creation (either by the user creating the resource or another tag default). If no value is set, resource creation is blocked.If the `isRequired` flag is set to true, the value is set during resource creation.If the `isRequired` flag is set to false, the value you enter is set during resource creation. | 
| value | string | X | √ | The default value for the tag definition. | 
| lifecycle_state | string | X | √ | The tag default's current state. | 
| time_created | timestamp | X | √ | Date and time the tagDefault was created. | 
| id | string | X | √ | The OCID of the tag default. | 
| tag_definition_name | string | X | √ | The name used in the tag definition. | 
| tag_namespace_id | string | X | √ | The OCID of the tag namespace that contains the tag definition. | 


