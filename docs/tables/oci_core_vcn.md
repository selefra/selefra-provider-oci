# Table: oci_core_vcn

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| time_created | timestamp | X | √ | The date and time the VCN was created. | 
| cidr_block | cidr | X | √ | The first CIDR IP address from cidrBlocks. | 
| default_dhcp_options_id | string | X | √ | The OCID for the VCN's default set of DHCP options. | 
| default_security_list_id | string | X | √ | The OCID for the VCN's default security list. | 
| cidr_blocks | json | X | √ | The list of IPv4 CIDR blocks the VCN will use. | 
| ipv6_cidr_blocks | json | X | √ | For an IPv6-enabled VCN, this is the list of IPv6 CIDR blocks for the VCN's IP address space. The CIDRs are provided by Oracle and the sizes are always /56. | 
| lifecycle_state | string | X | √ | The VCN's current state. | 
| vcn_domain_name | string | X | √ | The VCN's domain name, which consists of the VCN's DNS label, and the oraclevcn.com domain. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| tags | json | X | √ | A map of tags for the resource. | 
| display_name | string | X | √ | A user-friendly name. Does not have to be unique, and it's changeable. | 
| id | string | X | √ | The VCN's Oracle ID (OCID). | 
| title | string | X | √ | Title of the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| default_route_table_id | string | X | √ | The OCID of the instance. | 
| dns_label | string | X | √ | A DNS label for the VCN, used in conjunction with the VNIC's hostname and subnet's DNS label to form a fully qualified domain name (FQDN) for each VNIC within this subnet. | 
| ipv6_cidr_block | cidr | X | √ | For an IPv6-enabled VCN, this is the IPv6 CIDR block for the VCN's private IP address space. | 
| ipv6_public_cidr_block | string | X | √ | For an IPv6-enabled VCN, this is the IPv6 CIDR block for the VCN's public IP address space. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 


