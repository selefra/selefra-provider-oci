# Table: oci_core_instance

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| shape_config_baseline_ocpu_utilization | string | X | √ | The baseline OCPU utilization for a subcore burstable VM instance. | 
| region | string | X | √ | The OCI region in which the resource is located. | 
| id | string | X | √ | The OCID of the instance. | 
| dedicated_vm_host_id | string | X | √ | The OCID of dedicated VM host. | 
| ipxe_script | string | X | √ | When a bare metal or virtual machine instance boots, the iPXE firmware that runs on the instance is configured to run an iPXE script to continue the boot process. | 
| shape_config_memory_in_gbs | float | X | √ | The total amount of memory available to the instance, in gigabytes. | 
| platform_config | json | X | √ | The platform configuration for the instance. | 
| source_details | json | X | √ | Contains the details of the source image for the instance. | 
| launch_mode | string | X | √ | Specifies the configuration mode for launching virtual machine (VM) instances. | 
| shape_config_max_vnic_attachments | int | X | √ | The maximum number of VNIC attachments for the instance. | 
| shape_config_ocpus | float | X | √ | The total number of OCPUs available to the instance. | 
| extended_metadata | json | X | √ | Additional metadata key/value pairs that user provided to instance. | 
| tags | json | X | √ | A map of tags for the resource. | 
| title | string | X | √ | Title of the resource. | 
| compartment_id | string | X | √ | The OCID of the compartment in Tenant in which the resource is located. | 
| time_created | timestamp | X | √ | The date and time the instance was created | 
| capacity_reservation_id | string | X | √ | The OCID of the compute capacity reservation this instance is launched under. | 
| shape_config_gpus | int | X | √ | The number of GPUs available to the instance. | 
| shape_config_networking_bandwidth_in_gbps | float | X | √ | The networking bandwidth available to the instance, in gigabits per second. | 
| availability_config | json | X | √ | Options for defining the availability of a VM instance after a maintenance event that impacts the underlying hardware. | 
| freeform_tags | json | X | √ | Free-form tags for resource. This tags can be applied by any user with permissions on the resource. | 
| shape_config_local_disks_total_size_in_gbs | float | X | √ | The aggregate size of all local disks, in gigabytes. | 
| time_maintenance_reboot_due | timestamp | X | √ | The date and time the instance is expected to be stopped/started. After that time if instance hasn't been rebooted, Oracle will reboot the instance within 24 hours of the due time. | 
| agent_config | json | X | √ | Options for the Oracle Cloud Agent software running on the instance. | 
| shape_config | json | X | √ | The shape configuration for an instance. The shape configuration determines the resources allocated to an instance. | 
| defined_tags | json | X | √ | Defined tags for resource. Defined tags are set up in your tenancy by an administrator. Only users granted permission to work with the defined tags can apply them to resources. | 
| system_tags | json | X | √ | Tags added to instances by the service. | 
| availability_domain | string | X | √ | The availability domain the instance is running in. | 
| shape | string | X | √ | The shape of the instance. The shape determines the number of CPUs and the amount of memory allocated to the instance. | 
| launch_options | json | X | √ | LaunchOptions Options for tuning the compatibility and performance of VM shapes. | 
| instance_options | json | X | √ | Optional mutable instance options. | 
| display_name | string | X | √ | A user-friendly name. Does not have to be unique, and it's changeable. | 
| lifecycle_state | string | X | √ | The current state of the instance. | 
| fault_domain | string | X | √ | The name of the fault domain the instance is running in. A fault domain is a grouping of hardware and infrastructure within an availability domain. | 
| shape_config_local_disks | int | X | √ | The number of local disks available to the instance. | 
| tenant_id | string | X | √ | The OCID of the Tenant in which the resource is located. | 
| metadata | json | X | √ | Custom metadata that you provided to instance. | 


