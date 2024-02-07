package main

import (
	"errors"
	"io"
	"log"
	"os"
)

func injectTemplates(dir string) {
	log.Println("Injecting templates into application...")
	pathToComponent := dir + "\\src"
	_, err := os.Stat(pathToComponent)
	if err != nil {
		log.Fatal(err)
	}

	templateRoot := ".\\templates"
	_, err = os.Stat(templateRoot)
	if err != nil {
		log.Fatal(err)
	}

	injectNestedFiles(templateRoot, pathToComponent)
}

func injectAssets(dir string) {
	log.Println("Injecting assets...")
	pathToInject := dir + "\\src\\assets"
	_, err := os.Stat(pathToInject)
	if err != nil {
		log.Fatal(err)
	}

	pathToAssets := ".\\assets"
	_, err = os.Stat(pathToAssets)
	if err != nil {
		log.Fatal(err)
	}

	injectNestedFiles(pathToAssets, pathToInject)
}

func injectFavicon(dir string) {
	log.Println("Injecting favicon...")
	pathToInject := dir + "\\src\\favicon.ico"
	_, err := os.Stat(pathToInject)
	if err != nil {
		if errors.Is(os.ErrNotExist, err) {
			os.Mkdir(pathToInject, os.ModePerm)
		} else {
			log.Fatal(err)
		}
	}

	pathToFavicon := ".\\assets\\favicon.ico"
	_, err = os.Stat(pathToFavicon)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(pathToFavicon)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile(pathToInject, data, 0644)
}

func injectNestedFiles(root string, target string) {
	files, err := os.ReadDir(root)
	if err != nil {
		return
	}

	_, err = os.Stat(target)
	if err != nil {
		os.Mkdir(target, os.ModeDir)
	}

	for _, file := range files {
		if file.Name() == "favicon.ico" {
			continue
		}

		if file.IsDir() {
			injectNestedFiles(root+"\\"+file.Name(), target+"\\"+file.Name())
		} else {
			f, _ := os.Open(root + "\\" + file.Name())
			defer f.Close()

			b, _ := io.ReadAll(f)

			os.WriteFile(target+"\\"+file.Name(), b, 0644)
		}
	}
}
