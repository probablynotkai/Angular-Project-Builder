package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

var Red = "\033[31m"
var Reset = "\033[0m"

var rootDir string
var targetDir string
var applicationName string

func main() {
	flags, err := fetchMappedArgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	rootDir = filepath.Dir(ex)

	if len(flags) < 2 && os.Args[1] == "-help" {
		log.Printf("\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
			" ",
			"Usage: angular-gen -d <targetDir> -n <appName> -l <libraries>",
			" ",
			"No \"QUOTATION MARKS\" are needed as flag values are parsed between-the-flag.",
			" ",
			"Templates and assets will be injected into the Angular application starting from the src ",
			"or assets folder, respectively. Nested directories will be traversed and duplicated.",
			" ",
			"Flags:",
			" -d <directory> = Directory to create project",
			" -n <appName> = Name of the project",
			" -l <libraries> = (Optional) Comma separated list of libraries")
		return
	}

	// Create project
	targetDir, applicationName = generateAngularProject(flags)

	// Add required libraries
	addLibrary("@angular/material")
	addLibrary("jquery")
	addLibrary("bootstrap@4")

	// Add additional libraries, if included
	if flags["l"] != "" {
		log.Println("Retrieving and adding custom libraries...")
		customLibraries := split(",", flags["l"])
		for _, v := range customLibraries {
			addLibrary(v)
		}
	}

	// Update angular.json to add styles, scripts, and a few other things
	updateAngularJson()

	// Generate header component and http service
	generateHeader()
	generateErrorLogger()
	generateHttpService()

	// Inject templates & assets
	injectTemplates()
	injectAssets()
	injectFavicon()

	log.Printf("Successfully generated the Angular project '%s' in the directory %s.\n", applicationName, targetDir)
	log.Println("")
	log.Println(Red + "Please ensure that you change the <base href=\"/\"> element in your index.html in order to work with IIS." + Reset)
	log.Println("")
}

func split(s string, delimiter string) []string {
	args := []string{}

	var ce string
	for _, c := range s {
		if string(c) != delimiter {
			ce = ce + string(c)
		} else {
			if ce != "" {
				args = append(args, ce)
			}

			ce = ""
		}
	}

	if ce != "" {
		args = append(args, ce)
	}

	return args
}

func fetchMappedArgs(args []string) (map[string]string, error) {
	if len(args) < 1 {
		return nil, errors.New("no arguments supplied, use -help for more info")
	}

	mp := map[string]string{}

	flag := ""
	val := ""
	for _, v := range args {
		if v[0] == 055 {
			if flag != "" && flag != v {
				if val == "" {
					mp[flag[1:]] = "true"
				} else {
					mp[flag[1:]] = val
				}
				val = ""
			}

			flag = v
		} else {
			if val == "" {
				val = v
			} else {
				val = val + " " + v
			}
		}
	}

	if val == "" {
		mp[flag[1:]] = "true"
	} else {
		mp[flag[1:]] = val
	}

	return mp, nil
}
