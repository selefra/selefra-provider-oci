# Table: oci_budget_alert_rule

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| display_name | string | X | √ | The name of the alert rule. | 
| id | string | X | √ | The OCID of the alert rule. | 
| description | string | X | √ | The description of the alert rule. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| lifecycle_state | string | X | √ | The current state of the alert rule. | 
| time_created | timestamp | X | √ | Time that budget was created. | 
| message | string | X | √ | Custom message sent when alert is triggered. | 
| recipients | string | X | √ | Delimited list of email addresses to receive the alert when it triggers. | 
| time_updated | timestamp | X | √ | Time that budget was updated. | 
| type | string | X | √ | The type of alert. | 
| threshold | float | X | √ | The threshold for triggering the alert. If thresholdType is PERCENTAGE, the maximum value is 10000. | 
| threshold_type | string | X | √ | The type of threshold. | 
| version | int | X | √ | Version of the alert rule. Starts from 1 and increments by 1. | 
| title | string | X | √ | Title of the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| budget_id | string | X | √ | The OCID of the budget | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 


