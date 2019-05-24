package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	//Command name is the first arg after the current executable
	//Arguments follow
	commandName := os.Args[1]
	argsWithoutProg := os.Args[2:]

	//Name of the file to read and write last past timestamp
	var timestampsFile = commandName + "_timestamps.json"

	var changed = true
	var m map[string]int64
	m = make(map[string]int64)

	//Read file containing past timestamps
	content, err := ioutil.ReadFile(timestampsFile)
	if err == nil {
		err2 := json.Unmarshal(content, &m)
		if err2 != nil {
			fmt.Print(err2)
		}
	}

	//Search args for .tmp filename
	var tmpFile string
	for i := 0; i < len(argsWithoutProg); i++ {
		str := argsWithoutProg[i]

		//TODO: make this smarter
		if !(str[0] == '-') {
			tmpFile = str
			break
		}
	}

	if tmpFile != "" {
		info, err := os.Stat(tmpFile)
		if err == nil {

			//Get file timestamp
			time := info.ModTime().Unix()

			//Compare with past value
			if val, ok := m[tmpFile]; ok {
				changed = (val != time)
			}

			//Update map of timestamps
			m[tmpFile] = time
		}
	}

	//If value is different than the past one:
	// 1. Run
	// 2. Update values in timestamps file
	if changed {
		cmd := exec.Command(commandName, argsWithoutProg...)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		err := cmd.Run()
		fmt.Print(out.String())
		if err != nil {
			//As long as errors occur, do not update the timestamps file
			//Not ok, return error code
			os.Exit(1)
		} else {
			//all ok, write
			data, _ := json.MarshalIndent(m, "", "  ")
			_ = ioutil.WriteFile(timestampsFile, data, 0644)
		}
	}
}
