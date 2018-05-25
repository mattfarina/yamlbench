# YAML Benchmarking

This project provides some tools to benchmark YAML Unmarshaling in Go. The
motivation for this analysis was handling of large `index.yaml` files for
[Helm](https://helm.sh).

## Results

```
go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/mattfarina/yamlbench
BenchmarkGhodss-8                         	       1	37544364242 ns/op	7924747864 B/op	152573816 allocs/op
BenchmarkGoYaml-8                         	       1	17608089326 ns/op	3859904312 B/op	84502899 allocs/op
BenchmarkPrettyJson-8                     	       1	5630493765 ns/op	1292438552 B/op	11504330 allocs/op
BenchmarkJson-8                           	       1	4755094692 ns/op	1159920888 B/op	11504227 allocs/op
BenchmarkEmbedPointerPrettyJson-8         	       1	5799941405 ns/op	915286888 B/op	12504232 allocs/op
BenchmarkEmbedPointerJson-8               	       1	4879612802 ns/op	782781880 B/op	12504230 allocs/op
BenchmarkJsonParser-8                     	       2	 912638128 ns/op	433769756 B/op	 1518987 allocs/op
BenchmarkEmbedPointerPrettyJsonParser-8   	       1	1276917704 ns/op	807032664 B/op	 1519269 allocs/op
BenchmarkJsonParserFileSysRead-8          	       2	 915497066 ns/op	325590600 B/op	 1518962 allocs/op
BenchmarkPrettyJsonParserFileSysRead-8    	       1	1220139431 ns/op	458225912 B/op	 1519414 allocs/op
PASS
ok  	github.com/mattfarina/yamlbench	85.121s
```

Note, the last 4 items on the list are an entirely different style. Only the items you
query are allocated and the JSON itself is parsed as needed. The allocations and memory
are entirely dependent on what you need to get out and how often you need to do so.

The projects being evaluated are:

* github.com/ghodss/yaml
* gopkg.in/yaml.v2
* standard library encoding/json
* github.com/buger/jsonparser

Notes on Embeded pointers:

* JSON was evaluated with and without embedded structs in other structs as pointers.
The memory usage is different which is relevant.
* github.com/ghodss/yaml was used with embedded struct pointers which is how Helm
currently has it.
* gopkg.in/yaml.v2 does not work with embedded pointers to structs and instead
requires non-pointer embeds.

## Running Them Yourself

1. Make sure you have both [Go](https://golang.org) and [dep](https://github.com/golang/dep)
installed.
1. Run `make setup` to install the dependencies and to generate test files. Files
used for testing are not included because they are over 750mb in total.
1. Run `make bench` to run the benchmarks