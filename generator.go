package main

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
)

func generateAngularProject(args map[string]string) (string, string) {
	log.Println("Generating Angular project...")
	// Is valid directory supplied?
	if args["d"] == "" {
		log.Fatal(errors.New("no directory supplied"))
	}
	_, err := os.Stat(args["d"])
	if err != nil {
		os.Mkdir(args["d"], os.ModePerm)
	}

	// Is name supplied?
	if args["n"] == "" {
		log.Fatal(errors.New("no name supplied"))
	}
	space := containSpace(args["n"])
	if space {
		log.Fatal(errors.New("name has space, violates angular name schema"))
	}

	path, err := exec.LookPath("ng")
	if err != nil {
		log.Fatal(err)
	}

	cmdArgs := []string{"new", args["n"], "--defaults=true", "--routing=true", "--style=scss"}
	cmd := exec.Command(path, cmdArgs...)
	if !errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}

	var serr bytes.Buffer

	cmd.Dir = args["d"]
	cmd.Stderr = &serr
	_, err = cmd.Output()
	if err != nil {
		log.Fatal(errors.New(serr.String()))
	}

	log.Printf("Generated Angular project under name '%s'.", args["n"])
	return args["d"] + "\\" + args["n"], args["n"]
}

func generateHeader() {
	log.Println("Generating header component...")
	path, err := exec.LookPath("ng")
	if err != nil {
		log.Fatal(err)
	}

	var serr bytes.Buffer
	cmdArgs := []string{"g", "c", "components/Header"}

	cmd := exec.Command(path, cmdArgs...)
	cmd.Dir = targetDir
	cmd.Stderr = &serr
	if _, err := cmd.Output(); err != nil {
		log.Fatal(serr.String())
	}
}

func generateErrorLogger() {
	log.Println("Generating error logger template...")
	path, err := exec.LookPath("ng")
	if err != nil {
		log.Fatal(err)
	}

	var serr bytes.Buffer
	cmdArgs := []string{"g", "class", "logging/ErrorLogger"}

	cmd := exec.Command(path, cmdArgs...)
	cmd.Dir = targetDir
	cmd.Stderr = &serr
	if _, err := cmd.Output(); err != nil {
		log.Fatal(serr.String())
	}
}

func generateHttpService() {
	log.Println("Generating HTTP service...")
	path, err := exec.LookPath("ng")
	if err != nil {
		log.Fatal(err)
	}

	var serr bytes.Buffer
	cmdArgs := []string{"g", "s", "services/http"}

	cmd := exec.Command(path, cmdArgs...)
	cmd.Dir = targetDir
	cmd.Stderr = &serr
	if _, err := cmd.Output(); err != nil {
		log.Fatal(serr.String())
	}
}

func containSpace(s string) bool {
	for _, v := range s {
		if v == rune(040) {
			return true
		}
	}
	return false
}
