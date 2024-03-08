package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	type Config struct {
		Directory string `json:"Directory"`
	}
	userCommand := ""
	// userArg := ""

	if len(os.Args) >= 2 {
		userCommand = os.Args[1]
	}
	if len(os.Args) >= 3 {
		// userArg = os.Args[2]
	}
	d, err := os.Getwd()
	if err != nil {
		fmt.Println("Error trying to get current dir: ", err)
		return
	}

	if userCommand == "Init" {
		initJson := "New-Item -Path .\\config.json -ItemType File"
		jsonMsg := map[string]interface{}{
			// "" + userArg + "": "TestMessage",
			"Directory": "" + d + "",
		}
		jsonData, err := json.MarshalIndent(jsonMsg, "", "   ")
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		cmd := exec.Command("powershell", "-Command", initJson)

		err = cmd.Start()
		if err != nil {
			fmt.Println("Error creating JsonFile: ", err)
			return
		}
		err = os.WriteFile("config.json", jsonData, 0644)
		if err != nil {
			fmt.Println("Error Writing to config.json: ", err)
			return
		}
		fmt.Println("Initilzed Json")
		return
	}
	if userCommand == "Nvim" {
		cmd := exec.Command("nvim", ""+d+"/main.go")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error running nvim:", err)
			return
		}
	}
}
