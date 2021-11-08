package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func buildFile() {
	var stderr bytes.Buffer
	log.Println("[+] Building...")
	if ext := filepath.Ext(os.Args[1]); ext == ".cpp" {
		fBuild := fmt.Sprintf("%s.exe", strings.Split(os.Args[1], ".cpp")[0])
		cmd := exec.Command("g++", os.Args[1], "-std=c++14", "-o", fBuild)
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			log.Println("[+] Compile Error!")
			fmt.Print(stderr.String())
			os.Exit(1)
		}
	} else if ext == ".c" {
		fBuild := fmt.Sprintf("%s.exe", strings.Split(os.Args[1], ".c")[0])
		cmd := exec.Command("gcc", os.Args[1], "-o", fBuild)
		cmd.Stderr = &stderr
		err := cmd.Run()
		if err != nil {
			log.Println("[+] Compile Error!")
			fmt.Print(stderr.String())
			os.Exit(1)
		}
	} else {
		log.Fatal("[+] File invalid!")
	}
}

func runFile() {
	var stdout, stderr bytes.Buffer
	stdinFile := ""
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	files, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		name := f.Name()
		if strings.Contains(name, ".in") {
			dat, err := os.ReadFile(name)
			if err != nil {
				log.Fatal(err)
			}
			stdinFile = string(dat)
			break
		}
	}
	fBuild := fmt.Sprintf("%s.exe", strings.Split(os.Args[1], ".c")[0])
	time.Sleep(time.Second)
	cmd := exec.Command(fBuild)
	cmd.Stdin = strings.NewReader(stdinFile + "\n")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(stderr.String())
		os.Exit(1)
	}
	fmt.Println(stdout.String())
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s filename(.cpp|.c)", os.Args[0])
		return
	}
	startBuilding := time.Now()
	buildFile()
	timeToBuild := time.Since(startBuilding)
	log.Println("[+] Compiled", timeToBuild)
	fmt.Println("---------------------------------------------")
	runFile()
}
