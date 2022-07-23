package main

import (
	"encoding/json"
	"envars/mapper"
	"fmt"
	"io/ioutil"
	"os"
)

const help string = `
envars command allows to map env vars into a config.json file

Positional Arguments
jsonpath  the path of the config.json file to map the env vars

Flags
-h  --help  Show the command help`

func main() {
	if len(os.Args) == 1 {
		fmt.Println(help)
		return
	}

	args := os.Args[1:]

	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			fmt.Println(help)
			return
		}
	}

	jsonpath := args[0]

	if len(jsonpath) == 0 {
		fmt.Println("the 'jsonpath' positional argument cannot be empty")
		os.Exit(1)
		return
	}

	var config map[string]interface{}

	if b, e := ioutil.ReadFile(jsonpath); e != nil {
		fmt.Println(e)
		os.Exit(2)
		return
	} else {
		if e2 := json.Unmarshal(b, &config); e2 != nil {
			fmt.Println(e)
			os.Exit(3)
			return
		}
	}

	if e := mapper.MapVars(&config); e != nil {
		fmt.Println(e)
		os.Exit(4)
		return
	}

	if b, e := json.MarshalIndent(config, "", "  "); e != nil {
		fmt.Println(e)
		os.Exit(5)
		return
	} else {
		if e2 := ioutil.WriteFile(jsonpath, b, 0777); e != nil {
			fmt.Println(e2)
			os.Exit(6)
			return
		}
	}
}
