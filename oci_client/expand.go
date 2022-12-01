package oci_client

import (
	"context"
	"fmt"
	oci_common "github.com/oracle/oci-go-sdk/v44/common"
	oci_common_auth "github.com/oracle/oci-go-sdk/v44/common/auth"

	"github.com/oracle/oci-go-sdk/v44/identity"
	"github.com/selefra/selefra-provider-oci/constants"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-utils/pkg/pointer"
	"net/http"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
)

func BuildRegionList() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		client := taskClient.(*OciClient)

		if client.OciConfig != nil && len(client.OciConfig.Regions) != 0 {
			regions := client.OciConfig.Regions

			if len(getInvalidRegions(regions)) > 0 {
				panic(constants.NnConnectionconfighaveinvalidregions + strings.Join(getInvalidRegions(regions), constants.Constants_1))
			}

			slice := make([]*schema.ClientTaskContext, 0)
			for _, region := range regions {
				slice = append(slice, &schema.ClientTaskContext{
					Task:   task.Clone(),
					Client: client.CopyWithRegion(region),
				})
			}
			return slice
		}

		return []*schema.ClientTaskContext{
			&schema.ClientTaskContext{
				Task:   task.Clone(),
				Client: client.CopyWithRegion(constants.Sasaopaulo),
			},
		}
	}
}

func BuildCompartmentList() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		client := taskClient.(*OciClient)

		compartments, err := listAllCompartments(ctx, clientMeta, taskClient, task)
		if err != nil {
			if strings.Contains(err.Error(), constants.Properconfigurationforregion) || strings.Contains(err.Error(), constants.OCIREGION) {
				panic(constants.NnregionsmustbesetintheconnectionconfigurationEdityourconnectionconfigurationfileandthenrestartselefra)
			}
			panic(err)
		}

		slice := make([]*schema.ClientTaskContext, 0)
		for _, compartment := range compartments {
			slice = append(slice, &schema.ClientTaskContext{
				Task:   task.Clone(),
				Client: client.CopyWithCompartment(*compartment.Id),
			})
		}
		return slice
	}
}

func getInvalidRegions(regions []string) []string {

	ociRegions := []string{
		constants.Apchiyoda,
		constants.Apchuncheon,
		constants.Aphyderabad,
		constants.Apmelbourne,
		constants.Apmumbai,
		constants.Aposaka,
		constants.Apseoul,
		constants.Apsydney,
		constants.Aptokyo,
		constants.Camontreal,
		constants.Catoronto,
		constants.Euamsterdam,
		constants.Eufrankfurt,
		constants.Euzurich,
		constants.Medubai,
		constants.Mejeddah,
		constants.Sasantiago,
		constants.Sasaopaulo,
		constants.Savinhedo,
		constants.Ukcardiff,
		constants.Ukgovcardiff,
		constants.Ukgovlondon,
		constants.Uklondon,
		constants.Usashburn,
		constants.Usgovashburn,
		constants.Usgovchicago,
		constants.Usgovphoenix,
		constants.Uslangley,
		constants.Usluke,
		constants.Usphoenix,
		constants.Ussanjose,
	}

	var invalidRegions []string

	for _, region := range regions {

		if !in(region, ociRegions) {
			invalidRegions = append(invalidRegions, region)
		}
	}

	return invalidRegions
}

func in(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	if index < len(str_array) && str_array[index] == target {
		return true
	}
	return false
}

func listAllCompartments(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) ([]identity.Compartment, error) {

	session, err := IdentityService(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	compartments := []identity.Compartment{
		{
			Id: &session.TenancyID,
		},
	}

	request := identity.ListCompartmentsRequest{
		CompartmentId:          &session.TenancyID,
		CompartmentIdInSubtree: pointer.ToBoolPointer(true),
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: GetDefaultRetryPolicy(taskClient),
		},
	}

	pagesLeft := true
	for pagesLeft {
		response, err := session.IdentityClient.ListCompartments(ctx, request)
		if err != nil {
			return nil, err
		}

		for _, compartment := range response.Items {
			if compartment.LifecycleState == "ACTIVE" || compartment.LifecycleState == "INACTIVE" {
				compartments = append(compartments, compartment)
			}
		}

		if response.OpcNextPage != nil {
			request.Page = response.OpcNextPage
		} else {
			pagesLeft = false
		}
	}

	return compartments, err
}

