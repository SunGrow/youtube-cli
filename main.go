package main

import (
	"flag"
	"fmt"
	"youtube-cli/ttyinput"
	"youtube-cli/userset"
)

func main() {
	var user string;
	var configDir string;
	flag.StringVar(&user, "u", "default", "Specify the current user.");
	flag.StringVar(&configDir, "cd", userset.GetConfigDir(), "Specify current config directory.");
	flag.Usage = func() {
		fmt.Printf("Usage:	youtube-cli [options] command...\n");
		flag.PrintDefaults();
	}
	flag.Parse();
	function_map := ttyinput.GetArgFunctionMap();
	for _, arg:= range(flag.Args()){
		if function_map[arg] != nil {
			function_map[arg](user, configDir);
		}
	}

	return;
}
