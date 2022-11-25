# Table: oci_autoscaling_auto_scaling_configuration

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resource | json | X | √ | The resource details of this AutoScalingConfiguration. | 
| tags | json | X | √ | A map of tags for the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| title | string | X | √ | Title of the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| id | string | X | √ | The OCID of the autoscaling configuration. | 
| is_enabled | bool | X | √ | Indicates whether the autoscaling configuration is enabled. | 
| time_created | timestamp | X | √ | The date and time the AutoScalingConfiguration was created. | 
| cool_down_in_seconds | int | X | √ | The minimum period of time to wait between scaling actions. | 
| max_resource_count | int | X | √ | The maximum number of resources to scale out to. | 
| min_resource_count | int | X | √ | The minimum number of resources to scale in to. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| display_name | string | X | √ | A user-friendly name. | 
| policies | json | X | √ | Autoscaling policy definitions for the autoscaling configuration. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 


