# Table: oci_analytics_instance

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | The analytics instance's optional description. | 
| email_notification | string | X | √ | The email address receiving notifications. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| time_created | timestamp | X | √ | The date and time the instance was created. | 
| capacity_type | string | X | √ | The analytics instance's capacity model to use. | 
| capacity_value | int | X | √ | The analytics instance's capacity value selected. | 
| private_access_channels | json | X | √ | The private access channels of the analytics instance. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| time_updated | timestamp | X | √ | The date and time the instance was last updated. | 
| lifecycle_state | string | X | √ | The analytics instance's current state. | 
| feature_set | string | X | √ | The analytics instance's feature set. | 
| license_type | string | X | √ | The license used for the service. | 
| service_url | string | X | √ | The URL of the Analytics service. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| name | string | X | √ | A user-friendly name. Does not have to be unique, and it's changeable. | 
| network_endpoint_details | json | X | √ | The base representation of a network endpoint. | 
| vanity_url_details | json | X | √ | The vanity url configuration details of the analytic instance. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| id | string | X | √ | The analytics instance's Oracle ID (OCID). | 


