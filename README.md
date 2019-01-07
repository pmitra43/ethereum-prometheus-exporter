# Geth Prometheus Exporter

This service exports various metrics from Ethereum clients for consumption by [Prometheus](https://prometheus.io). It uses [JSON-RPC](https://github.com/ethereum/wiki/wiki/JSON-RPC) interface to collect the metrics. Any JSON-RPC 2.0 enabled client should be supported. Although, it was only tested with [Parity](https://wiki.parity.io/Parity-Ethereum).

## Exported Metrics

| Name | Description |
| ---- | ----------- |
| net_peers | The number of peers currently connected to the client. |
| eth_block_number | The number of most recent block. |
| eth_latest_block_transactions | The number of transactions in a latest block. |
| eth_pending_block_transactions | The number of transactions in a pending block. |
| last_block_seconds | duration in seconds of the last block since epoch |

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
