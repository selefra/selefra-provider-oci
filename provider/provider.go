package provider

import (
	"context"
	"github.com/selefra/selefra-provider-oci/constants"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-utils/pkg/pointer"
	"github.com/spf13/viper"
)

var Version = constants.V

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      "oci",
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				regions := config.GetStringSlice(constants.Providersregions)

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
				return `regions
	- "<xxxxx>"
	- "<xxxxx>"`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
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

