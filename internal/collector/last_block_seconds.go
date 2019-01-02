package collector

import (
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ethereum/go-ethereum/core/types"
	"encoding/json"
	"math"
)

type LastBlockSeconds struct {
	rpc  *rpc.Client
	desc *prometheus.Desc
}

func NewLastBlockSeconds(rpc *rpc.Client) *LastBlockSeconds{
	return &LastBlockSeconds{
		rpc: rpc,
		desc: prometheus.NewDesc(
			"last_block_seconds",
			"timestamp in seconds of the last block",
			nil,
			nil,
		),
	}
}

func (collector *LastBlockSeconds) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.desc
}

func (collector *LastBlockSeconds) Collect(ch chan<- prometheus.Metric) {
	var result json.RawMessage
	if err := collector.rpc.Call(&result, "eth_getBlockByNumber", "latest", false); err != nil {
		ch <- prometheus.NewInvalidMetric(collector.desc, err)
		return
	}
	var head *types.Header
	json.Unmarshal(result, &head)
	pow := math.Pow(10, 9)
	value := float64(head.Time.Int64())/pow
	ch <- prometheus.MustNewConstMetric(collector.desc, prometheus.GaugeValue, value)
}
