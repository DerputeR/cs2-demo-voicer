package main

import (
	"fmt"
	"os"

	// dem "github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	docopt "github.com/docopt/docopt-go"
)

type programConfig struct {
	ConfigPath  string `docopt:"config-path"`
	DemoInPath  string `docopt:"demo-in-path"`
	DemoOutPath string `docopt:"demo-out-path"`
	AudioDir    string `docopt:"audio-dir"`
}

var version string = "demovoicer 0.0.0"
var usage string = `Tool to add external voice comms into CS2 demos.

Usage:
  dmeovoicer <config_path>
  demovoicer <demo-in-path> <demo-out-path> <audio-dir> 
  demovoicer --version
  demovoicer -h | --help

Options:
  -h --help  Show this screen.
  --version  Show version.`

func main() {
	/// BASIC IO:
	// open demo
	argParser := &docopt.Parser{
		HelpHandler:   docopt.PrintHelpAndExit,
		OptionsFirst:  false,
		SkipHelpFlags: false,
	}
	var config programConfig
	opts, _ := argParser.ParseArgs(usage, os.Args[1:], version)
	opts.Bind(&config)

	fmt.Println(config.ConfigPath)
	fmt.Println(config.DemoInPath)
	fmt.Println(config.DemoOutPath)
	fmt.Println(config.AudioDir)

	// parse demo (load entirely into memory? prob should just stream it...)
	// figure out player id's (todo: find out if they change at half time switch)

	/// INITIAL CHECK:
	// rewrite parsed demo back to .dem
	// check hash to make sure they match (if they don't...what can we do?)
	// this initial check will not exist once we actually move onto the voice section

	/// AUDIO TECH:
	// figure out how to inject data into the demo that it didn't have before
	// figure out how how voice data is stored in the .dem/protobufs
	// figure out how the voice data is encoded
	// write/find a library to encode audio back into protobuf form
}
