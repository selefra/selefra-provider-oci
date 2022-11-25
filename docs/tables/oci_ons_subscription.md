# Table: oci_ons_subscription

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ | The OCID of the subscription. | 
| created_time | timestamp | X | √ | The time when this subscription was created. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| lifecycle_state | string | X | √ | The lifecycle state of the subscription. | 
| etag | string | X | √ | Used for optimistic concurrency control. | 
| title | string | X | √ | Title of the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| topic_id | string | X | √ | The OCID of the associated topic. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| endpoint | string | X | √ | A locator that corresponds to the subscription protocol. | 
| protocol | string | X | √ | The protocol used for the subscription. | 
| delivery_policy | json | X | √ | Delivery Policy of the subscription. | 
| tags | json | X | √ | A map of tags for the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 