type zoneInfo struct {
	identity.AvailabilityDomain
	Region string
}

func listAllzones(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) ([]zoneInfo, error) {

	zonesList := []zoneInfo{}

	regions := taskClient.(*OciClient).Region

	session, err := IdentityServiceRegional(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	request := identity.ListAvailabilityDomainsRequest{
		CompartmentId: &session.TenancyID,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: GetDefaultRetryPolicy(taskClient),
		},
	}

	response, err := session.IdentityClient.ListAvailabilityDomains(ctx, request)
	if err != nil {
		return nil, err
	}

	for _, zones := range response.Items {
		zonesList = append(zonesList, zoneInfo{zones, regions})
	}
	return zonesList, nil
}

func BuildCompartementZonalList() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
		client := taskClient.(*OciClient)

		compartments, err := listAllCompartments(ctx, clientMeta, taskClient, task)
		if err != nil {
			if strings.Contains(err.Error(), constants.Properconfigurationforregion) || strings.Contains(err.Error(), constants.OCIREGION) {
				panic(constants.NnregionsmustbesetintheconnectionconfigurationEdityourconnectionconfigurationfileandthenrestartselefra)
			}
			panic(err)
		}

		zones, err := listAllzones(ctx, clientMeta, taskClient, task)
		if err != nil {
			if strings.Contains(err.Error(), constants.Properconfigurationforregion) || strings.Contains(err.Error(), constants.OCIREGION) {
				panic(constants.NnregionsmustbesetintheconnectionconfigurationEdityourconnectionconfigurationfileandthenrestartselefra)
			}
			panic(err)
		}

		matrix := make([]*schema.ClientTaskContext, 0)

		for _, zone := range zones {
			for _, compartment := range compartments {
				matrix = append(matrix, &schema.ClientTaskContext{
					Task:   task.Clone(),
					Client: client.CopyWithAll(zone.Region, *compartment.Id, *zone.Name),
				})
			}
		}

		return matrix
	}
}

func getRegionFromEnvVar() string {
	if region, ok := os.LookupEnv(constants.OCIREGION); ok {
		return region
	} else if region, ok = os.LookupEnv(constants.OCICLIREGION); ok {
		return region
	}

	return getEnvSettingWithBlankDefault(constants.Region)
}

