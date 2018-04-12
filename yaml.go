package yamlbench

import (
	"time"
)

type IndexFile2 struct {
	APIVersion string                    `json:"apiVersion" yaml:"apiVersion"`
	Generated  time.Time                 `json:"generated" yaml:"generated"`
	Entries    map[string]ChartVersions2 `json:"entries" yaml:"entries"`
}

type ChartVersions2 []*ChartVersion2

type ChartVersion2 struct {
	*Metadata
	URLs    []string  `json:"urls" yaml:"urls"`
	Created time.Time `json:"created,omitempty" yaml:"created,omitempty"`
	Removed bool      `json:"removed,omitempty" yaml:"removed,omitempty"`
	Digest  string    `json:"digest,omitempty" yaml:"digest,omitempty"`
}

type IndexFile struct {
	APIVersion string                   `json:"apiVersion" yaml:"apiVersion"`
	Generated  time.Time                `json:"generated" yaml:"generated"`
	Entries    map[string]ChartVersions `json:"entries" yaml:"entries"`
}

type ChartVersions []ChartVersion

type ChartVersion struct {
	Metadata `yaml:",inline"`
	URLs     []string  `json:"urls" yaml:"urls"`
	Created  time.Time `json:"created,omitempty" yaml:"created,omitempty"`
	Removed  bool      `json:"removed,omitempty" yaml:"removed,omitempty"`
	Digest   string    `json:"digest,omitempty" yaml:"digest,omitempty"`
}

type Metadata struct {
	Name          string            `json:"name,omitempty" yaml:"name,omitempty"`
	Home          string            `json:"home,omitempty" yaml:"home,omitempty"`
	Sources       []string          `json:"sources,omitempty" yaml:"sources,omitempty"`
	Version       string            `json:"version,omitempty" yaml:"version,omitempty"`
	Description   string            `json:"description,omitempty" yaml:"description,omitempty"`
	Keywords      []string          `json:"keywords,omitempty" yaml:"keywords,omitempty"`
	Maintainers   []*Maintainer     `json:"maintainers,omitempty" yaml:"maintainers,omitempty"`
	Engine        string            `json:"engine,omitempty" yaml:"engine,omitempty"`
	Icon          string            `json:"icon,omitempty" yaml:"icon,omitempty"`
	ApiVersion    string            `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`
	Condition     string            `json:"condition,omitempty" yaml:"condition,omitempty"`
	Tags          string            `json:"tags,omitempty" yaml:"tags,omitempty"`
	AppVersion    string            `json:"appVersion,omitempty" yaml:"appVersion,omitempty"`
	Deprecated    bool              `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	TillerVersion string            `json:"tillerVersion,omitempty" yaml:"tillerVersion,omitempty"`
	Annotations   map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	KubeVersion   string            `json:"kubeVersion,omitempty" yaml:"kubeVersion,omitempty"`
}

type Maintainer struct {
	Name  string `json:"name,omitempty" yaml:"name,omitempty"`
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
	Url   string `json:"url,omitempty" yaml:"url,omitempty"`
}
