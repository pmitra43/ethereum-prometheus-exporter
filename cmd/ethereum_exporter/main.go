package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pmitra43/ethereum-prometheus-exporter/internal/collector"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

func main() {
	flag.Usage = func() {
		const (
			usage = "Usage: ethereum_exporter [options]\n\n" +
				"Prometheus exporter for Ethereum client metrics\n\n" +
				"Options:\n"
		)

		fmt.Fprint(flag.CommandLine.Output(), usage)
		flag.PrintDefaults()

		os.Exit(2)
	}

	url := flag.String("url", "http://localhost:8545", "Ethereum JSON-RPC URL")
	addr := flag.String("addr", ":9368", "listen address")

	flag.Parse()
	if len(flag.Args()) > 0 {
		flag.Usage()
	}

	rpc, err := rpc.Dial(*url)
	if err != nil {
		log.Fatal(err)
	}

	registry := prometheus.NewPedanticRegistry()
	registry.MustRegister(
		collector.NewEthBlockNumber(rpc),
		collector.NewEthLatestBlockTransactions(rpc),
		collector.NewEthPendingBlockTransactions(rpc),
		collector.NewNetPeerCount(rpc),
		collector.NewLastBlockSeconds(rpc),
	)

	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		ErrorLog:      log.New(os.Stderr, log.Prefix(), log.Flags()),
		ErrorHandling: promhttp.ContinueOnError,
	})

	log.Printf("Started listening %s\n", time.Now())

	http.Handle("/metrics", handler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
