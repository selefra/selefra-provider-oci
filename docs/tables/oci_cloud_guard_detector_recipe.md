# Table: oci_cloud_guard_detector_recipe

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| time_updated | timestamp | X | √ | The date and time the detector recipe was updated. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| system_tags | json | X | √ | System tags for resource. System tags can be viewed by users, but can only be created by the system. | 
| title | string | X | √ | Title of the resource. | 
| id | string | X | √ | Ocid for detector recipe. | 
| lifecycle_state | string | X | √ | The current state of the detector recipe. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| name | string | X | √ | DisplayName of detector recipe. | 
| tags | json | X | √ | A map of tags for the resource. | 
| time_created | timestamp | X | √ | The date and time the detector recipe was created. | 
| owner | string | X | √ | Owner of detector recipe. | 
| detector | string | X | √ | Type of detector. | 
| detector_rules | json | X | √ | List of detector rules for the detector type for recipe. | 
| effective_detector_rules | json | X | √ | List of detector rules for the detector type for recipe. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| source_detector_recipe_id | string | X | √ | Recipe Ocid of the Source Recipe to be cloned. | 
| description | string | X | √ | Detector recipe description. | 


