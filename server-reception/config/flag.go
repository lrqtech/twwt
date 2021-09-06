package config

import (
	"flag"
	"fmt"
	"os"
)

var cliIniFilePath = flag.String("f", "", "Input Your Path OF Ini File")

func init() {
	flag.Parse()
	if *cliIniFilePath == "" {
		fmt.Println("Param error")
		fmt.Println("Please try again")
		fmt.Println("Use 'server -h' for help")
		os.Exit(0)
	} else {
		IniFilePath = *cliIniFilePath
		loadIniFile()
	}
}
