# Table: oci_file_storage_snapshot

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| provenance_id | string | X | √ | An OCID identifying the parent from which this snapshot was cloned. | 
| is_clone_source | bool | X | √ | Specifies whether the snapshot has been cloned. | 
| lifecycle_details | string | X | √ | Additional information about the current 'lifecycleState'. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| title | string | X | √ | Title of the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| id | string | X | √ | The OCID of the snapshot. | 
| file_system_id | string | X | √ | The OCID of the file system from which the snapshot was created. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| time_created | timestamp | X | √ | The date and time the snapshot was created. | 
| tags | json | X | √ | A map of tags for the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| name | string | X | √ | Name of the snapshot. | 
| lifecycle_state | string | X | √ | The current state of the snapshot. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 


