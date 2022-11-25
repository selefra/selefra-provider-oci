# Table: oci_cloud_guard_managed_list

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | ManagedList display name. | 
| source_managed_list_id | string | X | √ | OCID of the source managed list. | 
| list_type | string | X | √ | Type of the list. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| tags | json | X | √ | A map of tags for the resource. | 
| id | string | X | √ | OCID for managed list. | 
| lifecycle_state | string | X | √ | The current state of the managed list. | 
| is_editable | bool | X | √ | If this list is editable or not. | 
| time_updated | timestamp | X | √ | The date and time the managed list was updated. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| system_tags | json | X | √ | System tags for resource. System tags can be viewed by users, but can only be created by the system. | 
| time_created | timestamp | X | √ | The date and time the managed list was created. | 
| description | string | X | √ | Managed list description. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| title | string | X | √ | Title of the resource. | 
| feed_provider | string | X | √ | Provider of the feed. | 
| lifecyle_details | string | X | √ | A message describing the current state in more detail. | 
| list_items | json | X | √ | List of managed list item. | 


