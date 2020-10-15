package cmd

import "github.com/famartinrh/appctl/pkg/types/app"

// var CfgFile string
var AppFile string

var Verbosity int

var AppConfig *app.AppConfig
var ProjectDir string
var ForceDowload bool

type AppctlConfig struct {
	Verbosity  int    `yaml:"verbosity"`
	CatalogURL string `yaml:"catalogURL"`
	Force      bool   `yaml:"force"`
}
