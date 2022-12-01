package oci_client

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/selefra/selefra-provider-oci/constants"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-utils/pkg/pointer"
	"math"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/user"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/oracle/oci-go-sdk/v44/analytics"
	"github.com/oracle/oci-go-sdk/v44/apigateway"
	"github.com/oracle/oci-go-sdk/v44/audit"
	"github.com/oracle/oci-go-sdk/v44/autoscaling"
	"github.com/oracle/oci-go-sdk/v44/budget"
	"github.com/oracle/oci-go-sdk/v44/cloudguard"
	oci_common "github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/containerengine"
	"github.com/oracle/oci-go-sdk/v44/core"
	"github.com/oracle/oci-go-sdk/v44/database"
	"github.com/oracle/oci-go-sdk/v44/dns"
	"github.com/oracle/oci-go-sdk/v44/events"
	"github.com/oracle/oci-go-sdk/v44/filestorage"
	"github.com/oracle/oci-go-sdk/v44/functions"
	"github.com/oracle/oci-go-sdk/v44/identity"
	"github.com/oracle/oci-go-sdk/v44/keymanagement"
	"github.com/oracle/oci-go-sdk/v44/loadbalancer"
	"github.com/oracle/oci-go-sdk/v44/logging"
	"github.com/oracle/oci-go-sdk/v44/monitoring"
	"github.com/oracle/oci-go-sdk/v44/mysql"
	"github.com/oracle/oci-go-sdk/v44/networkloadbalancer"
	"github.com/oracle/oci-go-sdk/v44/nosql"
	"github.com/oracle/oci-go-sdk/v44/objectstorage"
	"github.com/oracle/oci-go-sdk/v44/ons"
	"github.com/oracle/oci-go-sdk/v44/resourcemanager"
	"github.com/oracle/oci-go-sdk/v44/resourcesearch"
	"github.com/oracle/oci-go-sdk/v44/streaming"
	"github.com/oracle/oci-go-sdk/v44/vault"
)

type session struct {
	TenancyID                      string
	AnalyticsClient                analytics.AnalyticsClient
	ApiGatewayClient               apigateway.ApiGatewayClient
	AuditClient                    audit.AuditClient
	AutoScalingClient              autoscaling.AutoScalingClient
	BlockstorageClient             core.BlockstorageClient
	BudgetClient                   budget.BudgetClient
	CloudGuardClient               cloudguard.CloudGuardClient
	ComputeClient                  core.ComputeClient
	ContainerEngineClient          containerengine.ContainerEngineClient
	DatabaseClient                 database.DatabaseClient
	DnsClient                      dns.DnsClient
	EventsClient                   events.EventsClient
	FileStorageClient              filestorage.FileStorageClient
	FunctionsManagementClient      functions.FunctionsManagementClient
	IdentityClient                 identity.IdentityClient
	KmsManagementClient            keymanagement.KmsManagementClient
	KmsVaultClient                 keymanagement.KmsVaultClient
	LoggingManagementClient        logging.LoggingManagementClient
	LoadBalancerClient             loadbalancer.LoadBalancerClient
	MonitoringClient               monitoring.MonitoringClient
	MySQLConfigurationClient       mysql.MysqlaasClient
	MySQLChannelClient             mysql.ChannelsClient
	MySQLBackupClient              mysql.DbBackupsClient
	MySQLDBSystemClient            mysql.DbSystemClient
	NetworkLoadBalancerClient      networkloadbalancer.NetworkLoadBalancerClient
	NoSQLClient                    nosql.NosqlClient
	NotificationControlPlaneClient ons.NotificationControlPlaneClient
	NotificationDataPlaneClient    ons.NotificationDataPlaneClient
	ObjectStorageClient            objectstorage.ObjectStorageClient
	ResourceSearchClient           resourcesearch.ResourceSearchClient
	ResourceManagerClient          resourcemanager.ResourceManagerClient
	StreamAdminClient              streaming.StreamAdminClient
	VaultClient                    vault.VaultsClient
	VirtualNetworkClient           core.VirtualNetworkClient
}

