# Table: oci_database_db_system

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| scan_ip_ids | json | X | √ | The OCID of the single client access name (SCAN) IP addresses associated with the DB system. | 
| id | string | X | √ | The OCID of the DB system. | 
| next_maintenance_run_id | string | X | √ | The OCID of the next maintenance run. | 
| version | string | X | √ | The oracle database version of the DB system. | 
| zone_id | string | X | √ | The OCID of the zone the DB system is associated with. | 
| sparse_diskgroup | bool | X | √ | True, If sparse diskgroup is configured for exadata DB system. | 
| maintenance_window | json | X | √ | The maintenance window of the DB system. | 
| ssh_public_keys | json | X | √ | The public key portion of one or more key pairs used for SSH access to the DB system. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| domain | string | X | √ | The domain name for the DB system. | 
| license_model | string | X | √ | The oracle license model that applies to all the databases on the DB system. | 
| scan_dns_name | string | X | √ | The FQDN of the DNS record for the SCAN IP addresses that are associated with the DB system. | 
| source_db_system_id | string | X | √ | The OCID of the DB system from where the DB system is created. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| reco_storage_size_in_gb | int | X | √ | The RECO/REDO storage size, in gigabytes, that is currently allocated to the DB system. | 
| iorm_config_cache | json | X | √ | The IORM configuration of the DB system. | 
| nsg_ids | json | X | √ | A list of the OCIDs of the network security groups (NSGs) that this resource belongs to. | 
| lifecycle_state | string | X | √ | The current state of the DB system. | 
| data_storage_percentage | int | X | √ | The percentage assigned to data storage. | 
| host_name | string | X | √ | The hostname for the DB system. | 
| listener_port | int | X | √ | The port number configured for the listener on the DB system. | 
| subnet_id | string | X | √ | The OCID of the subnet the DB system is associated with. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| time_created | timestamp | X | √ | The date and time the DB system was created. | 
| data_storage_size_in_gbs | int | X | √ | The data storage size, in gigabytes, that is currently available to the DB system. | 
| db_system_options_storage_management | string | X | √ | The storage option used in DB system. | 
| node_count | int | X | √ | The number of nodes in the DB system. | 
| lifecycle_details | string | X | √ | Additional information about the current lifecycle state. | 
| fault_domains | json | X | √ | List of the fault domains in which this DB system is provisioned. | 
| vip_ids | json | X | √ | A list of the OCIDs of the virtual IP (VIP) addresses associated with the DB system. | 
| availability_domain | string | X | √ | The name of the availability domain that the DB system is located in. | 
| backup_subnet_id | string | X | √ | The OCID of the backup network subnet the DB system is associated with. | 
| disk_redundancy | string | X | √ | The type of redundancy configured for the DB system. | 
| last_patch_history_entry_id | string | X | √ | The OCID of the last patch history. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| display_name | string | X | √ | The user-friendly name for the DB system. The name does not have to be unique. | 
| cpu_core_count | int | X | √ | The number of CPU cores enabled on the DB system. | 
| point_in_time_data_disk_clone_timestamp | timestamp | X | √ | The point in time for a cloned database system when the data disks were cloned from the source database system. | 
| backup_network_nsg_ids | json | X | √ | A list of the OCIDs of the network security groups (NSGs) that the backup network of this DB system belongs to. | 
| time_zone | string | X | √ | The time zone of the DB system. | 
| cluster_name | string | X | √ | The cluster name for exadata and 2-node RAC virtual machine DB systems. | 
| kms_key_id | string | X | √ | The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations. | 
| last_maintenance_run_id | string | X | √ | The OCID of the last maintenance run. | 
| scan_dns_record_id | string | X | √ | The OCID of the DNS record for the SCAN IP addresses that are associated with the DB system. | 
| database_edition | string | X | √ | The oracle database edition that applies to all the databases on the DB system. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 


