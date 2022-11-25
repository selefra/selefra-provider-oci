# Table: oci_database_autonomous_database

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| infrastructure_type | string | X | √ | The infrastructure type this resource belongs to. | 
| permission_level | string | X | √ | The Autonomous Database permission level. | 
| system_tags | json | X | √ | System tags for resource. System tags can be viewed by users, but can only be created by the system. | 
| is_auto_scaling_enabled | bool | X | √ | Indicates if auto scaling is enabled for the Autonomous Database CPU core count. | 
| refreshable_status | string | X | √ | The refresh status of the clone. REFRESHING indicates that the clone is currently being refreshed with data from the source Autonomous Database. | 
| time_of_last_switchover | timestamp | X | √ | The timestamp of the last switchover operation for the Autonomous Database. | 
| whitelisted_ips | json | X | √ | The client IP access control list (ACL). | 
| tags | json | X | √ | A map of tags for the resource. | 
| id | string | X | √ | The OCID of the Autonomous Database. | 
| is_dedicated | bool | X | √ | True if the database uses dedicated Exadata infrastructure. | 
| license_model | string | X | √ | The Oracle license model that applies to the Oracle Autonomous Database. | 
| time_of_last_refresh | timestamp | X | √ | The date and time when last refresh happened. | 
| apex_details | json | X | √ | Information about Oracle APEX Application Development. | 
| available_upgrade_versions | json | X | √ | List of Oracle Database versions available for a database upgrade. If there are no version upgrades available, this list is empty. | 
| cpu_core_count | int | X | √ | The number of OCPU cores to be made available to the database. | 
| is_access_control_enabled | bool | X | √ | Indicates if the database-level access control is enabled. | 
| open_mode | string | X | √ | The `DATABASE OPEN` mode. You can open the database in `READ_ONLY` or `READ_WRITE` mode. | 
| source_id | string | X | √ | The OCID of the source Autonomous Database that was cloned to create the current Autonomous Database. | 
| time_maintenance_begin | timestamp | X | √ | The date and time when maintenance will begin. | 
| data_storage_size_in_gbs | int | X | √ | The quantity of data in the database, in gigabytes. | 
| is_refreshable_clone | bool | X | √ | Indicates whether the Autonomous Database is a refreshable clone. | 
| backup_config | json | X | √ | Autonomous Database configuration details for storing manual backups in the Object Storage service. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| lifecycle_details | string | X | √ | Information about the current lifecycle state. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| time_created | timestamp | X | √ | The date and time the Autonomous Database was created. | 
| is_data_guard_enabled | bool | X | √ | Indicates whether the Autonomous Database has Data Guard enabled. | 
| operations_insights_status | string | X | √ | Status of Operations Insights for this Autonomous Database. | 
| standby_db | json | X | √ | Autonomous Data Guard standby database details. | 
| title | string | X | √ | Title of the resource. | 
| refreshable_mode | string | X | √ | The refresh mode of the clone. AUTOMATIC indicates that the clone is automatically being refreshed with data from the source Autonomous Database. | 
| time_deletion_of_free_autonomous_database | timestamp | X | √ | The date and time the Always Free database will be automatically deleted because of inactivity. If the database is in the STOPPED state and without activity until this time, it will be deleted. | 
| time_of_last_failover | timestamp | X | √ | The timestamp of the last failover operation. | 
| connection_strings | json | X | √ | The connection string used to connect to the Autonomous Database. The username for the Service Console is ADMIN. Use the password you entered when creating the Autonomous Database for the password value. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| connection_urls | json | X | √ | The URLs for accessing Oracle Application Express (APEX) and SQL Developer Web with a browser from a Compute instance within your VCN or that has a direct connection to your VCN. Note that these URLs are provided by the console only for databases on dedicated Exadata infrastructure. | 
| standby_whitelisted_ips | json | X | √ | The client IP access control list (ACL). This feature is available for autonomous databases on shared Exadata infrastructure and on Exadata Cloud@Customer. Only clients connecting from an IP address included in the ACL may access the Autonomous Database instance. For shared Exadata infrastructure, this is an array of CIDR (Classless Inter-Domain Routing) notations for a subnet or VCN OCID. | 
| db_name | string | X | √ | The database name. | 
| data_safe_status | string | X | √ | Status of the Data Safe registration for this Autonomous Database. | 
| db_workload | string | X | √ | The Autonomous Database workload type. | 
| role | string | X | √ | The role of the Autonomous Data Guard-enabled Autonomous Container Database. | 
| time_of_next_refresh | timestamp | X | √ | The date and time of next refresh. | 
| failed_data_recovery_in_seconds | int | X | √ | Indicates the number of seconds of data loss for a Data Guard failover. | 
| is_free_tier | bool | X | √ | Indicates if this is an Always Free resource. The default value is false. | 
| private_endpoint_ip | string | X | √ | The private endpoint Ip address for the resource. | 
| service_console_url | string | X | √ | The URL of the Service Console for the Autonomous Database. | 
| subnet_id | string | X | √ | The OCID of the subnet the resource is associated with. | 
| autonomous_container_database_id | string | X | √ | The Autonomous Container Database OCID. | 
| key_store_id | string | X | √ | The OCID of the key store. | 
| private_endpoint | string | X | √ | The private endpoint for the resource. | 
| display_name | string | X | √ | The user-friendly name for the Autonomous Database. The name does not have to be unique. | 
| are_primary_whitelisted_ips_used | bool | X | √ | This field will be null if the Autonomous Database is not Data Guard enabled or Access Control is disabled. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| private_endpoint_label | string | X | √ | The private endpoint label for the resource. | 
| time_of_last_refresh_point | timestamp | X | √ | The refresh point timestamp (UTC). The refresh point is the time to which the database was most recently refreshed. Data created after the refresh point is not included in the refresh. | 
| time_reclamation_of_free_autonomous_database | timestamp | X | √ | The date and time the Always Free database will be stopped because of inactivity. If this time is reached without any database activity, the database will automatically be put into the STOPPED state. | 
| nsg_ids | json | X | √ | A list of the OCIDs of the network security groups (NSGs) that this resource belongs to. | 
| lifecycle_state | string | X | √ | The current state of the Autonomous Database. | 
| data_storage_size_in_tbs | int | X | √ | The quantity of data in the database, in terabytes. | 
| db_version | string | X | √ | A valid Oracle Database version for Autonomous Database. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| is_preview | bool | X | √ | Indicates if the Autonomous Database version is a preview version. | 
| key_store_wallet_name | string | X | √ | The wallet name for Oracle Key Vault. | 
| time_maintenance_end | timestamp | X | √ | The date and time when maintenance will end. | 
| used_data_storage_size_in_tbs | int | X | √ | The amount of storage that has been used, in terabytes. | 