func getProvider(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (oci_common.ConfigurationProvider, error) {

	provider, err := GetEnv(ctx, clientMeta, taskClient, task)

	if err != nil {
		return nil, err
	}

	return provider, nil
}

func ApiGatewayService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := apigateway.NewApiGatewayClientWithConfigurationProvider(provider)

	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()

	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:        tenantId,
		ApiGatewayClient: client,
	}

	return sess, nil
}

func AuditService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := audit.NewAuditClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:   tenantId,
		AuditClient: client,
	}

	return sess, nil
}

func AutoScalingService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := autoscaling.NewAutoScalingClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:         tenantId,
		AutoScalingClient: client,
	}

	return sess, nil
}

func IdentityService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	if err != nil {
		return nil, err
	}

	client, err := identity.NewIdentityClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:      tenantId,
		IdentityClient: client,
	}

	return sess, nil
}

func IdentityServiceRegional(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := identity.NewIdentityClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:      tenantId,
		IdentityClient: client,
	}

	return sess, nil
}

func LoggingManagementService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := logging.NewLoggingManagementClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:               tenantId,
		LoggingManagementClient: client,
	}

	return sess, nil
}

func CoreBlockStorageService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := core.NewBlockstorageClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:          tenantId,
		BlockstorageClient: client,
	}

	return sess, nil
}

func ContainerEngineService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := containerengine.NewContainerEngineClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:             tenantId,
		ContainerEngineClient: client,
	}

	return sess, nil
}

func EventsService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := events.NewEventsClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:    tenantId,
		EventsClient: client,
	}

	return sess, nil
}

func FileStorageService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := filestorage.NewFileStorageClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantID, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:         tenantID,
		FileStorageClient: client,
	}
	return sess, nil
}

func FunctionsManagementService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := functions.NewFunctionsManagementClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantID, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:                 tenantID,
		FunctionsManagementClient: client,
	}

	return sess, nil
}

func KmsManagementService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {

	sess := &session{}
	return sess, nil
}

func KmsVaultService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := keymanagement.NewKmsVaultClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:      tenantId,
		KmsVaultClient: client,
	}

	return sess, nil
}

func LoadBalancerService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:          tenantId,
		LoadBalancerClient: client,
	}

	return sess, nil
}

func ObjectStorageService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:           tenantId,
		ObjectStorageClient: client,
	}

	return sess, nil
}

func OnsNotificationControlPlaneService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := ons.NewNotificationControlPlaneClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:                      tenantId,
		NotificationControlPlaneClient: client,
	}

	return sess, nil
}

func OnsNotificationDataPlaneService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := ons.NewNotificationDataPlaneClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:                   tenantId,
		NotificationDataPlaneClient: client,
	}
	return sess, nil
}

func CoreComputeService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := core.NewComputeClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:     tenantId,
		ComputeClient: client,
	}

	return sess, nil
}

func CoreVirtualNetworkService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := core.NewVirtualNetworkClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantID, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:            tenantID,
		VirtualNetworkClient: client,
	}

	return sess, nil
}

func CloudGuardService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := cloudguard.NewCloudGuardClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantID, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:        tenantID,
		CloudGuardClient: client,
	}

	return sess, nil
}

func DnsService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := dns.NewDnsClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantID, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID: tenantID,
		DnsClient: client,
	}
	return sess, nil
}

func DatabaseService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := database.NewDatabaseClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantID, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:      tenantID,
		DatabaseClient: client,
	}

	return sess, nil
}

func BudgetService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := budget.NewBudgetClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantID, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:    tenantID,
		BudgetClient: client,
	}

	return sess, nil
}

func MonitoringService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := monitoring.NewMonitoringClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantID, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:        tenantID,
		MonitoringClient: client,
	}
	return sess, nil
}

func MySQLChannelService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := mysql.NewChannelsClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantID, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:          tenantID,
		MySQLChannelClient: client,
	}

	return sess, nil
}

