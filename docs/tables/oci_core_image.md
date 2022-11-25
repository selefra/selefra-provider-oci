# Table: oci_core_image

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| operating_system | string | X | √ | The image's operating system. | 
| agent_features | json | X | √ | Oracle Cloud Agent features supported on the image. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| display_name | string | X | √ | A user-friendly name for the image. It does not have to be unique, and it's changeable. | 
| id | string | X | √ | The OCID of the image. | 
| create_image_allowed | bool | X | √ | Indicates whether instances launched with this image can be used to create new images. | 
| operating_system_version | string | X | √ | The image's operating system version. | 
| size_in_mbs | int | X | √ | The boot volume size for an instance launched from this image. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| lifecycle_state | string | X | √ | The image's current state. | 
| time_created | timestamp | X | √ | The date and time the image was created. | 
| launch_mode | string | X | √ | Specifies the configuration mode for launching virtual machine (VM) instances. | 
| launch_options | json | X | √ | LaunchOptions Options for tuning the compatibility and performance of VM shapes. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| base_image_id | string | X | √ | The OCID of the image originally used to launch the instance. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 


