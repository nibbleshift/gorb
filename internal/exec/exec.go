package exec

import (
	"fmt"
	"os/exec"
	"strings"

	benchmark "golang.org/x/tools/benchmark/parse"
)

type Exec struct {
	padding int
}

func NewExec() *Exec {
	return &Exec{}
}

type Result struct {
	OS      string
	Arch    string
	CPU     string
	Package string
	Bench   []benchmark.Benchmark `json:"bench"`
	Pass    bool
}

func (e *Exec) Run(cmd ...string) Result {
	out, err := exec.Command("go", "test", "-bench=.", "-benchtime=1x").Output()
	if err != nil {
		panic(err)
	}

	r := Result{}

	for _, line := range strings.Split(string(out), "\n") {
		if strings.HasPrefix(line, "Benchmark") {
			b, err := benchmark.ParseLine(line)

			if err != nil {
				fmt.Println(err)
				continue
			}
			r.Bench = append(r.Bench, *b)
		} else if strings.HasPrefix(line, "goos") {
			fmt.Sscanf(line, "goos: %s", &r.OS)
		} else if strings.HasPrefix(line, "goarch") {
			fmt.Sscanf(line, "goarch: %s", &r.Arch)
		} else if strings.HasPrefix(line, "pkg") {
			fmt.Sscanf(line, "pkg: %s", &r.Package)
		} else if strings.HasPrefix(line, "cpu") {
			r.CPU, _ = strings.CutPrefix(line, "cpu: ")
		} else if line == "PASS" {
			r.Pass = true
		} else {
			fmt.Println(line)
		}
	}
	return r
}
