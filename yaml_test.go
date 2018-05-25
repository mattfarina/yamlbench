package yamlbench

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/buger/jsonparser"
	"github.com/ghodss/yaml"

	y2 "gopkg.in/yaml.v2"
)

var index []byte
var jsonindex []byte
var prettyjsonindex []byte

func getIndex() []byte {
	if len(index) == 0 {
		var err error
		index, err = ioutil.ReadFile("./index.yaml")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	return index
}

func getJSONIndex() []byte {
	if len(jsonindex) == 0 {
		var err error
		jsonindex, err = ioutil.ReadFile("./index.json")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	return jsonindex
}

func getPrettyJSONIndex() []byte {
	if len(prettyjsonindex) == 0 {
		var err error
		prettyjsonindex, err = ioutil.ReadFile("./index.pretty.json")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	return prettyjsonindex
}

func BenchmarkGhodss(b *testing.B) {
	yml := getIndex()
	for n := 0; n < b.N; n++ {
		var in IndexFile2
		err := yaml.Unmarshal(yml, &in)
		if err != nil {
			b.Errorf("github.com/ghodss/yaml err: %s", err)
		}
	}
}

func BenchmarkGoYaml(b *testing.B) {
	yml := getIndex()
	for n := 0; n < b.N; n++ {
		var in IndexFile
		err := y2.Unmarshal(yml, &in)
		if err != nil {
			b.Errorf("gopkg.in/yaml.v2 err: %s", err)
		}
	}
}

func BenchmarkPrettyJson(b *testing.B) {
	index, err := ioutil.ReadFile("./index.pretty.json")
	if err != nil {
		b.Error(err)
	}
	for n := 0; n < b.N; n++ {
		var in IndexFile
		err := json.Unmarshal(index, &in)
		if err != nil {
			b.Errorf("json err: %s", err)
		}
	}
}

func BenchmarkJson(b *testing.B) {
	index, err := ioutil.ReadFile("./index.json")
	if err != nil {
		b.Error(err)
	}
	for n := 0; n < b.N; n++ {
		var in IndexFile
		err := json.Unmarshal(index, &in)
		if err != nil {
			b.Errorf("json err: %s", err)
		}
	}
}

func BenchmarkEmbedPointerPrettyJson(b *testing.B) {
	index, err := ioutil.ReadFile("./index.pretty.json")
	if err != nil {
		b.Error(err)
	}
	for n := 0; n < b.N; n++ {
		var in IndexFile2
		err := json.Unmarshal(index, &in)
		if err != nil {
			b.Errorf("json err: %s", err)
		}
	}
}

func BenchmarkEmbedPointerJson(b *testing.B) {
	index, err := ioutil.ReadFile("./index.json")
	if err != nil {
		b.Error(err)
	}
	for n := 0; n < b.N; n++ {
		var in IndexFile2
		err := json.Unmarshal(index, &in)
		if err != nil {
			b.Errorf("json err: %s", err)
		}
	}
}

func BenchmarkJsonParser(b *testing.B) {
	in, err := ioutil.ReadFile("./index.json")
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {

		// Since the json file needs to be in memory to perform actions on it
		// we are making the entire thing available here. This is a copy of
		// memory to memory so as to skip the allocations dealing with reading
		// from the file system.
		index := make([]byte, len(in))
		copy(index, in)

		// This slice allows us to store data without using it but keeping
		// that in memory.
		i := map[string]string{}
		var err error

		// You only act on the things you need. It's an entirely different model.
		// This test looks at working over the whole data set.
		i["apiversion"], err = jsonparser.GetString(index, "apiVersion")
		if err != nil {
			b.Error(err)
		}

		i["generated"], err = jsonparser.GetString(index, "generated")
		if err != nil {
			b.Error(err)
		}

		jsonparser.ObjectEach(index, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			jsonparser.ArrayEach(value, func(value2 []byte, dataType jsonparser.ValueType, offset2 int, err error) {
				n, err := jsonparser.GetString(value2, "name")
				if err != nil {
					b.Error(err)
				}
				v, err := jsonparser.GetString(value2, "version")
				if err != nil {
					b.Error(err)
				}
				i[string(key)+v] = n
			})

			return nil
		}, "entries")

	}
}

func BenchmarkEmbedPointerPrettyJsonParser(b *testing.B) {
	in, err := ioutil.ReadFile("./index.pretty.json")
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {

		// Since the json file needs to be in memory to perform actions on it
		// we are making the entire thing available here. This is a copy of
		// memory to memory so as to skip the allocations dealing with reading
		// from the file system.
		index := make([]byte, len(in))
		copy(index, in)

		// This slice allows us to store data without using it but keeping
		// that in memory.
		i := map[string]string{}
		var err error

		// You only act on the things you need. It's an entirely different model.
		// This test looks at working over the whole data set.
		i["apiversion"], err = jsonparser.GetString(index, "apiVersion")
		if err != nil {
			b.Error(err)
		}

		i["generated"], err = jsonparser.GetString(index, "generated")
		if err != nil {
			b.Error(err)
		}

		jsonparser.ObjectEach(index, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			jsonparser.ArrayEach(value, func(value2 []byte, dataType jsonparser.ValueType, offset2 int, err error) {
				n, err := jsonparser.GetString(value2, "name")
				if err != nil {
					b.Error(err)
				}
				v, err := jsonparser.GetString(value2, "version")
				if err != nil {
					b.Error(err)
				}
				i[string(key)+v] = n
			})

			return nil
		}, "entries")

	}
}

func BenchmarkJsonParserFileSysRead(b *testing.B) {

	for n := 0; n < b.N; n++ {

		// This includes the read from the filesystem so we can look at allocations
		index, err := ioutil.ReadFile("./index.json")
		if err != nil {
			b.Error(err)
		}

		// This slice allows us to store data without using it but keeping
		// that in memory.
		i := map[string]string{}

		// You only act on the things you need. It's an entirely different model.
		// This test looks at working over the whole data set.
		i["apiversion"], err = jsonparser.GetString(index, "apiVersion")
		if err != nil {
			b.Error(err)
		}

		i["generated"], err = jsonparser.GetString(index, "generated")
		if err != nil {
			b.Error(err)
		}

		jsonparser.ObjectEach(index, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			jsonparser.ArrayEach(value, func(value2 []byte, dataType jsonparser.ValueType, offset2 int, err error) {
				n, err := jsonparser.GetString(value2, "name")
				if err != nil {
					b.Error(err)
				}
				v, err := jsonparser.GetString(value2, "version")
				if err != nil {
					b.Error(err)
				}
				i[string(key)+v] = n
			})

			return nil
		}, "entries")

	}
}
func BenchmarkPrettyJsonParserFileSysRead(b *testing.B) {

	for n := 0; n < b.N; n++ {

		// This includes the read from the filesystem so we can look at allocations
		index, err := ioutil.ReadFile("./index.pretty.json")
		if err != nil {
			b.Error(err)
		}

		// This slice allows us to store data without using it but keeping
		// that in memory.
		i := map[string]string{}

		// You only act on the things you need. It's an entirely different model.
		// This test looks at working over the whole data set.
		i["apiversion"], err = jsonparser.GetString(index, "apiVersion")
		if err != nil {
			b.Error(err)
		}

		i["generated"], err = jsonparser.GetString(index, "generated")
		if err != nil {
			b.Error(err)
		}

		jsonparser.ObjectEach(index, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			jsonparser.ArrayEach(value, func(value2 []byte, dataType jsonparser.ValueType, offset2 int, err error) {
				n, err := jsonparser.GetString(value2, "name")
				if err != nil {
					b.Error(err)
				}
				v, err := jsonparser.GetString(value2, "version")
				if err != nil {
					b.Error(err)
				}
				i[string(key)+v] = n
			})

			return nil
		}, "entries")

	}
}
