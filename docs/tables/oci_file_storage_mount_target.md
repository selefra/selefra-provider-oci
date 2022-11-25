# Table: oci_file_storage_mount_target

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| display_name | string | X | √ | A user-friendly name of the Mount Target. | 
| availability_domain | string | X | √ | The availability domain the mount target is in. | 
| export_set_id | string | X | √ | The OCID of the associated export set. | 
| nsg_ids | json | X | √ | A list of Network Security Group OCIDs associated with this mount target. | 
| private_ip_ids | json | X | √ | The OCIDs of the private IP addresses associated with this mount target. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| id | string | X | √ | The OCID of the Mount Target. | 
| lifecycle_state | string | X | √ | The current state of the Mount Target. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| subnet_id | string | X | √ | The OCIDs of the subnet the mount target is in. | 
| time_created | timestamp | X | √ | The date and time the Mount Target was created. | 
| lifecycle_details | string | X | √ | Additional information about the current 'lifecycleState'. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 


