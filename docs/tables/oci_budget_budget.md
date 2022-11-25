# Table: oci_budget_budget

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ | The description of the budget. | 
| amount | float | X | √ | The amount of the budget expressed in the currency of the customer's rate card. | 
| reset_period | string | X | √ | The reset period for the budget. | 
| targets | json | X | √ | The list of targets on which the budget is applied. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 
| lifecycle_state | string | X | √ | The current state of the budget. | 
| id | string | X | √ | The OCID of the budget. | 
| time_created | timestamp | X | √ | Time that budget was created. | 
| alert_rule_count | int | X | √ | Total number of alert rules in the budget. | 
| forecasted_spend | float | X | √ | The forecasted spend in currency by the end of the current budget cycle. | 
| time_updated | timestamp | X | √ | Time that budget was updated. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| display_name | string | X | √ | The display name of the budget. | 
| target_compartment_id | string | X | √ | This is DEPRECATED. For backwards compatability, the property will be populated when targetType is COMPARTMENT AND targets contains EXACT ONE target compartment ocid. For all other scenarios, this property will be left empty. | 
| budget_processing_period_start_offset | int | X | √ | The number of days offset from the first day of the month, at which the budget processing period starts. | 
| target_type | string | X | √ | The type of target on which the budget is applied. | 
| time_spend_computed | timestamp | X | √ | The time that the budget spend was last computed. | 
| version | int | X | √ | Version of the budget. Starts from 1 and increments by 1. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| title | string | X | √ | Title of the resource. | 
| actual_spend | float | X | √ | The actual spend in currency for the current budget cycle. | 


