package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var (
	addToEndC  = "printf(\"\\n-------------------------\\n\");system(\"pause\");return 0;}"
	addToEndCpp = "cout << endl << \"-------------------------\" << endl;system(\"pause\");return 0;}"
	tempFolder = ""
)

func buildFile() {
	var stderr bytes.Buffer
	log.Println("[+] Building...")
	localApp, err := os.UserCacheDir()

	if err != nil {
		log.Fatal("[+] Folder appdata not found!")
	}
	tempFolder = fmt.Sprintf("%s\\Temp", localApp)

	if _, err := os.Stat(tempFolder); os.IsNotExist(err) {
		log.Fatal("[+] Folder Temp not found!")
	}

	if ext := filepath.Ext(os.Args[1]); ext == ".cpp" {
		dat, er := os.ReadFile(os.Args[1])
		if er != nil {
			log.Fatal("[+] Could not read file ", os.Args[1])
		}

		fileContent := strings.TrimSpace(string(dat))
		fileContent = strings.TrimSuffix(fileContent, "}")
		fileContent = strings.TrimSpace(fileContent)
		fileContent = strings.TrimSuffix(fileContent, "return 0;")
		fileContent += addToEndCpp

		newSrcFile := fmt.Sprintf("%s\\%s", tempFolder, os.Args[1])

		f, err := os.Create(newSrcFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		_, err = f.WriteString(fileContent)
		if err != nil {
			log.Fatal(err)
		}

		fBuild := fmt.Sprintf("%s.exe", strings.Split(os.Args[1], ".cpp")[0])
		cmd := exec.Command("g++", newSrcFile, "-std=c++14", "-o", fBuild)
		cmd.Stderr = &stderr
		err = cmd.Run()
		if err != nil {
			log.Println("[+] Compile Error!")
			fmt.Print(stderr.String())
			os.Exit(1)
		}
	} else if ext == ".c" {
		dat, er := os.ReadFile(os.Args[1])
		if er != nil {
			log.Fatal("[+] Could not read file ", os.Args[1])
		}

		fileContent := strings.TrimSpace(string(dat))
		fileContent = strings.TrimSuffix(fileContent, "}")
		fileContent = strings.TrimSpace(fileContent)
		fileContent = strings.TrimSuffix(fileContent, "return 0;")
		fileContent += addToEndC

		newSrcFile := fmt.Sprintf("%s\\%s", tempFolder, os.Args[1])

		f, err := os.Create(newSrcFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		_, err = f.WriteString(fileContent)
		if err != nil {
			log.Fatal(err)
		}
		fBuild := fmt.Sprintf("%s.exe", strings.Split(os.Args[1], ".c")[0])
		cmd := exec.Command("gcc", newSrcFile, "-o", fBuild)
		cmd.Stderr = &stderr
		err = cmd.Run()
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
	fBuild := fmt.Sprintf("%s.exe", strings.Split(os.Args[1], ".c")[0])
	time.Sleep(time.Second)
	cmd := exec.Command("cmd", "/c", "start", fmt.Sprintf(".\\%s", fBuild))
	cmd.Run()
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
	runFile()
}
