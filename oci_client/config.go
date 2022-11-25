package oci_client

type OciConfig struct {
	Auth                  *string  `cty:"auth"`
	ConfigPath            *string  `cty:"config_path"`
	Fingerprint           *string  `cty:"fingerprint"`
	PrivateKey            *string  `cty:"private_key"`
	PrivateKeyPassword    *string  `cty:"private_key_password"`
	PrivateKeyPath        *string  `cty:"private_key_path"`
	Profile               *string  `cty:"config_file_profile"`
	Regions               []string `cty:"regions"`
	TenancyOCID           *string  `cty:"tenancy_ocid"`
	UserOCID              *string  `cty:"user_ocid"`
	MaxErrorRetryAttempts *int     `cty:"max_error_retry_attempts"`
	MinErrorRetryDelay    *int     `cty:"min_error_retry_delay"`
}
