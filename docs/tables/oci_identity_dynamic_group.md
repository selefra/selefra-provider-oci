# Table: oci_identity_dynamic_group

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| name | string | X | √ | The name you assign to the group during creation. | 
| id | string | X | √ | The OCID of the group. | 
| lifecycle_state | string | X | √ | The group's current state. | 
| inactive_status | int | X | √ | The detailed status of INACTIVE lifecycleState. | 
| matching_rule | string | X | √ | A rule string that defines which instance certificates will be matched. | 
| description | string | X | √ | The description you assign to the group. | 
| time_created | timestamp | X | √ | Date and time the group was created, in the format defined by RFC3339. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 


