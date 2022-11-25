# Table: oci_core_subnet

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| tags | json | X | √ | A map of tags for the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| id | string | X | √ | The subnet's Oracle ID (OCID). | 
| route_table_id | string | X | √ | The OCID of the route table that the subnet uses. | 
| ipv6_virtual_router_ip | ip | X | √ | For an IPv6-enabled subnet, this is the IPv6 address of the virtual router. | 
| prohibit_public_ip_on_vnic | bool | X | √ | Indicates whether VNICs within this subnet can have public IP addresses. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| dns_label | string | X | √ | A DNS label for the subnet, used in conjunction with the VNIC's hostname and VCN's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet. | 
| security_list_ids | json | X | √ | The OCIDs of the security list or lists that the subnet uses. | 
| title | string | X | √ | Title of the resource. | 
| vcn_id | string | X | √ | The OCID of the VCN the subnet is in. | 
| lifecycle_state | string | X | √ | The subnet's current state. | 
| availability_domain | string | X | √ | The subnet's availability domain. | 
| cidr_block | cidr | X | √ | The subnet's CIDR block. | 
| dhcp_options_id | string | X | √ | The OCID of the set of DHCP options that the subnet uses. | 
| time_created | timestamp | X | √ | The date and time the subnet was created. | 
| ipv6_cidr_block | cidr | X | √ | For an IPv6-enabled subnet, this is the IPv6 CIDR block for the subnet's private IP address space. | 
| ipv6_public_cidr_block | cidr | X | √ | For an IPv6-enabled subnet, this is the IPv6 CIDR block for the subnet's public IP address space. | 
| virtual_router_ip | ip | X | √ | The IP address of the virtual router. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| display_name | string | X | √ | A user-friendly name. Does not have to be unique, and it's changeable. | 
| subnet_domain_name | string | X | √ | The subnet's domain name, which consists of the subnet's DNS label, the VCN's DNS label, and the `oraclevcn.com` domain. | 
| virtual_router_mac | string | X | √ | The MAC address of the virtual router. | 


