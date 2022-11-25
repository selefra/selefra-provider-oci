# Table: oci_mysql_configuration_custom

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| display_name | string | X | √ | The display name of the configuration. | 
| type | string | X | √ | The configuration type, DEFAULT or CUSTOM. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 
| lifecycle_state | string | X | √ | The current state of the configuration. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| title | string | X | √ | Title of the resource. | 
| id | string | X | √ | The OCID of the configuration. | 
| parent_configuration_id | string | X | √ | The OCID of the configuration from which this configuration is derived. | 
| description | string | X | √ | User-provided data about the configuration. | 
| variables | string | X | √ | User controllable service variables of the configuration. | 
| time_created | timestamp | X | √ | The date and time the configuration was created. | 
| shape_name | string | X | √ | The name of the associated shape. | 
| time_updated | timestamp | X | √ | The date and time the configuration was last updated. | 


