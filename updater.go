package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

var standardStyles = []string{
	"./node_modules/@angular/material/prebuilt-themes/indigo-pink.css",
	"./node_modules/bootstrap/scss/bootstrap.scss",
	"src/styles.scss",
}

var standardScripts = []string{
	"./node_modules/jquery/dist/jquery.min.js",
	"./node_modules/bootstrap/dist/js/bootstrap.min.js",
}

func updateAngularJson() {
	log.Println("Updating and optimising angular.json...")
	file, err := os.Open(targetDir + "\\angular.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var cfg AngularConfig
	json.Unmarshal(b, &cfg) // Get values
	by, _ := cfg.ForProject(applicationName)
	json.Unmarshal(by, &cfg) // Update map values

	if entry, ok := cfg.Projects[applicationName]; ok {
		orgBudgets := []Budget{{
			Type:           "initial",
			MaximumWarning: "4mb",
			MaximumError:   "5mb",
		}, {
			Type:           "anyComponentStyle",
			MaximumWarning: "2kb",
			MaximumError:   "4kb",
		}}

		entry.Architect.Build.Options.Styles = standardStyles
		entry.Architect.Build.Options.Scripts = standardScripts
		entry.Architect.Build.DefaultConfiguration = "development"
		entry.Architect.Build.Configurations.Production.Budgets = orgBudgets

		cfg.Projects[applicationName] = entry
	}

	nb, err := json.MarshalIndent(cfg, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(targetDir+"\\angular.json", nb, 0644)
}
