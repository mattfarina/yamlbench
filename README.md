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
BenchmarkGhodss-8                   	       1	39737474696 ns/op	7933007208 B/op	152831528 allocs/op
BenchmarkGoYaml-8                   	       1	20787000729 ns/op	3886198200 B/op	84590302 allocs/op
BenchmarkPrettyJson-8               	       1	6387788762 ns/op	1559883688 B/op	11620231 allocs/op
BenchmarkJson-8                     	       1	5089828723 ns/op	1427313304 B/op	11620118 allocs/op
BenchmarkEmbedPointerPrettyJson-8   	       1	6092759874 ns/op	920753368 B/op	12620129 allocs/op
BenchmarkEmbedPointerJson-8         	       1	5241154178 ns/op	788197752 B/op	12620129 allocs/op
```

The projects being evaluated are:

* github.com/ghodss/yaml
* gopkg.in/yaml.v2
* standard library encoding/json

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