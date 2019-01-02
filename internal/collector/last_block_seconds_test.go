package collector

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
)

func TestLastBlockSecondsError(t *testing.T) {
	rpc, err := rpc.DialHTTP("http://localhost")
	if err != nil {
		t.Fatalf("rpc connection error: %#v", err)
	}

	collector := NewLastBlockSeconds(rpc)
	ch := make(chan prometheus.Metric, 1)

	collector.Collect(ch)
	close(ch)

	if got := len(ch); got != 1 {
		t.Fatalf("got %v, want 1", got)
	}

	var metric dto.Metric
	for result := range ch {
		err := result.Write(&metric)
		if err == nil {
			t.Fatalf("expected invalid metric, got %#v", metric)
		}
		if _, ok := err.(*url.Error); !ok {
			t.Fatalf("unexpected error %#v", err)
		}
	}
}

func TestLastBlockSecondsCollect(t *testing.T) {
	rpcServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{"result":{"difficulty":"0x20000","extraData":"0x0000000000000000000000000000000000000000000000000000000000000000f84432b84130f7b38d287781158c87bc40bd310cbc310b2a351f1008edf1e7351bad0ebfe4786630101c3446456f343b6d0f0b582220211e66c3c975757a415f111c665ea400","gasLimit":"0x47faa3","gasUsed":"0xcf08","hash":"0xb3c16f71258e3b5f50a6f79017c516c9f5d2526f5395483ed6aab873139d3fe0","logsBloom":"0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000","miner":"0x0000000000000000000000000000000000000000","mixHash":"0x0000000000000000000000000000000000000000000000000000000000000000","nonce":"0x0000000000000000","number":"0xf","parentHash":"0xd7845cbf53d8b30d83b2afd1c1e57505eb0c237809bd9386f0de44643ffc5406","receiptsRoot":"0x65c4d1b533902562730f0d110bcce51643ce4725ca03d85c33d5812c0ed308bf","sha3Uncles":"0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347","size":"0x2bd","stateRoot":"0x4b332dd3b0484955f7c0f2bb46ab61f3d10799c9d9422f83718237c1aa0ea26e","timestamp":"0x157654eb67de5048","totalDifficulty":"0x1e0000","transactions":["0x92a98d4fc1377703f10f15c1630cda38913cf47bac28d2eae7ec3e72e3e96d40"],"transactionsRoot":"0x8d93c79cdcc93580cfc0c0062dfcf7c7e7431d1a66ae7ff2d2a67d0ffaf3b179","uncles":[]}}"`))
		if err != nil {
			t.Fatalf("could not write a response: %#v", err)
		}
	}))
	defer rpcServer.Close()

	rpc, err := rpc.DialHTTP(rpcServer.URL)
	if err != nil {
		t.Fatalf("rpc connection error: %#v", err)
	}

	collector := NewLastBlockSeconds(rpc)
	ch := make(chan prometheus.Metric, 1)

	collector.Collect(ch)
	close(ch)

	if got := len(ch); got != 1 {
		t.Fatalf("got %v, want 1", got)
	}

	var metric dto.Metric
	for result := range ch {
		if err := result.Write(&metric); err != nil {
			t.Fatalf("expected metric, got %#v", err)
		}
		if got := len(metric.Label); got > 0 {
			t.Fatalf("expected 0 labels, got %d", got)
		}
		if got := *metric.Gauge.Value; got != 1546516892.085014600 {
			t.Fatalf("got %v, want 3220", got)
		}
	}
}
