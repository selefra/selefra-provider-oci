# Table: oci_events_rule

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| is_enabled | bool | X | √ | Indicates whether or not this rule is currently enabled. | 
| description | string | X | √ | A string that describes the details of the rule. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| compartment_id | string | X | √ | ColumnDescriptionCompartment | 
| lifecycle_state | string | X | √ | The current state of the rule. | 
| condition | json | X | √ | A filter that specifies the event that will trigger actions associated with this rule. | 
| title | string | X | √ | Title of the resource. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| display_name | string | X | √ | A string that describes the rule. | 
| time_created | timestamp | X | √ | The time this rule was created. | 
| actions | json | X | √ | An object that represents an action to trigger for events that match a rule. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| tags | json | X | √ | A map of tags for the resource. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| id | string | X | √ | The OCID of this rule. | 
| lifecycle_message | string | X | √ | A message generated by the Events service about the current state of this rule. | 


