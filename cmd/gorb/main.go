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

	var bench client.Bench

	bench.Os = result.OS
	bench.Arch = result.Arch
	bench.CPU = result.CPU
	bench.Package = result.Package
	bench.Pass = result.Pass

	for _, r := range result.Results {
		br := &BenchResult{
			Name:              r.Name,
			N:                 r.N,
			NsPerOp:           r.NsPerOp,
			AllocedBytesPerOp: r.AllocedBytesPerOp,
			AllocsPerOp:       r.AllocsPerOp,
			MBPerS:            r.MBPerS,
			Measured:          r.Measured,
			Ord:               r.Ord,
		}
	}

	_, err = gorbClient.CreateBench(ctx, bench)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}

func NewGorbClient(c *clientv2.Client) *gorb.Client {
	return &gorb.Client{Client: c}
}
