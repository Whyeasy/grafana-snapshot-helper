FROM alpine

COPY grafana-snapshot-helper /usr/bin/
ENTRYPOINT ["/usr/bin/grafana-snapshot-helper"]