package main

import (
	"encoding/json"
)

type AngularConfig struct {
	Schema      string                     `json:"$schema"`
	Version     int64                      `json:"version"`
	ProjectRoot string                     `json:"newProjectRoot"`
	Projects    map[string]ProjectTemplate `json:"projects"`
}

// type Project struct {
// 	Main ProjectTemplate `json:"templateProject"`
// }

type ProjectTemplate struct {
	Type       string    `json:"projectType"`
	Schematics Blank     `json:"schematics"`
	Root       string    `json:"root"`
	SourceRoot string    `json:"sourceRoot"`
	Prefix     string    `json:"prefix"`
	Architect  Architect `json:"architect"`
}

type Blank struct{}

type Architect struct {
	Build       Build       `json:"build"`
	Serve       Serve       `json:"serve"`
	ExtractI18n ExtractI18n `json:"extract-i18n"`
	Test        Test        `json:"test"`
}

type Build struct {
	Builder              string         `json:"builder"`
	Options              BuildOptions   `json:"options"`
	Configurations       Configurations `json:"configurations"`
	DefaultConfiguration string         `json:"defaultConfiguration"`
}

type BuildOptions struct {
	OutputPath string   `json:"outputPath"`
	Index      string   `json:"index"`
	Main       string   `json:"main"`
	Polyfills  string   `json:"polyfills"`
	TsConfig   string   `json:"tsConfig"`
	Assets     []string `json:"assets"`
	Styles     []string `json:"styles"`
	Scripts    []string `json:"scripts"`
}

type Configurations struct {
	Production  ProdConfiguration `json:"production"`
	Development DevConfiguration  `json:"development"`
}

type ProdConfiguration struct {
	Budgets          []Budget `json:"budgets"`
	FileReplacements []struct {
		Replace string `json:"replace"`
		With    string `json:"with"`
	} `json:"fileReplacements"`
	OutputHashing string `json:"outputHashing"`
}

type DevConfiguration struct {
	BuildOptimizer  bool `json:"buildOptimizer"`
	Optimization    bool `json:"optimization"`
	VendorChunk     bool `json:"vendorChunk"`
	ExtractLicenses bool `json:"extractLicenses"`
	SourceMap       bool `json:"sourceMap"`
	NamedChunks     bool `json:"namedChunks"`
}

type Budget struct {
	Type           string `json:"type"`
	MaximumWarning string `json:"maximumWarning"`
	MaximumError   string `json:"maximumError"`
}

type Serve struct {
	Builder              string              `json:"builder"`
	Configurations       ServeConfigurations `json:"configurations"`
	DefaultConfiguration string              `json:"defaultConfiguration"`
}

type ServeConfigurations struct {
	Production struct {
		BrowserTarget string `json:"browserTarget"`
	} `json:"production"`
	Development struct {
		BrowserTarget string `json:"browserTarget"`
	} `json:"development"`
}

type ExtractI18n struct {
	Builder string `json:"builder"`
	Options struct {
		BrowserTarget string `json:"browserTarget"`
	} `json:"options"`
}

type Test struct {
	Builder string      `json:"builder"`
	Options TestOptions `json:"options"`
}

type TestOptions struct {
	Main        string   `json:"main"`
	Polyfills   string   `json:"polyfills"`
	TsConfig    string   `json:"tsConfig"`
	KarmaConfig string   `json:"karmaConfig"`
	Assets      []string `json:"assets"`
	Styles      []string `json:"styles"`
	Scripts     []string `json:"scripts"`
}

func (ac *AngularConfig) ForProject(projectName string) ([]byte, error) {
	m, _ := json.Marshal(ac.Projects)

	var a interface{}
	json.Unmarshal(m, &a)
	b := a.(map[string]interface{})

	b[projectName] = b["templateProject"]
	delete(b, "templateProject")

	return json.Marshal(b)
}
