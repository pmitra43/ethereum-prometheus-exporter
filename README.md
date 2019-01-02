# Ethereum Prometheus Exporter

This service exports various metrics from Ethereum clients for consumption by [Prometheus](https://prometheus.io). It uses [JSON-RPC](https://github.com/ethereum/wiki/wiki/JSON-RPC) interface to collect the metrics. Any JSON-RPC 2.0 enabled client should be supported. Although, it was only tested with [Parity](https://wiki.parity.io/Parity-Ethereum).

## Exported Metrics

| Name | Description |
| ---- | ----------- |
| net_peers | The number of peers currently connected to the client. |
| eth_block_number | The number of most recent block. |
| eth_gas_price | The current price per gas in wei. *Might be inaccurate*. |
| eth_earliest_block_transactions | The number of transactions in an earliest block. |
| eth_latest_block_transactions | The number of transactions in a latest block. |
| eth_pending_block_transactions | The number of transactions in a pending block. |
| eth_hashrate | The number of hashes per second that the node is mining with. |
| eth_sync_starting | The block at which the import started. |
| eth_sync_current | The number of most recent block. |
| eth_sync_highest | The estimated highest block. |
| parity_net_active_peers | The number of active peers. *Available only for Parity*. |
| parity_net_connected_peers | The number of peers currently connected to the client. *Available only for Parity*. |

## Development

[Go modules](https://golang.org/doc/go1.11#modules) is used for dependency management. Hence Go 1.11 is a minimum required version.

[CircleCI Local CLI](https://circleci.com/docs/2.0/local-cli/) can be used to ensure that everything builds locally.

    circleci build --job lint
    circleci build --job test
    circleci build --job build

## Contributing

Contributions are greatly appreciated. The project follows the typical GitHub pull request model. Before starting any work, please either comment on an existing issue or file a new one.

## Support and Donate

You can show your appreciation for the project and support future development by donating.

[![Donate with Ethereum](https://en.cryptobadges.io/badge/big/0xcaDe516c2c2d916eDf44b958ED5B52C01039fad6)](https://en.cryptobadges.io/donate/0xcaDe516c2c2d916eDf44b958ED5B52C01039fad6)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
