![build](https://github.com/Whyeasy/grafana-snapshot-helper/workflows/build/badge.svg)
![status-badge](https://goreportcard.com/badge/github.com/Whyeasy/grafana-snapshot-helper)
![Github go.mod Go version](https://img.shields.io/github/go-mod/go-version/Whyeasy/grafana-snapshot-helper)

# Grafana snapshot Helper

A Grafana sidecar that allows easy usage of the snapshot capability with a stateless Grafana.

## Requirements

Provide a Grafana username that can create API keys; `--username <string>` or as env variable `GRAFANA_USER`.

Provide a the corresponding password for this user; `--password <string>` or as env variables `GRAFANA_PASSWORD`
