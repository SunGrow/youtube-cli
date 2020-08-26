package main

import (
	"flag"
	"fmt"
	"youtube-cli/ttyinput"
	"youtube-cli/userset"
)


func main() {
	var state userset.InputState;
	flag.StringVar(&state.User, "u", "default", "Specify the current user.");
	flag.StringVar(&state.Output, "o", userset.GetWD()+"/feed.html", "Specify the output file.");
	flag.StringVar(&state.ConfigFile, "cf", userset.GetConfigDir()+"/"+state.User+"/sub_list.csv", "Specify the csv subscription feed file.");
	flag.Usage = func() {
		fmt.Print(
			"Usage:	youtube-cli [options] command...\n",
			"Commands:\n",
			"  feed\n",
			"\tGenerate user subscription feed\n",
			"  sublist\n",
			"\tGenerate user subscription list to generate subcription feed from\n",
			"Options:\n",
		);
		flag.PrintDefaults();
	}
	flag.Parse();
	function_map := ttyinput.GetArgFunctionMap();
	for _, arg:= range(flag.Args()){
		if function_map[arg] != nil {
			function_map[arg](state);
		}
	}

	return;
}
