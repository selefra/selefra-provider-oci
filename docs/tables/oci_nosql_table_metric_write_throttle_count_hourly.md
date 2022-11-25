# Table: oci_nosql_table_metric_write_throttle_count_hourly

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| maximum | float | X | √ | The maximum metric value for the data point. | 
| sum | float | X | √ | The sum of the metric values for the data point. | 
| unit | string | X | √ | The standard unit for the data point. | 
| timestamp | timestamp | X | √ | The time stamp used for the data point. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| compartment_id | string | X | √ | The ID of the compartment. | 
| average | float | X | √ | The average of the metric values that correspond to the data point. | 
| metric_name | string | X | √ | The name of the metric. | 
| namespace | string | X | √ | The metric namespace. | 
| minimum | float | X | √ | The minimum metric value for the data point. | 
| sample_count | float | X | √ | The number of metric values that contributed to the aggregate value of this data point. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| name | string | X | √ | The name of the NoSQL table. | 