func MySQLDBSystemService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := mysql.NewDbSystemClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantID, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:           tenantID,
		MySQLDBSystemClient: client,
	}

	return sess, nil
}

func NoSQLDatabaseService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := nosql.NewNosqlClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantID, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:   tenantID,
		NoSQLClient: client,
	}

	return sess, nil
}

func MySQLBackupService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := mysql.NewDbBackupsClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantID, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:         tenantID,
		MySQLBackupClient: client,
	}

	return sess, nil
}

func MySQLConfigurationService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)
	client, err := mysql.NewMysqlaasClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantID, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:                tenantID,
		MySQLConfigurationClient: client,
	}

	return sess, nil
}

func NetworkLoadBalancerService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := networkloadbalancer.NewNetworkLoadBalancerClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantID, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:                 tenantID,
		NetworkLoadBalancerClient: client,
	}

	return sess, nil
}

func ResourceSearchService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := resourcesearch.NewResourceSearchClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:            tenantId,
		ResourceSearchClient: client,
	}

	return sess, nil
}

func ResourceManagerService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := resourcemanager.NewResourceManagerClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:             tenantId,
		ResourceManagerClient: client,
	}

	return sess, nil
}

func StreamAdminService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := streaming.NewStreamAdminClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:         tenantId,
		StreamAdminClient: client,
	}

	return sess, nil
}

func VaultService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := vault.NewVaultsClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:   tenantId,
		VaultClient: client,
	}

	return sess, nil
}

func AnalyticsService(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) (*session, error) {
	provider, err := getProvider(ctx, clientMeta, taskClient, task)

	client, err := analytics.NewAnalyticsClientWithConfigurationProvider(provider)
	if err != nil {
		return nil, err
	}

	tenantId, err := provider.TenancyOCID()
	if err != nil {
		return nil, err
	}

	sess := &session{
		TenancyID:       tenantId,
		AnalyticsClient: client,
	}

	return sess, nil
}

func expandPath(filepath string) string {
	if strings.HasPrefix(filepath, fmt.Sprintf(constants.C, os.PathSeparator)) {
		filepath = path.Join(getHomeFolder(), filepath[2:])
	}
	return path.Clean(filepath)
}

func getHomeFolder() string {
	current, e := user.Current()
	if e != nil {

		home := os.Getenv(constants.HOME)
		if home == constants.Constants_21 {
			home = os.Getenv(constants.USERPROFILE)
		}
		return home
	}
	return current.HomeDir
}

func checkProfile(profile string, path string) (err error) {
	var profileRegex = regexp.MustCompile(`^\[(.*)\]`)
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	content := string(data)
	splitContent := strings.Split(content, constants.N)
	for _, line := range splitContent {
		if match := profileRegex.FindStringSubmatch(line); len(match) > 1 && match[1] == profile {
			return nil
		}
	}

	return fmt.Errorf(constants.Configurationfiledidnotcontainprofiles, profile)
}

func getEnvSettingWithBlankDefault(s string) string {
	return getEnvSettingWithDefault(s, constants.Constants_22)
}

func getEnvSettingWithDefault(s string, dv string) string {
	v := os.Getenv(constants.TFVAR + s)
	if v != constants.Constants_23 {
		return v
	}
	v = os.Getenv(constants.OCI + s)
	if v != constants.Constants_24 {
		return v
	}
	v = os.Getenv(s)
	if v != constants.Constants_25 {
		return v
	}
	return dv
}

func getCLIEnvVariables(variableName string) string {
	v := os.Getenv(constants.OCICLI + variableName)
	if v != constants.Constants_26 {
		return v
	}
	v = os.Getenv(constants.OCI + variableName)
	if v != constants.Constants_27 {
		return v
	}
	return constants.Constants_28
}