func GetEnv(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (oci_common.ConfigurationProvider, error) {
	ociConfig := taskClient.(*OciClient).OciConfig
	region := taskClient.(*OciClient).Region

	if region == constants.Constants_2 {
		region = getRegionFromEnvVar()
	}

	authType := constants.ApiKey

	if ociConfig.Auth != nil && (*ociConfig.Auth != constants.ApiKey && *ociConfig.Auth != constants.Constants_3) {
		authType = *ociConfig.Auth
	}

	if authType == constants.SecurityToken {
		return getProviderForSecurityToken(region, ociConfig)
	}

	if authType == constants.InstancePrincipal {
		return getProviderForInstancePrincipal(region)
	}

	if authType == constants.ApiKey {
		return getProviderForAPIkey(region, ociConfig)
	}

	regionInfo := oci_common.NewRawConfigurationProvider(constants.Constants_4, constants.Constants_5, region, constants.Constants_6, constants.Constants_7, nil)
	provider, err := oci_common.ComposingConfigurationProvider([]oci_common.ConfigurationProvider{regionInfo, oci_common.DefaultConfigProvider()})
	if err != nil {
		return nil, err
	}

	return provider, nil
}

func getProviderForSecurityToken(region string, config *OciConfig) (oci_common.ConfigurationProvider, error) {
	regionInfo := oci_common.NewRawConfigurationProvider(constants.Constants_8, constants.Constants_9, region, constants.Constants_10, constants.Constants_11, nil)

	if config.Profile == nil {
		return nil, fmt.Errorf(constants.NnconfigfileprofilemustbesetintheconnectionconfigurationforSecurityTokenauthenticationEdityourconnectionconfigurationfileandthenrestartSelefra)
	}

	profileString := *config.Profile
	defaultPath := path.Join(getHomeFolder(), constants.Oci, constants.Config)
	if err := checkProfile(profileString, defaultPath); err != nil {
		return nil, err
	}

	securityTokenBasedAuthConfigProvider := oci_common.CustomProfileConfigProvider(defaultPath, profileString)

	keyId, err := securityTokenBasedAuthConfigProvider.KeyID()
	if err != nil || !strings.HasPrefix(keyId, constants.ST) {
		return nil, fmt.Errorf(constants.Securitytokenisinvalid)
	}

	return oci_common.ComposingConfigurationProvider([]oci_common.ConfigurationProvider{regionInfo, securityTokenBasedAuthConfigProvider})
}

func getProviderForInstancePrincipal(region string) (oci_common.ConfigurationProvider, error) {

	instancePrincipalAuthClientModifier := func(client oci_common.HTTPRequestDispatcher) (oci_common.HTTPRequestDispatcher, error) {
		if acceptLocalCerts := getEnvSettingWithBlankDefault(constants.Acceptlocalcerts); acceptLocalCerts != constants.Constants_12 {
			if bool, err := strconv.ParseBool(acceptLocalCerts); err == nil {
				modifiedClient := buildHttpClient()
				modifiedClient.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify = bool
				return modifiedClient, nil
			}
		}
		return client, nil
	}

	cfg, err := oci_common_auth.InstancePrincipalConfigurationForRegionWithCustomClient(oci_common.StringToRegion(region), instancePrincipalAuthClientModifier)
	if err != nil {
		return nil, err
	}

	return oci_common.ComposingConfigurationProvider([]oci_common.ConfigurationProvider{cfg})
}

func getProviderForAPIkey(region string, config *OciConfig) (oci_common.ConfigurationProvider, error) {

	regionInfo := oci_common.NewRawConfigurationProvider(constants.Constants_13, constants.Constants_14, region, constants.Constants_15, constants.Constants_16, nil)

	if config.Profile != nil {
		configPath := constants.Constants_17
		if config.ConfigPath != nil {
			configPath = *config.ConfigPath
		}

		configProvider := oci_common.CustomProfileConfigProvider(configPath, *config.Profile)
		configProviderEnvironmentVariables := oci_common.ConfigurationProviderEnvironmentVariables(constants.OCI, constants.Constants_18)

		return oci_common.ComposingConfigurationProvider([]oci_common.ConfigurationProvider{regionInfo, configProvider, configProviderEnvironmentVariables})
	}

	if config.UserOCID != nil {
		pemFilePassword := constants.Constants_19
		pemFileContent := constants.Constants_20
		if config.PrivateKey != nil {
			pemFileContent = *config.PrivateKey
		}
		if config.PrivateKeyPath != nil {
			resolvedPath := expandPath(*config.PrivateKeyPath)
			pemFileData, err := os.ReadFile(resolvedPath)
			if err != nil {
				return nil, fmt.Errorf(constants.CannotreadprivatekeyfromsErrorq, *config.PrivateKeyPath, err)
			}
			pemFileContent = string(pemFileData)
		}

		if config.PrivateKeyPassword != nil {
			pemFilePassword = *config.PrivateKeyPassword
		}

		configProvider := oci_common.NewRawConfigurationProvider(*config.TenancyOCID, *config.UserOCID, region, *config.Fingerprint, pemFileContent, &pemFilePassword)
		configProviderEnvironmentVariables := oci_common.ConfigurationProviderEnvironmentVariables(constants.OCI, pemFilePassword)

		return oci_common.ComposingConfigurationProvider([]oci_common.ConfigurationProvider{regionInfo, configProvider, configProviderEnvironmentVariables})
	}

	var providers []oci_common.ConfigurationProvider
	providers = append(providers, regionInfo, oci_common.DefaultConfigProvider())
	cliProvider, _ := getProviderFromCLIEnvironmentVariables()
	if cliProvider != nil {
		providers = append(providers, cliProvider)
	}

	return oci_common.ComposingConfigurationProvider(providers)
}

