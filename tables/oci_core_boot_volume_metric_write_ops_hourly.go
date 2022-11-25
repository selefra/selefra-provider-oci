package tables

import (
	"context"
	"strings"
	"time"

	"github.com/oracle/oci-go-sdk/v44/common"
	"github.com/oracle/oci-go-sdk/v44/core"
	"github.com/oracle/oci-go-sdk/v44/monitoring"
	"github.com/selefra/selefra-provider-oci/oci_client"
	"github.com/selefra/selefra-provider-oci/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

type TableOciCoreBootVolumeMetricWriteOpsHourlyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableOciCoreBootVolumeMetricWriteOpsHourlyGenerator{}

func (x *TableOciCoreBootVolumeMetricWriteOpsHourlyGenerator) GetTableName() string {
	return "oci_core_boot_volume_metric_write_ops_hourly"
}

func (x *TableOciCoreBootVolumeMetricWriteOpsHourlyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableOciCoreBootVolumeMetricWriteOpsHourlyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableOciCoreBootVolumeMetricWriteOpsHourlyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableOciCoreBootVolumeMetricWriteOpsHourlyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			volume := task.ParentRawResult.(core.BootVolume)
			region := taskClient.(*oci_client.OciClient).Region

			_, err := listMonitoringMetricStatistics(ctx, clientMeta, taskClient, task, resultChannel, "HOURLY", "oci_blockstore", "VolumeWriteOps", "resourceId", *volume.Id, *volume.CompartmentId, region)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}

			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
		},
	}
}

type MetricData struct {
	CompartmentId *string
	PointValue    *float64
	Timestamp     *time.Time
}
type MonitoringMetricRow struct {
	CompartmentId *string

	DimensionName *string

	DimensionValue *string

	Namespace *string

	MetricName *string

	Average *float64

	Maximum *float64

	Minimum *float64

	SampleCount *float64

	Sum *float64

	Timestamp *time.Time

	Unit *string

	Metadata map[string]string

	Region string
}

