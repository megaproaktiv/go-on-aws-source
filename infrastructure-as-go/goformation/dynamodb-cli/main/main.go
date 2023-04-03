package main

import (
	// go
	"fmt"
	"os"
	"strings"

	// own
	"gfcli"
	// utils
	"path/filepath"
	// goformation
	"github.com/awslabs/goformation/v5"
	// cli-tools
	"github.com/thatisuday/clapper"
)



func main() {

	gfcli.Logger.Info("Starting")
	
	const cmdDestroyString = "destroy"
	const cmdDeployString = "deploy"
	const cmdStatusString = "status"
	const cmdHelpString = "help"

	const flagTemplate = "template"
	// Look for commands
	registry := clapper.NewRegistry()
	registry.Register(cmdDeployString)
	registry.Register(cmdDestroyString)
	registry.Register(cmdStatusString)
	registry.Register(cmdHelpString)
	


	// parse command-line arguments
	command, err := registry.Parse(os.Args[1:])

	// check for command line error
	if err != nil {
		gfcli.Logger.Errorf("error => %#v\n", err)
		help()
		return
	}

	// no command
	if( len(command.Name) == 0 ) {
		help()
		os.Exit(1)
	}
	cmd := command.Name;

	
	var stackname = "demotemplate"
	template,err := gfcli.CreateTemplate(stackname);
	if err != nil {
		gfcli.Logger.Errorf("error => %#v\n", err)
	}

	for flagName, flagValue := range command.Flags {
		if strings.Compare(flagName, flagTemplate) == 0 {
			// Check exist
			path := (flagValue.Value)
			if _, err := os.Stat(path); os.IsNotExist(err) {
				gfcli.Logger.Errorf("Template file ",path, "does not exist")
				panic(err)
			}
			template, err = goformation.Open(path)
			if err != nil {
				gfcli.Logger.Fatal("There was an error processing the template: %s", err)
				
			}
			// Template name is filename
			basename := filepath.Base(path)
			name := strings.TrimSuffix(basename, filepath.Ext(basename))
			stackname = name
			
		}
	}
	if cmd == cmdDeployString {
		// Create a new CloudFormation template	
		// Check template switch
		
		gfcli.CreateStack(gfcli.Client,stackname, template)
		gfcli.ShowStatus(gfcli.Client,stackname,template,gfcli.StatusCreateComplete);
	}
	
	if cmd == cmdDestroyString {
		gfcli.DeleteStack(gfcli.Client,stackname)
		gfcli.ShowStatus(gfcli.Client,stackname,template,gfcli.StatusDeleteComplete);
	}
	
	if cmd == cmdStatusString {
		gfcli.ShowStatus(gfcli.Client,stackname,template,gfcli.StatusCreateComplete);
	}

	
	if cmd == cmdHelpString {
		help();
		os.Exit(0);
	}

}

func help(){
	fmt.Println("CloudFormation deploy app. ")
	fmt.Println("CloudFormation Template is generated automatically.")
	fmt.Println("Please call with [deploy|destroy|status|help] . ")
}