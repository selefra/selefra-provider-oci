# Table: oci_mysql_heat_wave_cluster

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| time_created | timestamp | X | √ | The date and time the HeatWave cluster was created. | 
| cluster_size | int | X | √ | The number of analytics-processing compute instances, of the specified shape, in the HeatWave cluster. | 
| lifecycle_details | string | X | √ | Additional information about the current lifecycleState. | 
| lifecycle_state | string | X | √ | The current state of the HeatWave cluster. | 
| shape_name | string | X | √ | The shape determines resources to allocate to the HeatWave nodes - CPU cores, memory. | 
| time_updated | timestamp | X | √ | The time the HeatWave cluster was last updated. | 
| cluster_nodes | json | X | √ | A HeatWave node is a compute host that is part of a HeatWave cluster. | 
| db_system_id | string | X | √ | The OCID of the parent DB System this HeatWave cluster is attached to. | 