func filterMetricStatistic(metricStatistic monitoring.SummarizeMetricsDataResponse) []MetricData {
	metricData := []MetricData{}
	for _, item := range metricStatistic.Items {
		for _, data := range item.AggregatedDatapoints {
			metricData = append(metricData, MetricData{
				CompartmentId: item.CompartmentId,
				PointValue:    data.Value,
				Timestamp:     &data.Timestamp.Time,
			})
		}
	}
	return metricData
}
func getMonitoringPeriodForGranularity(granularity string) string {
	switch strings.ToUpper(granularity) {
	case "DAILY":

		return "1d"
	case "HOURLY":

		return "1h"
	}

	return "5m"
}
func getMonitoringStartDateForGranularity(granularity string) time.Time {
	switch strings.ToUpper(granularity) {
	case "DAILY":

		return time.Now().AddDate(0, 0, -90)
	case "HOURLY":

		return time.Now().AddDate(0, 0, -60)
	}

	return time.Now().AddDate(0, 0, -5)
}
func getStatisticForColumnByTimestamp(timestamp time.Time, compartmentId string, metricData []MetricData) *float64 {
	var value *float64
	for _, t := range metricData {
		if *t.Timestamp == timestamp && compartmentId == *t.CompartmentId {
			value = t.PointValue
			break
		}
	}

	return value
}
func listMonitoringMetricStatistics(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any, granularity string, namespace string, metricName string, dimensionName string, dimensionValue string, compartmentId string, region string) (*monitoring.SummarizeMetricsDataResponse, error) {

	session, err := oci_client.MonitoringService(ctx, clientMeta, taskClient, task)
	if err != nil {
		return nil, err
	}

	queryString := metricName + "[" + getMonitoringPeriodForGranularity(granularity) + "]" + "{" + dimensionName + " = \"" + dimensionValue + "\"}"
	queryStringavg := queryString + ".grouping().mean()"
	querystringMin := queryString + ".grouping().min()"
	querystringMax := queryString + ".grouping().max()"
	querystringSum := queryString + ".grouping().sum()"
	querystringCount := queryString + ".grouping().count()"

	interval := getMonitoringPeriodForGranularity(granularity)
	metricDetails := monitoring.SummarizeMetricsDataDetails{
		Namespace:  &namespace,
		StartTime:  &common.SDKTime{Time: getMonitoringStartDateForGranularity(granularity)},
		EndTime:    &common.SDKTime{Time: time.Now()},
		Query:      &queryStringavg,
		Resolution: &interval,
	}

	requestParam := monitoring.SummarizeMetricsDataRequest{
		CompartmentId:               &compartmentId,
		SummarizeMetricsDataDetails: metricDetails,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: oci_client.GetDefaultRetryPolicy(taskClient),
		},
	}

	avgStatistics, err := session.MonitoringClient.SummarizeMetricsData(ctx, requestParam)
	if err != nil {
		return nil, err
	}

	metricDetails.Query = &querystringMin
	requestParam.SummarizeMetricsDataDetails = metricDetails
	minStatistics, err := session.MonitoringClient.SummarizeMetricsData(ctx, requestParam)
	if err != nil {
		return nil, err
	}
	metricDetailsMin := filterMetricStatistic(minStatistics)

	metricDetails.Query = &querystringMax
	requestParam.SummarizeMetricsDataDetails = metricDetails
	maxStatistics, err := session.MonitoringClient.SummarizeMetricsData(ctx, requestParam)
	if err != nil {
		return nil, err
	}
	metricDetailsMax := filterMetricStatistic(maxStatistics)

	metricDetails.Query = &querystringSum
	requestParam.SummarizeMetricsDataDetails = metricDetails
	sumStatistics, err := session.MonitoringClient.SummarizeMetricsData(ctx, requestParam)
	if err != nil {
		return nil, err
	}
	metricDetailsSum := filterMetricStatistic(sumStatistics)

	metricDetails.Query = &querystringCount
	requestParam.SummarizeMetricsDataDetails = metricDetails
	countStatistics, err := session.MonitoringClient.SummarizeMetricsData(ctx, requestParam)
	if err != nil {
		return nil, err
	}
	metricDetailsCount := filterMetricStatistic(countStatistics)

	for _, item := range avgStatistics.Items {
		for _, datapoint := range item.AggregatedDatapoints {
			resultChannel <- &MonitoringMetricRow{
				CompartmentId:  item.CompartmentId,
				DimensionValue: &dimensionValue,
				DimensionName:  &dimensionName,
				Namespace:      &namespace,
				MetricName:     &metricName,
				Average:        datapoint.Value,
				Maximum:        getStatisticForColumnByTimestamp(datapoint.Timestamp.Time.UTC(), *item.CompartmentId, metricDetailsMax),
				Minimum:        getStatisticForColumnByTimestamp(datapoint.Timestamp.Time.UTC(), *item.CompartmentId, metricDetailsMin),
				Timestamp:      &datapoint.Timestamp.Time,
				SampleCount:    getStatisticForColumnByTimestamp(datapoint.Timestamp.Time.UTC(), *item.CompartmentId, metricDetailsCount),
				Sum:            getStatisticForColumnByTimestamp(datapoint.Timestamp.Time.UTC(), *item.CompartmentId, metricDetailsSum),
				Metadata:       item.Metadata,
				Region:         region,
			}
		}
	}

	return nil, err
}

func (x *TableOciCoreBootVolumeMetricWriteOpsHourlyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return oci_client.BuildCompartementZonalList()
}

func (x *TableOciCoreBootVolumeMetricWriteOpsHourlyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("sample_count").ColumnType(schema.ColumnTypeFloat).Description("The number of metric values that contributed to the aggregate value of this data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sum").ColumnType(schema.ColumnTypeFloat).Description("The sum of the metric values for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("unit").ColumnType(schema.ColumnTypeString).Description("The standard unit for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timestamp").ColumnType(schema.ColumnTypeTimestamp).Description("The time stamp used for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("The OCID of the boot volume.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("average").ColumnType(schema.ColumnTypeFloat).Description("The average of the metric values that correspond to the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maximum").ColumnType(schema.ColumnTypeFloat).Description("The maximum metric value for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("minimum").ColumnType(schema.ColumnTypeFloat).Description("The minimum metric value for the data point.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).Description("The OCID of the Tenant in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metric_name").ColumnType(schema.ColumnTypeString).Description("The name of the metric.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).Description("The metric namespace.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).Description("The OCI region in which the resource is located.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compartment_id").ColumnType(schema.ColumnTypeString).Description("The ID of the compartment.").Build(),
	}
}

func (x *TableOciCoreBootVolumeMetricWriteOpsHourlyGenerator) GetSubTables() []*schema.Table {
	return nil
}
