#!/bin/bash
set -e
/entrypoint.sh &
API=/bin/statshouse-api
META=/bin/statshouse-metadata
AGENT=/bin/statshouse
AGGREGATOR=/bin/statshouse
DELVE="/bin/dlv exec --headless --listen=:8000 --api-version=2"
case $STATSHOUSE_DEBUG_TARGET in
  agent)
    AGENT="$DELVE $AGENT --"
    ;;
  aggregator)
    AGGREGATOR="$DELVE $AGGREGATOR --"
    ;;
  api)
    API="$DELVE $API --"
    ;;
  meta)
    META="$DELVE $META --"
    ;;
esac
$META --db-path=/var/lib/statshouse/metadata/db --binlog-prefix=/var/lib/statshouse/metadata/binlog/bl &
until clickhouse-client --query="SELECT 1"; do sleep 0.2; done
$AGGREGATOR aggregator --cluster=test_shard_aggregator --log-level=trace --agg-addr=':13336' --kh=127.0.0.1:8123 \
  --auto-create -cache-dir=/var/lib/statshouse/cache/aggregator -u=root -g=root &
$AGENT agent --cluster=test_shard_aggregator --log-level=trace --remote-write-enabled \
  --agg-addr='127.0.0.1:13336,127.0.0.1:13336,127.0.0.1:13336' --cache-dir=/var/lib/statshouse/cache/agent \
  -u=root -g=root &
$API --verbose --insecure-mode --local-mode --access-log --clickhouse-v1-addrs= --clickhouse-v2-addrs=127.0.0.1:9000 \
  --listen-addr=:10888 --statshouse-addr=127.0.0.1:13337 --disk-cache=/var/lib/statshouse/cache/api/mapping_cache.sqlite3 \
  --static-dir=/usr/lib/statshouse-api/statshouse-ui/ &
wait -n
exit $?
