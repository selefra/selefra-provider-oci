# Table: oci_database_software_image

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| included_patches_summary | string | X | √ | The patches included in the image and the version of the image. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| display_name | string | X | √ | The user-friendly name for the database software image. The name does not have to be unique. | 
| id | string | X | √ | The OCID of the database software image. | 
| time_created | timestamp | X | √ | The date and time the database software image was created. | 
| title | string | X | √ | Title of the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| tags | json | X | √ | A map of tags for the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| image_type | string | X | √ | The type of software image. It can be grid or database. | 
| lifecycle_state | string | X | √ | The current state of the database software image. | 
| image_shape_family | string | X | √ | The shape that the image is meant for. | 
| lifecycle_details | string | X | √ | Detailed message for the lifecycle state. | 
| is_upgrade_supported | bool | X | √ | True if this database software image is supported for upgrade. | 
| database_software_image_one_off_patches | json | X | √ | List of one-off patches for database homes. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| database_version | string | X | √ | The database version with which the database software image is to be built. | 
| ls_inventory | string | X | √ | The output from lsinventory which will get passed as a string. | 
| patch_set | string | X | √ | The PSU or PBP or release updates. | 
| database_software_image_included_patches | json | X | √ | List of one-off patches for database homes. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 


