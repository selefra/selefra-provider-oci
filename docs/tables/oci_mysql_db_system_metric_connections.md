# Table: oci_mysql_db_system_metric_connections

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| metric_name | string | X | √ | The name of the metric. | 
| namespace | string | X | √ | The metric namespace. | 
| minimum | float | X | √ | The minimum metric value for the data point. | 
| unit | string | X | √ | The standard unit for the data point. | 
| timestamp | timestamp | X | √ | The time stamp used for the data point. | 
| compartment_id | string | X | √ | The ID of the compartment. | 
| id | string | X | √ | The OCID of the DB System. | 
| average | float | X | √ | The average of the metric values that correspond to the data point. | 
| maximum | float | X | √ | The maximum metric value for the data point. | 
| sample_count | float | X | √ | The number of metric values that contributed to the aggregate value of this data point. | 
| sum | float | X | √ | The sum of the metric values for the data point. | 


