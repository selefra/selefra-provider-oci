# Table: oci_cloud_guard_responder_recipe

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | Responder recipe description. | 
| effective_responder_rules | json | X | √ | List of responder rules for the responder type for recipe. | 
| system_tags | json | X | √ | System tags for resource. System tags can be viewed by users, but can only be created by the system. | 
| title | string | X | √ | Title of the resource. | 
| name | string | X | √ | Display name of responder recipe. | 
| id | string | X | √ | OCID for responder recipe. | 
| source_responder_recipe_id | string | X | √ | Recipe OCID of the source recipe to be cloned. | 
| time_updated | timestamp | X | √ | The date and time the responder recipe was updated. | 
| owner | string | X | √ | Owner of responder recipe. | 
| time_created | timestamp | X | √ | The date and time the responder recipe was created. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| lifecycle_state | string | X | √ | The current state of the responder recipe. | 
| lifecycle_details | string | X | √ | A message describing the current state in more detail. | 
| responder_rules | json | X | √ | List of responder rules for the responder type for recipe. | 
| tags | json | X | √ | A map of tags for the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 


