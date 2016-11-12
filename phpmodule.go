package main

import (
	"os"
	"phpmodule/scaffold"
	"fmt"
)

// Main project function.
func main() {
	// Create the structure

	// Act on first argument depending on its value.
	switch os.Args[1] {

		// Build a module
		case "init":

			//Check the arg length, ensuring there is a name.
			if len(os.Args) > 2 {

				// Build the structure.
				scaffold.Build(os.Args[2]);
			}else {

				// There must be a name.
				fmt.Println("You must specify a name. Please see --help for information.")
			}

		break

		// Create a component within the module.
		case "create":
		break
	}
}

/*

	//********
	// MODULE
	//********

	Format: program action name

	// Create Module
	$ phpmodule init Billing



	//*******************
	// MODULE COMPONENTS
	//*******************

	Format: program action type (name)

	// Create Module Data
	$ phpmodule create data Customer

	// Create Module Enum
	phpmodule create enum PlanCode

	// Create Module Exception
	phpmodule create exception (name) 	# Name Optional

	// Create Module Test
	phpmodule create test Customer


*/