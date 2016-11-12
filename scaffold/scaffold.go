package scaffold

import (
	"os"
	"fmt"
	"html/template"
)

// The base directory
const OUTPUT_DIR = "output"

// The base directory
const TEMPLATE_DIR = "scaffold/templates"



// ------
// Wrapper for the methods that do the work.
// --
func Build(name string) {

	// Build directory structure
	Directories(name);

	// Populate with basic files.
	Files(name);
}


// ------
// Handles the creation of the folder structure for the module.
// --
func Directories(name string) {

	// Tell the user whats going on.
	fmt.Println("Building the folder structure...");

	// List of paths to create.
	paths := []string{
		name + "/Config/Dynamic",
		name + "/Container",
		name + "/Data",
		name + "/Enum",
		name + "/Exception",
		name + "/Handler",
		name + "/Library",
		name + "/Migration",
		name + "/Processor/Cron",
		name + "/Processor/Event",
		name + "/Processor/Interval",
		name + "/Processor/Single",
		name + "/Processor/Worker",
		name + "/Processor/Test",
	}

	// Loop the paths.
	for _, path := range paths {

		// Make the path.
		os.MkdirAll(OUTPUT_DIR + "/" + path, os.ModePerm)
	}

	// Directories taken care of, let them know.
	fmt.Println("Done.");
}


// ------
// Handles the creations of files for the basic module.
// --
func Files (name string) {

	// Tell the user whats going on.
	fmt.Println("Adding in some files you may find useful...");

	// Template Values - For All!
	values := map[string]interface{} {
		"Name": name,
	}

	// Desired Output Path => Template Path
	files := map[string] string {
		OUTPUT_DIR + "/composer.json": TEMPLATE_DIR + "/_composer.json",
		//OUTPUT_DIR + "/Module.php": TEMPLATE_DIR + "/_ ",
	}

	// Loop them files!
	for outputPath, tpl := range files {

		// Parse the template.
		t, err := template.ParseFiles(tpl)

		// Check for errors.
		checkError(err)

		// Create output path.
		file, err := os.Create(outputPath)
		// Execute.
		err = t.Execute(file, values)
	}

	// Base files all sorted, let them know.
	fmt.Println("Done.");
}


// ------
// Fatal error checking.
// --
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}






