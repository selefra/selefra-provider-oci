# Table: oci_containerengine_cluster

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| image_policy_config_enabled | bool | X | √ | Whether the image verification policy is enabled. Defaults to false. If set to true, the images will be verified against the policy at runtime. | 
| endpoints | json | X | √ | Endpoints served up by the cluster masters. | 
| endpoint_config | json | X | √ | The network configuration for access to the Cluster control plane. | 
| metadata | json | X | √ | Metadata about the cluster. | 
| title | string | X | √ | Title of the resource. | 
| name | string | X | √ | A user-friendly name. It does not have to be unique, and it is changeable. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| lifecycle_state | string | X | √ | The state of the cluster masters. | 
| vcn_id | string | X | √ | The OCID of the virtual cloud network (VCN) in which the cluster exists. | 
| options | json | X | √ | Optional attributes for the cluster. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| id | string | X | √ | The OCID of the cluster. | 
| kms_key_id | string | X | √ | The OCID of the KMS key to be used as the master encryption key for Kubernetes secret encryption. | 
| kubernetes_version | string | X | √ | The version of Kubernetes running on the cluster masters. | 
| lifecycle_details | string | X | √ | Additional information about the current 'lifecycleState'. | 
| available_kubernetes_upgrades | json | X | √ | Available Kubernetes versions to which the clusters masters may be upgraded. | 


