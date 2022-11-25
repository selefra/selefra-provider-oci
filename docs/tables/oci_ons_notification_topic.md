# Table: oci_ons_notification_topic

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ | The name of the topic. | 
| topic_id | string | X | √ | The OCID of the topic. | 
| tags | json | X | √ | A map of tags for the resource. | 
| description | string | X | √ | The description of the topic. | 
| short_topic_id | string | X | √ | A unique short topic Id. This is used only for SMS subscriptions. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| etag | string | X | √ | Used for optimistic concurrency control. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| lifecycle_state | string | X | √ | The lifecycle state of the topic. | 
| time_created | timestamp | X | √ | The time the topic was created. | 
| api_endpoint | string | X | √ | The endpoint for managing subscriptions or publishing messages to the topic. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| title | string | X | √ | Title of the resource. | 


