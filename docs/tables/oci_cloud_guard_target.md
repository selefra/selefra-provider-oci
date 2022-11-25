# Table: oci_cloud_guard_target

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| time_updated | timestamp | X | √ | The date and time the target was updated. | 
| inherited_by_compartments | json | X | √ | List of inherited compartments. | 
| title | string | X | √ | Title of the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| target_resource_id | string | X | √ | Resource ID which the target uses to monitor. | 
| description | string | X | √ | The target description. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| id | string | X | √ | OCID for target. | 
| time_created | timestamp | X | √ | The date and time the target was created. | 
| target_resource_type | string | X | √ | Possible type of targets(compartment/HCMCloud/ERPCloud). | 
| tags | json | X | √ | A map of tags for the resource. | 
| lifecyle_details | string | X | √ | A message describing the current state in more detail. | 
| recipe_count | int | X | √ | Total number of recipes attached to target. | 
| target_detector_recipes | json | X | √ | List of detector recipes associated with target. | 
| target_responder_recipes | json | X | √ | List of responder recipes associated with target. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| system_tags | json | X | √ | System tags for resource. System tags can be viewed by users, but can only be created by the system. | 
| name | string | X | √ | Target display name. | 
| lifecycle_state | string | X | √ | The current state of the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 