func getProviderFromCLIEnvironmentVariables() (oci_common.ConfigurationProvider, error) {
	var providers []oci_common.ConfigurationProvider
	privateKeyPath := getCLIEnvVariables(constants.KEYFILE)
	pemFileContent := constants.Constants_29
	if privateKeyPath != constants.Constants_30 {
		resolvedPath := expandPath(privateKeyPath)
		pemFileData, err := os.ReadFile(resolvedPath)
		if err != nil {
			return nil, fmt.Errorf(constants.CannotreadprivatekeyfromsErrorq, privateKeyPath, err)
		}
		pemFileContent = string(pemFileData)
	}

	cliApiKeyProvider := oci_common.NewRawConfigurationProvider(
		getCLIEnvVariables(constants.TENANCY),
		getCLIEnvVariables(constants.USER),
		getCLIEnvVariables(constants.REGION),
		getCLIEnvVariables(constants.FINGERPRINT),
		pemFileContent,
		pointer.ToStringPointer(constants.Constants_31),
	)
	if cliApiKeyProvider != nil {
		providers = append(providers, cliApiKeyProvider)
	}

	cliFileWithProfileProvider, _ := oci_common.ConfigurationProviderFromFileWithProfile(
		getCLIEnvVariables(constants.CONFIGFILE),
		getCLIEnvVariables(constants.PROFILE),
		getCLIEnvVariables(constants.Constants_32),
	)

	if cliFileWithProfileProvider != nil {
		providers = append(providers, cliFileWithProfileProvider)
	}

	cliFromFileProvider, _ := oci_common.ConfigurationProviderFromFile(
		getCLIEnvVariables(constants.CONFIGFILE),
		getCLIEnvVariables(constants.Constants_33),
	)

	if cliFromFileProvider != nil {
		providers = append(providers, cliFromFileProvider)
	}

	if len(providers) > 0 {
		return oci_common.ComposingConfigurationProvider(providers)
	}
	return nil, nil
}

func buildHttpClient() (httpClient *http.Client) {
	httpClient = &http.Client{
		Timeout: 0,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: 10000000000,
			}).DialContext,
			TLSHandshakeTimeout: 10000000000,
			TLSClientConfig:     &tls.Config{MinVersion: tls.VersionTLS12},
			Proxy:               http.ProxyFromEnvironment,
		},
	}
	return
}

func GetDefaultRetryPolicy(taskClient any) *oci_common.RetryPolicy {
	attempts := uint(9)
	minRetryDelay := 25 * time.Millisecond

	config := taskClient.(*OciClient).OciConfig
	if config.MaxErrorRetryAttempts != nil {
		attempts = uint(*config.MaxErrorRetryAttempts)
	}

	if config.MinErrorRetryDelay != nil {
		minRetryDelay = time.Duration(*config.MinErrorRetryDelay) * time.Millisecond
	}

	retryOnResponseCodes := func(r oci_common.OCIOperationResponse) bool {
		if r.Response != nil && r.Response.HTTPResponse() != nil {
			statusCode := strconv.Itoa(r.Response.HTTPResponse().StatusCode)
			return (r.Error != nil && in(statusCode, []string{constants.Constants_34, constants.Constants_35, constants.Constants_36}))
		}
		return false
	}
	return getExponentialBackoffRetryPolicy(attempts, minRetryDelay, retryOnResponseCodes)
}

func getExponentialBackoffRetryPolicy(n uint, minRetryDelay time.Duration, fn func(r oci_common.OCIOperationResponse) bool) *oci_common.RetryPolicy {

	exponentialBackoff := func(r oci_common.OCIOperationResponse) time.Duration {

		var jitter = float64(rand.Intn(120-80)+80) / 100

		maxDelayTime := time.Duration(int(float64(int(minRetryDelay.Nanoseconds())*int(math.Pow(3, float64(r.AttemptNumber)))) * jitter))
		if maxDelayTime > time.Duration(3*time.Minute) {
			return time.Duration(3 * time.Minute)
		}

		return maxDelayTime
	}
	policy := oci_common.NewRetryPolicy(n, fn, exponentialBackoff)
	return &policy
}

