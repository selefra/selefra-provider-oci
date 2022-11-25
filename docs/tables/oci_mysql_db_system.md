# Table: oci_mysql_db_system

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| lifecycle_state | string | X | √ | The current state of the DB System. | 
| availability_domain | string | X | √ | The Availability Domain where the primary DB System should be located. | 
| data_storage_size_in_gbs | int | X | √ | Initial size of the data volume in GiBs that will be created and attached. | 
| ip_address | ip | X | √ | The IP address the DB System is configured to listen on. | 
| shape_name | string | X | √ | The shape of the primary instances of the DB System. | 
| source | json | X | √ | DbSystemSource Parameters detailing how to provision the initial data of the DB System. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| time_created | timestamp | X | √ | The date and time the DB System was created. | 
| description | string | X | √ | User-provided data about the DB System. | 
| backup_policy | json | X | √ | BackupPolicy The Backup policy for the DB System. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| display_name | string | X | √ | The user-friendly name for the DB System. It does not have to be unique. | 
| subnet_id | string | X | √ | The OCID of the subnet the DB System is associated with. | 
| hostname_label | string | X | √ | The hostname for the primary endpoint of the DB System. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| id | string | X | √ | The OCID of the DB System. | 
| fault_domain | string | X | √ | The name of the fault domain the DB System is located in. | 
| analytics_cluster | json | X | √ | A summary of an Analytics Cluster. | 
| endpoints | json | X | √ | The network endpoints available for this DB System. | 
| mysql_version | string | X | √ | Name of the MySQL Version in use for the DB System. | 
| port_x | int | X | √ | The network port on which X Plugin listens for TCP/IP connections. This is the X Plugin equivalent of port. | 
| is_analytics_cluster_attached | bool | X | √ | If the DB System has an Analytics Cluster attached. | 
| is_heat_wave_cluster_attached | bool | X | √ | Whether the DB System has a HeatWave cluster attached. | 
| lifecycle_details | string | X | √ | Additional information about the current lifecycleState. | 
| channels | json | X | √ | A list with a summary of all the Channels attached to the DB System. | 
| configuration_id | string | X | √ | The OCID of the Configuration to be used for Instances in this DB System. | 
| time_updated | timestamp | X | √ | The time the DB System was last updated. | 
| maintenance | json | X | √ | The Maintenance Policy for the DB System. | 
| port | int | X | √ | The port for primary endpoint of the DB System to listen on. | 
| tags | json | X | √ | A map of tags for the resource. | 


