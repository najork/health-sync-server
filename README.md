# health-sync-server
The health-sync-server retrieves timeseries data from external sources such as Garmin Connect and syncs them to Prometheus. Part of the [health timesteries platform](https://github.com/najork/health-timeseries-platform).

## Usage

Start the server.

```
./health-sync-server server
```

Export the metrics from an activity.

```
curl -XPOST -k https://localhost:8443/health-sync/api/collect -H 'Content-Type: application/json' -d '{"id":$activity_id,"maxPoints":50}'
```

Import the metrics into Prometheus. See the [documentation](https://github.com/prometheus/prometheus/blob/v2.29.1/docs/storage.md#backfilling-from-openmetrics-format) for more information.

```
promtool tsdb create-blocks-from openmetrics metrics.out
```
