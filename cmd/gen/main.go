// The gen command builds an example yaml file that's fairly massive to use in
// benchmarking.
//
// Some quick design notes...
// - Instead of using yaml tooling using templates
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
	"github.com/mattfarina/yamlbench"
)

var header = `apiVersion: v1
entries:
`

var footer = "generated: 2018-04-11T16:56:56.656249201Z"

var chart = `  dummy-chart-{{.Num}}:
`

var release = `  - created: 2017-07-06T01:33:50.952906435Z
    description: Example description
    digest: 249e27501dbfe1bd93d4039b04440f0ff19c707ba720540f391b5aefa3571455
    home: https://example.com
    icon: https://example.com/foo.png
    keywords:
    - A
    - B
    maintainers:
    - email: bar@example.com
      name: Bar
    name: dummy-chart-{{.Num}}
    sources:
    - https://example.com
    - https://example.com
    urls:
    - https://example.com
    version: 1.2.{{.Num2}}
`

type wrapper struct {
	Num  int
	Num2 int
}

func main() {

	// Generate a YAML file for testing
	genYaml()

	// Also generate a json version of the same file content for testing
	genJson()

	fmt.Println("Done generating testing files")
}

func genYaml() {
	// TODO(mattfarina): create an argument to capture the file name.

	fmt.Println("Generating index.yaml for testing")

	f, err := os.Create("./index.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.WriteString(header)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	charttmpl, err := template.New("chart").Parse(chart)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	reltmpl, err := template.New("release").Parse(release)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var w wrapper
	for i := 0; i < 5000; i++ {
		w = wrapper{Num: i}
		err = charttmpl.Execute(f, w)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for j := 0; j < 100; j++ {
			w.Num2 = j
			err = reltmpl.Execute(f, w)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	_, err = f.WriteString(footer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func genJson() {
	yml, err := ioutil.ReadFile("./index.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Generating json files")

	in := &yamlbench.IndexFile{}
	err = yaml.Unmarshal(yml, &in)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Generating index.pretty.json for testing")
	out, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile("./index.pretty.json", out, 0644)

	fmt.Println("Generating index.json for testing")
	out, err = json.Marshal(in)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile("./index.json", out, 0644)
}
