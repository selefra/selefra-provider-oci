# Table: oci_core_image_custom

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| title | string | X | √ | Title of the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| display_name | string | X | √ | A user-friendly name for the image. It does not have to be unique, and it's changeable. | 
| create_image_allowed | bool | X | √ | Indicates whether instances launched with this image can be used to create new images. | 
| launch_mode | string | X | √ | Specifies the configuration mode for launching virtual machine (VM) instances. | 
| operating_system_version | string | X | √ | The image's operating system version. | 
| id | string | X | √ | The OCID of the image. | 
| lifecycle_state | string | X | √ | The image's current state. | 
| agent_features | json | X | √ | Oracle Cloud Agent features supported on the image. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| time_created | timestamp | X | √ | The date and time the image was created. | 
| launch_options | json | X | √ | LaunchOptions Options for tuning the compatibility and performance of VM shapes. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| base_image_id | string | X | √ | The OCID of the image originally used to launch the instance. | 
| operating_system | string | X | √ | The image's operating system. | 
| size_in_mbs | int | X | √ | The boot volume size for an instance launched from this image. | 
| tags | json | X | √ | A map of tags for the resource. | 


