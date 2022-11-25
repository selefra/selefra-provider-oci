# Table: oci_mysql_configuration

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| title | string | X | √ | Title of the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| time_updated | timestamp | X | √ | The date and time the Configuration was last updated. | 
| description | string | X | √ | User-provided data about the Configuration. | 
| type | string | X | √ | The Configuration type, DEFAULT or CUSTOM. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| tags | json | X | √ | A map of tags for the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| display_name | string | X | √ | The display name of the Configuration. | 
| parent_configuration_id | string | X | √ | The OCID of the Configuration from which this Configuration is derived. | 
| time_created | timestamp | X | √ | The date and time the Configuration was created. | 
| variables | string | X | √ | User controllable service variables of the Configuration. | 
| id | string | X | √ | The OCID of the Configuration. | 
| shape_name | string | X | √ | The name of the associated Shape. | 
| lifecycle_state | string | X | √ | The current state of the Configuration. | 


