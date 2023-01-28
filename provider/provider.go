package provider

import (
	"context"
	"github.com/selefra/selefra-provider-oci/constants"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-utils/pkg/pointer"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var Version = constants.V

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      "oci",
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				regions := config.GetStringSlice("regions")

				if len(regions) == 0 {
					regionData := os.Getenv("OCI_REGIONS")

					var regionList []string

					if regionData != "" {
						regionList = strings.Split(regionData, ",")
					}

					regions = regionList
				}

				if len(regions) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg("analysis config err: no configuration")
				}

				var ociConfig *oci_client.OciConfig

				ociConfig = &oci_client.OciConfig{
					Auth:    pointer.ToStringPointer(constants.ApiKey),
					Regions: regions,
				}

				return []any{
					&oci_client.OciClient{
						Region:    regions[0],
						OciConfig: ociConfig,
					},
				}, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# regions:
# - ap-seoul-1
# - us-phoenix-1`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				regions := config.GetStringSlice("regions")

				if len(regions) == 0 {
					regionData := os.Getenv("OCI_REGIONS")

					var regionList []string

					if regionData != "" {
						regionList = strings.Split(regionData, ",")
					}

					regions = regionList
				}

				if len(regions) == 0 {
					return schema.NewDiagnostics().AddErrorMsg("analysis config err: no configuration")
				}

				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{},
			DataSourcePullResultAutoExpand:       true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{
			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
