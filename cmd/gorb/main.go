package main

import (
	"context"

	_ "github.com/lib/pq"
	"github.com/nibbleshift/gorb/internal/exec"

	"fmt"
	"net/http"
	"os"

	"github.com/Yamashou/gqlgenc/clientv2"
	"github.com/nibbleshift/gorb/pkg/client"
	gorb "github.com/nibbleshift/gorb/pkg/client"
)

func main() {
	var (
		err error
	)
	e := exec.NewExec()
	result := e.Run("/usr/local/bin/go", "-test", "-bench=.")

	gorbClient := NewGorbClient(clientv2.NewClient(http.DefaultClient, "http://localhost:8081/query", nil,
		func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {

			return next(ctx, req, gqlInfo, res)
		}))
	ctx := context.Background()

	var bench client.CreateBenchInput

	bench.Os = result.OS
	bench.Arch = result.Arch
	bench.CPU = result.CPU
	bench.Package = result.Package
	bench.Pass = result.Pass

	var benchResults []*client.CreateBenchResultInput
	for _, r := range result.Bench {
		br := &client.CreateBenchResultInput{
			Name:              r.Name,
			N:                 int64(r.N),
			NsPerOp:           r.NsPerOp,
			AllocedBytesPerOp: int64(r.AllocedBytesPerOp),
			AllocsPerOp:       int64(r.AllocsPerOp),
			MbPerS:            float64(r.MBPerS),
			Measured:          int64(r.Measured),
			Ord:               int64(r.Ord),
		}
		benchResults = append(benchResults, br)
	}

	_, err = gorbClient.CreateBench(ctx, bench, benchResults)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}

func NewGorbClient(c *clientv2.Client) *gorb.Client {
	return &gorb.Client{Client: c}
}
