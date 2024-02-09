package main

import (
	"log"
	"os/exec"
)

func addLibrary(libName string) {
	log.Printf("Adding library %s...\n", libName)
	path, err := exec.LookPath("ng")
	if err != nil {
		log.Fatal(err)
	}

	cmdArgs := []string{"add", libName, "--skip-confirmation"}
	cmd := exec.Command(path, cmdArgs...)
	if cmd.Err != nil {
		log.Fatal(cmd.Err)
	}

	cmd.Dir = targetDir
	cmd.Output()
}
