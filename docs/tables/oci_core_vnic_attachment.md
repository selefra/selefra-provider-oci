# Table: oci_core_vnic_attachment

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| title | string | X | √ | Title of the resource. | 
| id | string | X | √ | The OCID of the VNIC attachment. | 
| skip_source_dest_check | bool | X | √ | Whether the source/destination check is disabled on the VNIC. Defaults to `false`, which means the check is performed. | 
| subnet_id | string | X | √ | The OCID of the subnet to create the VNIC in. | 
| nsg_ids | json | X | √ | A list of the OCIDs of the network security groups that the VNIC belongs to. | 
| vnic_name | string | X | √ | A user-friendly name for the VNIC. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| tags | json | X | √ | A map of tags for the resource. | 
| instance_id | string | X | √ | The OCID of the instance. | 
| is_primary | bool | X | √ | Whether the VNIC is the primary VNIC (the VNIC that is automatically created and attached during instance launch). | 
| nic_index | int | X | √ | The physical network interface card (NIC) the VNIC uses. | 
| vnic_id | string | X | √ | The OCID of the VNIC. | 
| hostname_label | string | X | √ | The hostname for the VNIC's primary private IP. | 
| vlan_id | string | X | √ | The OCID of the VLAN to create the VNIC in. | 
| display_name | string | X | √ | A user-friendly name for the VNIC attachment. | 
| lifecycle_state | string | X | √ | The current state of the VNIC attachment. Possible values include: 'ATTACHING', 'ATTACHED', 'DETACHING', 'DETACHED'. | 
| private_ip | string | X | √ | The private IP address of the primary `privateIp` object on the VNIC. | 
| public_ip | string | X | √ | The public IP address of the VNIC, if one is assigned. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| availability_domain | string | X | √ | The availability domain of the instance. | 
| time_created | timestamp | X | √ | The date and time the VNIC attachment was created. | 
| mac_address | string | X | √ | The MAC address of the VNIC. | 
| vlan_tag | int | X | √ | The OCID of the VNIC. | 


