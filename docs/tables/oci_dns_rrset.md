# Table: oci_dns_rrset

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| domain | string | X | √ | The fully qualified domain name where the record can be located. | 
| record_hash | string | X | √ | A unique identifier for the record within its zone. | 
| is_protected | bool | X | √ | A Boolean flag indicating whether or not parts of the record are unable to be explicitly managed. | 
| rrset_version | string | X | √ | The latest version of the record's zone in which its RRSet differs from the preceding version. | 
| rtype | string | X | √ | The type of DNS record, such as A or CNAME. | 
| title | string | X | √ | Title of the resource. | 
| rdata | string | X | √ | The record's data. | 
| ttl | string | X | √ | The Time To Live for the record, in seconds. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 


