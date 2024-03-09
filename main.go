package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// All comment are personal just to remind myself of some of the flow
func main() {
	type Config struct {
		Directory string `json:"Directory"`
	}
	userCommand := ""
	// userArg := ""

	// os.Args by default is 1 because when you use >My-Workspace that is already 1
	// Look for 2 or more when looking for arguments
	if len(os.Args) >= 2 {
		userCommand = os.Args[1]
	}
	if len(os.Args) >= 3 {
		// userArg = os.Args[2]
	}
	//wd = Working Directory
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
		//Marshaling is taking an Obj or Struct and formating into JSON string
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
